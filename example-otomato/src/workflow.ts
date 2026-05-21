// Builds the Otomato workflow from a validated WorkflowConfig.
//
// Topology:
//
//   trigger_1 ─┐
//              ├─► SPLIT ─► AI(protocol_X) ─► IF(eq "true") ─┬─► HTTP_REQUEST → relay /detect
//   trigger_N ─┘                                             ├─► SEND_EMAIL recipient_1
//                                                            └─► ... SEND_EMAIL recipient_M
//
// Each Twitter account becomes one trigger (OR-combined via isOptional).
// Each protocol becomes one branch off the SPLIT node.
// On a positive hit (AI returns "true") the workflow fires the relay
// webhook AND emails every address in alertEmails.

import {
  Action,
  Trigger,
  Workflow,
  Edge,
  ConditionGroup,
  LOGIC_OPERATORS,
} from 'otomato-sdk';

import { ACTIONS, TRIGGERS } from './otomato-types.js';
import type { WorkflowConfig, Protocol, MonitoredAccount } from './config.js';
import { buildDetectionPrompt } from './prompt.js';

export const WORKFLOW_NAME = 'thatsRekt detector';

export function buildWorkflow(config: WorkflowConfig): Workflow {
  const { tracking, env } = config;

  const triggers = tracking.monitoredAccounts.map(buildTrigger);

  const firstTrigger = triggers[0];
  if (!firstTrigger) throw new Error('monitoredAccounts is empty');

  // Output variables are sourced from the first trigger; the split
  // node's children forward them through the pipeline automatically.
  const tweetIDVar        = firstTrigger.getOutputVariableName('tweetId');
  const tweetURLVar       = firstTrigger.getOutputVariableName('tweetURL');
  const tweetContentVar   = firstTrigger.getOutputVariableName('tweetContent');
  const tweetAccountVar   = firstTrigger.getOutputVariableName('account');
  const tweetTimestampVar = firstTrigger.getOutputVariableName('timestamp');
  const tweetImagesVar    = firstTrigger.getOutputVariableName('images');

  const split = new Action(ACTIONS.CORE.SPLIT.SPLIT);
  split.setParams('branchesAmount', tracking.protocols.length);

  const branches = tracking.protocols.map((protocol) =>
    buildProtocolBranch({
      protocol,
      alertEmails: tracking.alertEmails,
      webhookBaseURL: env.WEBHOOK_BASE_URL,
      webhookToken: env.WEBHOOK_TOKEN,
      webhookChain: env.WEBHOOK_CHAIN,
      tweetIDVar,
      tweetURLVar,
      tweetContentVar,
      tweetAccountVar,
      tweetTimestampVar,
      tweetImagesVar,
    }),
  );

  const edges: Edge[] = [
    ...triggers.map((t) => new Edge({ source: t, target: split })),
    ...branches.map((b) => new Edge({ source: split, target: b.ai })),
    ...branches.flatMap((b) => [
      new Edge({ source: b.ai, target: b.ifAction }),
      new Edge({ source: b.ifAction, target: b.http, label: 'true', value: true }),
      ...b.emails.map(
        (e) => new Edge({ source: b.ifAction, target: e, label: 'true', value: true }),
      ),
    ]),
  ];

  const allNodes = [
    ...triggers,
    split,
    ...branches.flatMap((b) => [b.ai, b.ifAction, b.http, ...b.emails]),
  ];

  return new Workflow(WORKFLOW_NAME, allNodes, edges);
}

function buildTrigger(account: MonitoredAccount): Trigger {
  const t = new Trigger(TRIGGERS.SOCIALS.X.X_POST_TRIGGER);
  t.setParams('username', account.username);
  t.setParams('includeRetweets', account.includeRetweets);
  // isOptional OR-combines multiple triggers — any single account firing
  // starts the pipeline.
  (t as unknown as { isOptional: boolean }).isOptional = true;
  return t;
}

interface BranchInput {
  readonly protocol: Protocol;
  readonly alertEmails: readonly string[];
  readonly webhookBaseURL: string;
  readonly webhookToken: string;
  readonly webhookChain: string;
  readonly tweetIDVar: string;
  readonly tweetURLVar: string;
  readonly tweetContentVar: string;
  readonly tweetAccountVar: string;
  readonly tweetTimestampVar: string;
  readonly tweetImagesVar: string;
}

interface ProtocolBranch {
  readonly ai: Action;
  readonly ifAction: Action;
  readonly http: Action;
  readonly emails: readonly Action[];
}

function buildProtocolBranch(input: BranchInput): ProtocolBranch {
  const {
    protocol,
    alertEmails,
    webhookBaseURL,
    webhookToken,
    webhookChain,
    tweetIDVar,
    tweetURLVar,
    tweetContentVar,
    tweetAccountVar,
    tweetTimestampVar,
    tweetImagesVar,
  } = input;

  const ai = new Action(ACTIONS.AI.AI.AI);
  ai.setParams('prompt', buildDetectionPrompt(protocol));
  ai.setParams('defaultMode', false);
  const aiResultVar = ai.getOutputVariableName('result');

  const ifAction = new Action(ACTIONS.CORE.CONDITION.IF);
  ifAction.setParams('logic', LOGIC_OPERATORS.OR);
  const condGroup = new ConditionGroup(LOGIC_OPERATORS.OR);
  condGroup.addConditionCheck(aiResultVar, 'eq', 'true');
  ifAction.setParams('groups', [condGroup]);

  // Raw tweet text in the body; metadata travels as headers so the
  // relay can synthesize the on-chain title without parsing JSON.
  const http = new Action(ACTIONS.CORE.HTTP_REQUEST.HTTP_REQUEST);
  http.setParams('url', `${webhookBaseURL}/detect`);
  http.setParams('method', 'POST');
  http.setParams(
    'headers',
    JSON.stringify({
      Authorization: `Bearer ${webhookToken}`,
      'Content-Type': 'text/plain',
      'X-Idempotency-Key': tweetIDVar,
      'X-Tweet-URL': tweetURLVar,
      'X-Tweet-Account': tweetAccountVar,
      'X-Tweet-Timestamp': tweetTimestampVar,
      'X-Chain': webhookChain,
      'X-Protocol': protocol.name,
      'X-Tweet-Images': tweetImagesVar,
    }),
  );
  http.setParams('body', tweetContentVar);

  const emails = alertEmails.map((to) =>
    buildEmailAction({ to, protocol, tweetAccountVar, tweetContentVar, tweetURLVar }),
  );

  return { ai, ifAction, http, emails };
}

interface EmailInput {
  readonly to: string;
  readonly protocol: Protocol;
  readonly tweetAccountVar: string;
  readonly tweetContentVar: string;
  readonly tweetURLVar: string;
}

function buildEmailAction(input: EmailInput): Action {
  const { to, protocol, tweetAccountVar, tweetContentVar, tweetURLVar } = input;
  const email = new Action(ACTIONS.NOTIFICATIONS.EMAIL.SEND_EMAIL);
  email.setParams('to', to);
  email.setParams('subject', `[thatsRekt] ${protocol.name} security alert`);
  email.setParams(
    'body',
    [
      `${tweetAccountVar} posted content that may indicate a ${protocol.name} security incident.`,
      ``,
      `Tweet: ${tweetContentVar}`,
      `Link:  ${tweetURLVar}`,
      ``,
      `(This alert is also being submitted on-chain to thatsRekt.)`,
    ].join('\n'),
  );
  return email;
}
