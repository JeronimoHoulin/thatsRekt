// `npm run create` — build and deploy the Otomato workflow.
//
// Reads tracking.json + .env, ships the workflow to Otomato, and writes
// the resulting workflow id to workflow-ids.json (gitignored).
//
// Each run creates a NEW workflow. To replace an existing one you must
// delete it first from app.otomato.xyz (Otomato has no delta-update API).

import dotenv from 'dotenv';
import { writeFileSync } from 'node:fs';
import { resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { apiServices } from 'otomato-sdk';

import { loadConfig } from './config.js';
import { buildWorkflow, WORKFLOW_NAME } from './workflow.js';

const projectRoot = resolve(fileURLToPath(import.meta.url), '..', '..');
const idsFile = resolve(projectRoot, 'workflow-ids.json');

dotenv.config({ path: resolve(projectRoot, '.env') });

async function main(): Promise<void> {
  const cfg = loadConfig();

  apiServices.setUrl(cfg.env.OTOMATO_API_URL);
  apiServices.setAuth(cfg.env.OTOMATO_API_KEY);

  const triggerCount    = cfg.tracking.monitoredAccounts.length;
  const branchCount     = cfg.tracking.protocols.length;
  const recipientCount  = cfg.tracking.alertEmails.length;
  const totalNodes      = triggerCount + 1 /* split */ + branchCount * (3 + recipientCount);

  console.log(
    `\nBuilding "${WORKFLOW_NAME}":\n` +
      `  ${triggerCount} X triggers\n` +
      `  ${branchCount} protocol branches\n` +
      `  ${recipientCount} email recipients per branch\n` +
      `  → ${totalNodes} total nodes\n` +
      `  webhook: ${cfg.env.WEBHOOK_BASE_URL}/detect  (chain=${cfg.env.WEBHOOK_CHAIN})\n`,
  );

  const workflow = buildWorkflow(cfg);

  const result = await workflow.create();
  if (!result.success) {
    console.error('workflow.create failed:', result.error);
    process.exit(1);
  }

  const id = workflow.id;
  if (typeof id !== 'string' || id.length === 0) {
    console.error('workflow.create returned without an id');
    process.exit(1);
  }

  console.log(`Workflow created — id: ${id}`);

  await workflow.run();
  console.log(`State: ${workflow.getState() ?? 'unknown'}`);

  writeFileSync(
    idsFile,
    JSON.stringify({ [WORKFLOW_NAME]: id }, null, 2) + '\n',
    'utf-8',
  );
  console.log(`Workflow id saved to ${idsFile}`);
}

main().catch((err: unknown) => {
  console.error('create failed:', err);
  process.exit(1);
});
