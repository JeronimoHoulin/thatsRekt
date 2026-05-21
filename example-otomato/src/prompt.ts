import type { Protocol } from './config.js';

// Builds the binary-classification prompt for a single protocol branch.
//
// Otomato's AI block returns a short string; the downstream IF gate only
// supports `eq`/`neq`. The only reliable contract is: return the exact
// lowercase literal "true" or "false" — no JSON, no markdown, nothing else.
export function buildDetectionPrompt(protocol: Protocol): string {
  const keywordList = protocol.keywords.join(', ');
  const handleClause = protocol.twitterHandle
    ? ` or directly tags/mentions the official account @${protocol.twitterHandle}`
    : '';

  return [
    `Task: classify whether the tweet describes an active security incident for the protocol "${protocol.name}".`,
    ``,
    `Output contract:`,
    `  - Reply with EXACTLY the literal lowercase string "true" or "false". Nothing else.`,
    `  - No JSON, no quotes, no punctuation, no prose, no whitespace.`,
    `  - The downstream gate compares your output with === "true" — any deviation fails silently.`,
    ``,
    `Decision rules:`,
    ``,
    `1. Output "true" only if the tweet directly states ${protocol.name} (keywords: ${keywordList}${handleClause}) has been hacked, exploited, drained, is under attack, or is experiencing a serious security/funds incident.`,
    ``,
    `2. Output "false" for price talk, governance, audits, partnerships, listings, positive announcements, marketing, or any unrelated content. When uncertain, default to "false".`,
    ``,
    `3. Match keywords as EXACT, COMPLETE tokens. A keyword that appears only as a substring of a longer token does NOT count (e.g. "tETH" must NOT match "stETH"). A token boundary is the start/end of the tweet OR a non-alphanumeric character on each side.`,
    ``,
    `4. A retweet that merely mentions the protocol name in passing → "false". Only direct reporting of an incident → "true".`,
    ``,
    `Reply now with EXACTLY one word: true or false.`,
  ].join('\n');
}
