// `npm run check` — print the current state of deployed workflows.
//
// Reads workflow-ids.json (written by `npm run create`) and queries
// Otomato for each workflow's live state.

import dotenv from 'dotenv';
import { existsSync, readFileSync } from 'node:fs';
import { resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { Workflow, apiServices } from 'otomato-sdk';

import { loadEnv } from './config.js';

const projectRoot = resolve(fileURLToPath(import.meta.url), '..', '..');
const idsFile = resolve(projectRoot, 'workflow-ids.json');

dotenv.config({ path: resolve(projectRoot, '.env') });

async function main(): Promise<void> {
  if (!existsSync(idsFile)) {
    console.error(`workflow-ids.json not found. Run \`npm run create\` first.`);
    process.exit(1);
  }

  const env = loadEnv();
  apiServices.setUrl(env.OTOMATO_API_URL);
  apiServices.setAuth(env.OTOMATO_API_KEY);

  const ids = JSON.parse(readFileSync(idsFile, 'utf-8')) as Record<string, string>;
  const entries = Object.entries(ids);

  if (entries.length === 0) {
    console.log('No workflows in workflow-ids.json.');
    return;
  }

  console.log(`\nChecking ${entries.length} workflow(s)...\n`);

  for (const [name, id] of entries) {
    let state = 'unknown';
    try {
      const wf = new Workflow(name, [], []);
      await wf.load(id);
      state = wf.getState() ?? 'unknown';
    } catch (err) {
      state = `ERROR (${(err as Error).message})`;
    }
    const icon = state === 'active' ? '[OK]' : state.startsWith('ERROR') ? '[ERR]' : '[??]';
    console.log(`${icon}  ${name.padEnd(32)} ${state.padEnd(20)} ${id}`);
  }
}

main().catch((err: unknown) => {
  console.error('check failed:', err);
  process.exit(1);
});
