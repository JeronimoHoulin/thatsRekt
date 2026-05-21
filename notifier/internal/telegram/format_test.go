package telegram_test

import (
	"strings"
	"testing"

	"github.com/ThatsRekt/thatsRekt/notifier/internal/graphql"
	"github.com/ThatsRekt/thatsRekt/notifier/internal/telegram"
)

// v2Note builds a self-describing note string for tests.
func v2Note(summary, chains, txs, sources string) string {
	var b strings.Builder
	if summary != "" {
		b.WriteString("summary: ")
		b.WriteString(summary)
		b.WriteString("\n")
	}
	if chains != "" {
		b.WriteString("chains: ")
		b.WriteString(chains)
		b.WriteString("\n")
	}
	if txs != "" {
		b.WriteString("txs: ")
		b.WriteString(txs)
		b.WriteString("\n")
	}
	if sources != "" {
		b.WriteString("sources: ")
		b.WriteString(sources)
		b.WriteString("\n")
	}
	return b.String()
}

func makePost(opts struct {
	title       string
	note        string
	actionCount int
	attackers   []string
	victims     []string
	chain       graphql.Chain
	updatedAt   string
}) graphql.Post {
	return graphql.Post{
		ID:           opts.chain.Slug + "-7",
		Chain:        opts.chain,
		Title:        opts.title,
		Note:         opts.note,
		ActionCount:  opts.actionCount,
		Attackers:    opts.attackers,
		Victims:      opts.victims,
		LastUpdatedAt: opts.updatedAt,
	}
}

var baseChain = graphql.Chain{
	ChainID: 8453,
	Slug:    "base",
	Name:    "Base",
}

// TestFormatPostMessage_CreatePostOnly covers a freshly created post (rev 1,
// no amendments, no victims). This is acceptance criterion 1 + 3 + 4.
func TestFormatPostMessage_CreatePostOnly(t *testing.T) {
	note := v2Note(
		"Butter Bridge V3.1 drained via reentrancy",
		"base",
		"0x31e5e5f2a8b3c4d5e6f7a8b9c0d1e2f3a4b5c6d7e8f9a0b1c2d3e4f5a6b7c8d9e0664a",
		"@rektreporter",
	)
	p := makePost(struct {
		title       string
		note        string
		actionCount int
		attackers   []string
		victims     []string
		chain       graphql.Chain
		updatedAt   string
	}{
		title:       "Butter Bridge V3.1",
		note:        note,
		actionCount: 1, // one createPost action → rev 1
		attackers:   []string{"0x4059e47b062D9F959e2059b48cD6dB264EF5279F"},
		victims:     nil,
		chain:       baseChain,
		updatedAt:   "2026-05-20T14:00:00Z",
	})

	msg := telegram.FormatPostMessage(p)

	// Must contain the title and revision
	assertContains(t, msg, "Butter Bridge V3.1", "title")
	assertContains(t, msg, "rev 1", "rev derived from action count")

	// Summary from parsed note
	assertContains(t, msg, "Butter Bridge V3.1 drained via reentrancy", "summary from note")

	// Attacker section present with abbreviated address
	assertContains(t, msg, "Attackers:", "attackers section")
	assertContains(t, msg, "0x4059", "attacker addr prefix")
	assertContains(t, msg, "279F", "attacker addr suffix")

	// No victims section for a post without victims
	assertAbsent(t, msg, "Victims:", "victims section must be absent when no victims")

	// Source attribution
	assertContains(t, msg, "@rektreporter", "source attribution")

	// No score/confidence
	assertAbsent(t, msg, "score", "no score")
	assertAbsent(t, msg, "confidence", "no confidence")
	assertAbsent(t, msg, "✓", "no confirmation count")
	assertAbsent(t, msg, "✗", "no disconfirmation count")
}

// TestFormatPostMessage_AmendedPost covers an amended post (rev > 1).
// Acceptance criterion 3: rev derived from action count.
func TestFormatPostMessage_AmendedPost(t *testing.T) {
	note := v2Note(
		"MapProtocol bridge exploited for 3M USDC",
		"ethereum, arbitrum",
		"0xaaaa1111bbbb2222cccc3333dddd4444eeee5555ffff6666aaaa1111bbbb2222cc",
		"@rektreporter",
	)
	p := makePost(struct {
		title       string
		note        string
		actionCount int
		attackers   []string
		victims     []string
		chain       graphql.Chain
		updatedAt   string
	}{
		title:       "MapProtocol Bridge Exploit",
		note:        note,
		actionCount: 3, // createPost + 2 amendments → rev 3
		attackers:   []string{"0xDEADBEEF0000000000000000000000000000DEAD"},
		victims:     nil,
		chain:       graphql.Chain{ChainID: 1, Slug: "ethereum", Name: "Ethereum"},
		updatedAt:   "2026-05-21T10:00:00Z",
	})

	msg := telegram.FormatPostMessage(p)

	assertContains(t, msg, "rev 3", "rev derived from action count of 3")
	assertContains(t, msg, "MapProtocol Bridge Exploit", "title")
	assertContains(t, msg, "MapProtocol bridge exploited for 3M USDC", "summary")
	// Multi-chain attacked set from note
	assertContains(t, msg, "ethereum", "attacked chain from note")
	assertContains(t, msg, "arbitrum", "attacked chain from note")
}

// TestFormatPostMessage_WithVictims covers a post that lists victims.
// Acceptance criterion 5: victims render in their own section.
func TestFormatPostMessage_WithVictims(t *testing.T) {
	note := v2Note(
		"Hack with both attackers and victims",
		"base",
		"0xbbbb1111cccc2222dddd3333eeee4444ffff5555aaaa6666bbbb1111cccc222200",
		"@analyst",
	)
	p := makePost(struct {
		title       string
		note        string
		actionCount int
		attackers   []string
		victims     []string
		chain       graphql.Chain
		updatedAt   string
	}{
		title:       "FooProtocol Hack",
		note:        note,
		actionCount: 1,
		attackers:   []string{"0x1111111111111111111111111111111111111111"},
		victims:     []string{"0x2222222222222222222222222222222222222222", "0x3333333333333333333333333333333333333333"},
		chain:       baseChain,
		updatedAt:   "2026-05-21T11:00:00Z",
	})

	msg := telegram.FormatPostMessage(p)

	assertContains(t, msg, "Attackers:", "attackers section")
	assertContains(t, msg, "Victims:", "victims section must be present when victims exist")
	// Both victim addresses must appear abbreviated
	assertContains(t, msg, "0x2222", "first victim prefix")
	assertContains(t, msg, "0x3333", "second victim prefix")
}

// TestFormatPostMessage_WithoutVictims ensures the Victims section is omitted.
// Acceptance criterion 5: victims section omitted when absent.
func TestFormatPostMessage_WithoutVictims(t *testing.T) {
	note := v2Note("No victims here", "base", "0xcc", "@anon")
	p := makePost(struct {
		title       string
		note        string
		actionCount int
		attackers   []string
		victims     []string
		chain       graphql.Chain
		updatedAt   string
	}{
		title:       "Victimless",
		note:        note,
		actionCount: 1,
		attackers:   []string{"0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"},
		victims:     nil,
		chain:       baseChain,
		updatedAt:   "2026-05-21T12:00:00Z",
	})

	msg := telegram.FormatPostMessage(p)
	assertAbsent(t, msg, "Victims:", "victims section must be absent")
}

// TestFormatPostMessage_ExplorerLinks checks that attacker addresses and tx
// hashes render with HTML anchor tags pointing to a block explorer.
func TestFormatPostMessage_ExplorerLinks(t *testing.T) {
	txHash := "0x31e5e5f2a8b3c4d5e6f7a8b9c0d1e2f3a4b5c6d7e8f9a0b1c2d3e4f5a6b7664a"
	note := v2Note("Test summary", "ethereum", txHash, "@src")
	p := makePost(struct {
		title       string
		note        string
		actionCount int
		attackers   []string
		victims     []string
		chain       graphql.Chain
		updatedAt   string
	}{
		title:       "LinkTest",
		note:        note,
		actionCount: 1,
		attackers:   []string{"0x4059e47b062D9F959e2059b48cD6dB264EF5279F"},
		victims:     nil,
		chain:       graphql.Chain{ChainID: 1, Slug: "ethereum", Name: "Ethereum"},
		updatedAt:   "2026-05-21T13:00:00Z",
	})

	msg := telegram.FormatPostMessage(p)

	// Attacker address should be an HTML link
	assertContains(t, msg, `href=`, "attacker should be an HTML anchor")
	// The tx hash should appear abbreviated: first 6 chars + last 4
	assertContains(t, msg, "0x31e5", "tx hash prefix")
	assertContains(t, msg, "664a", "tx hash suffix")
}

// TestFormatPostMessage_HTMLEscaping checks that injected HTML in title or
// summary is escaped before being written into the message body.
func TestFormatPostMessage_HTMLEscaping(t *testing.T) {
	note := v2Note("<script>alert(1)</script>", "base", "0xaa", "@anon")
	p := makePost(struct {
		title       string
		note        string
		actionCount int
		attackers   []string
		victims     []string
		chain       graphql.Chain
		updatedAt   string
	}{
		title:       "<b>bold title</b>",
		note:        note,
		actionCount: 1,
		attackers:   []string{"0x4059e47b062D9F959e2059b48cD6dB264EF5279F"},
		victims:     nil,
		chain:       baseChain,
		updatedAt:   "2026-05-21T14:00:00Z",
	})

	msg := telegram.FormatPostMessage(p)

	// Must not contain raw injected tags
	assertAbsent(t, msg, "<script>", "raw script tag must be escaped")
	assertAbsent(t, msg, "<b>bold title</b>", "raw bold title must be escaped")
	// Must contain escaped form
	assertContains(t, msg, "&lt;script&gt;", "escaped script tag in summary")
}

// --- helpers ---

func assertContains(t *testing.T, haystack, needle, label string) {
	t.Helper()
	if !strings.Contains(haystack, needle) {
		t.Errorf("[%s] expected %q in message, not found.\nMessage:\n%s", label, needle, haystack)
	}
}

func assertAbsent(t *testing.T, haystack, needle, label string) {
	t.Helper()
	if strings.Contains(haystack, needle) {
		t.Errorf("[%s] did not expect %q in message, but found it.\nMessage:\n%s", label, needle, haystack)
	}
}
