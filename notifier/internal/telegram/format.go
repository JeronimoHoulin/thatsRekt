package telegram

import (
	"fmt"
	"strings"

	"github.com/ThatsRekt/thatsRekt/notifier/internal/graphql"
	"github.com/ThatsRekt/thatsRekt/notifier/internal/note"
)

// FormatPostMessage renders the uniform v2 Telegram message for any on-chain
// post — createPost or amendment — built purely from on-chain data.
//
// Target format:
//
//	🚨 HACK VERIFIED
//	<title> on <chain>
//	updated · rev <N>
//
//	<summary>
//
//	Attackers:
//	  <addr> (<explorer link>)
//
//	[Victims:
//	  <addr> (<explorer link>)]
//
//	Tx:
//	  <txHash> (<explorer link>)
//	  ...
//
//	Source: @<handle>
//
// Rules:
//   - All content is sourced from the self-describing on-chain note. The
//     formatter parses the note for summary, attacked-chain set, exploit tx
//     hashes, and sources.
//   - `rev N` is derived from p.ActionCount (1 createPost + N-1 amendments).
//     If ActionCount is 0 (indexer not yet upgraded), falls back to rev 1.
//   - No confidence/score is shown.
//   - Victims section is rendered the same way as Attackers when present;
//     omitted entirely when p.Victims is empty.
//   - Addresses and tx hashes are abbreviated (first 6 + last 4 chars) with
//     an HTML anchor pointing to the appropriate block explorer.
//   - All user-supplied text is HTML-escaped before insertion.
func FormatPostMessage(p graphql.Post) string {
	parsed := note.ParseNote(p.Note)

	chainName := chainDisplayName(p.Chain)

	title := strings.TrimSpace(p.Title)
	if title == "" {
		title = "(untitled alert)"
	}

	rev := p.ActionCount
	if rev < 1 {
		rev = 1
	}

	// The "on <chains>" line uses the full attacked-chain set from the
	// self-describing note (which may span multiple chains). Fall back to
	// the posting chain when the note carries no chains.
	chains := strings.Join(parsed.AttackedChains, ", ")
	if chains == "" {
		chains = chainName
	}

	var b strings.Builder

	// Header
	fmt.Fprintf(&b, "🚨 <b>HACK VERIFIED</b>\n")
	fmt.Fprintf(&b, "%s on %s\n", html(title), html(chains))
	fmt.Fprintf(&b, "updated · rev %d\n", rev)

	// Summary from parsed note
	if summary := strings.TrimSpace(parsed.Summary); summary != "" {
		fmt.Fprintf(&b, "\n%s\n", html(summary))
	}

	// Attackers
	if len(p.Attackers) > 0 {
		fmt.Fprintf(&b, "\nAttackers:\n")
		for _, addr := range p.Attackers {
			link := explorerAddrURL(p.Chain, addr)
			fmt.Fprintf(&b, "  %s (%s)\n", addrAbbrev(addr), explorerLink(link, addrAbbrev(addr)))
		}
	}

	// Victims (only when present)
	if len(p.Victims) > 0 {
		fmt.Fprintf(&b, "\nVictims:\n")
		for _, addr := range p.Victims {
			link := explorerAddrURL(p.Chain, addr)
			fmt.Fprintf(&b, "  %s (%s)\n", addrAbbrev(addr), explorerLink(link, addrAbbrev(addr)))
		}
	}

	// Exploit tx hashes from parsed note
	if len(parsed.ExploitTxHashes) > 0 {
		fmt.Fprintf(&b, "\nTx:\n")
		for _, txHash := range parsed.ExploitTxHashes {
			link := explorerTxURL(p.Chain, txHash)
			fmt.Fprintf(&b, "  %s (%s)\n", txAbbrev(txHash), explorerLink(link, txAbbrev(txHash)))
		}
	}

	// Sources from parsed note
	if len(parsed.Sources) > 0 {
		fmt.Fprintf(&b, "\nSource: %s\n", html(strings.Join(parsed.Sources, ", ")))
	}

	return strings.TrimRight(b.String(), "\n")
}

// VoteKeyboard builds the cosmetic ✓/✗ inline keyboard. The callback_data
// payload is `vote:{up|down}:{postId}` so the press handler can identify
// which post + direction without needing a separate lookup table.
//
// These counts are TELEGRAM-side only — they do NOT affect the on-chain
// confirm/disconfirm state. The OG-card preview shows the on-chain numbers
// for canonical truth; the buttons are a low-effort engagement signal for
// chat readers who don't have a wallet handy.
func VoteKeyboard(postID string, up, down int) *InlineKeyboardMarkup {
	return &InlineKeyboardMarkup{
		InlineKeyboard: [][]InlineKeyboardButton{
			{
				{Text: fmt.Sprintf("✓  %d", up), CallbackData: "vote:up:" + postID},
				{Text: fmt.Sprintf("✗  %d", down), CallbackData: "vote:down:" + postID},
			},
		},
	}
}

// --- explorer URL helpers ---

// explorerAddrURL builds the block explorer URL for an address given the
// post's chain. Returns an empty string for unknown chains, in which case
// the link is omitted and only the abbreviated address is shown.
func explorerAddrURL(c graphql.Chain, addr string) string {
	base := explorerBase(c)
	if base == "" {
		return ""
	}
	return base + "/address/" + addr
}

// explorerTxURL builds the block explorer URL for a transaction hash.
func explorerTxURL(c graphql.Chain, txHash string) string {
	base := explorerBase(c)
	if base == "" {
		return ""
	}
	return base + "/tx/" + txHash
}

// explorerBase returns the block explorer base URL for a chain. Returns an
// empty string for unknown chains.
func explorerBase(c graphql.Chain) string {
	switch c.ChainID {
	case 1:
		return "https://etherscan.io"
	case 10:
		return "https://optimistic.etherscan.io"
	case 56:
		return "https://bscscan.com"
	case 100:
		return "https://gnosisscan.io"
	case 137:
		return "https://polygonscan.com"
	case 8453:
		return "https://basescan.org"
	case 42161:
		return "https://arbiscan.io"
	case 43114:
		return "https://snowtrace.io"
	default:
		return ""
	}
}

// explorerLink wraps label in an HTML anchor when url is non-empty;
// returns the label unchanged otherwise.
func explorerLink(url, label string) string {
	if url == "" {
		return label
	}
	return fmt.Sprintf(`<a href="%s">%s</a>`, url, label)
}

// --- abbreviation helpers ---

// addrAbbrev renders a hex address or tx hash as `0x1234…abcd`
// (first 6 chars + last 4). Inputs shorter than 10 chars are returned as-is.
func addrAbbrev(addr string) string {
	if len(addr) < 10 {
		return addr
	}
	return addr[:6] + "…" + addr[len(addr)-4:]
}

// txAbbrev is an alias for addrAbbrev — tx hashes use the same abbreviation.
func txAbbrev(txHash string) string {
	return addrAbbrev(txHash)
}

// chainDisplayName returns a human-readable chain name, falling back to the
// uppercased slug when Name is empty.
func chainDisplayName(c graphql.Chain) string {
	if c.Name != "" {
		return c.Name
	}
	return strings.ToUpper(c.Slug)
}

// html escapes the four characters Telegram's HTML parse mode treats
// specially: `<`, `>`, `&`, `"`. (The Bot API parses HTML strictly enough
// that an un-escaped `<` in a poster's note will fail the entire message
// with `can't parse entities`.)
func html(s string) string {
	r := strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
		"\"", "&quot;",
	)
	return r.Replace(s)
}
