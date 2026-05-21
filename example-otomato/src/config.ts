import { readFileSync } from 'node:fs';
import { resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { z } from 'zod';

const projectRoot = resolve(fileURLToPath(import.meta.url), '..', '..');

// ─── Schemas ─────────────────────────────────────────────────────────────────

export const MonitoredAccountSchema = z.object({
  username: z
    .string()
    .min(1)
    .regex(/^[A-Za-z0-9_]+$/, 'X usernames are alphanumeric + underscore'),
  includeRetweets: z.boolean(),
});

export const ProtocolSchema = z.object({
  name: z.string().min(1).max(50),
  twitterHandle: z.string().regex(/^[A-Za-z0-9_]+$/).optional(),
  keywords: z.array(z.string().min(1)).min(1).max(20),
});

export const TrackingFileSchema = z.object({
  alertEmails: z.array(z.string().email()).min(1).max(10),
  monitoredAccounts: z.array(MonitoredAccountSchema).min(1).max(20),
  protocols: z.array(ProtocolSchema).min(1).max(50),
});

export type MonitoredAccount = z.infer<typeof MonitoredAccountSchema>;
export type Protocol = z.infer<typeof ProtocolSchema>;
export type TrackingFile = z.infer<typeof TrackingFileSchema>;

// Env carries deployment-time secrets that should never appear in tracking.json.
export const EnvSchema = z.object({
  OTOMATO_API_KEY: z.string().min(1, 'OTOMATO_API_KEY is required'),
  OTOMATO_API_URL: z.string().url().default('https://api.otomato.xyz/api'),
  WEBHOOK_BASE_URL: z
    .string()
    .url('WEBHOOK_BASE_URL must be a full URL (e.g. https://your-relay.up.railway.app)')
    .refine((u) => !u.endsWith('/'), 'WEBHOOK_BASE_URL must not end with a trailing slash'),
  WEBHOOK_TOKEN: z.string().min(8, 'WEBHOOK_TOKEN is too short — use a real secret'),
  WEBHOOK_CHAIN: z.string().min(1).default('ethereum'),
});

export type Env = z.infer<typeof EnvSchema>;

export interface WorkflowConfig {
  readonly tracking: TrackingFile;
  readonly env: Env;
}

// ─── Loaders ─────────────────────────────────────────────────────────────────

export function loadTrackingFile(path?: string): TrackingFile {
  const target = path ?? resolve(projectRoot, 'tracking.json');
  const raw = readFileSync(target, 'utf-8');
  return TrackingFileSchema.parse(JSON.parse(raw) as unknown);
}

export function loadEnv(env: NodeJS.ProcessEnv = process.env): Env {
  return EnvSchema.parse(env);
}

export function loadConfig(): WorkflowConfig {
  return {
    tracking: loadTrackingFile(),
    env: loadEnv(),
  };
}
