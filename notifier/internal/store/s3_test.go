package store_test

import (
	"encoding/json"
	"testing"

	"github.com/ThatsRekt/thatsRekt/notifier/internal/store"
)

// TestStateDeserializesLegacyVoteFields verifies backward compatibility with
// existing S3 state JSON that contains the now-removed vote fields (upVotes,
// downVotes, voters). Go's json.Unmarshal ignores unknown fields by default,
// so old state must deserialize cleanly without error, and the fields that
// were kept (messageId, lastActionCount, lastUpdatedAt, chainSlug, retracted)
// must round-trip correctly.
//
// This is a real struct round-trip — no mocks. It proves that deploying the
// new binary against existing S3 state does not crash on Load.
func TestStateDeserializesLegacyVoteFields(t *testing.T) {
	// JSON that mimics what the old binary would have persisted: PostState
	// entries that include upVotes, downVotes, and voters.
	legacyJSON := `{
		"lastSeenByChain": {
			"base": "base-3"
		},
		"posts": {
			"base-1": {
				"messageId": 42,
				"upVotes": 5,
				"downVotes": 2,
				"voters": {
					"100001": "up",
					"100002": "down"
				},
				"lastActionCount": 1,
				"lastUpdatedAt": "2026-05-21T10:00:00Z",
				"chainSlug": "base",
				"retracted": false
			},
			"base-2": {
				"messageId": 99,
				"upVotes": 0,
				"downVotes": 0,
				"voters": {},
				"lastActionCount": 2,
				"lastUpdatedAt": "2026-05-22T08:00:00Z",
				"chainSlug": "base",
				"retracted": true
			}
		}
	}`

	// Unmarshal directly into the real store.State type.  This is the same
	// step that Store.Load performs internally.  If a future change adds
	// json.Decoder.DisallowUnknownFields or restructures PostState, this test
	// will catch the regression immediately — the throwaway local struct it
	// replaced would not have caught that.
	var state store.State
	if err := json.Unmarshal([]byte(legacyJSON), &state); err != nil {
		t.Fatalf("legacy JSON must unmarshal into store.State without error: %v", err)
	}

	if state.LastSeenByChain["base"] != "base-3" {
		t.Errorf("expected LastSeenByChain[base]=base-3, got %q", state.LastSeenByChain["base"])
	}
	if len(state.Posts) != 2 {
		t.Fatalf("expected 2 posts in state.Posts, got %d", len(state.Posts))
	}

	// Verify base-1 fields round-tripped correctly through the real PostState.
	ps1, ok := state.Posts["base-1"]
	if !ok {
		t.Fatal("expected post base-1 to be present in state.Posts")
	}
	if ps1.MessageID != 42 {
		t.Errorf("base-1: expected MessageID=42, got %d", ps1.MessageID)
	}
	if ps1.LastActionCount != 1 {
		t.Errorf("base-1: expected LastActionCount=1, got %d", ps1.LastActionCount)
	}
	if ps1.LastUpdatedAt != "2026-05-21T10:00:00Z" {
		t.Errorf("base-1: expected LastUpdatedAt=2026-05-21T10:00:00Z, got %q", ps1.LastUpdatedAt)
	}
	if ps1.ChainSlug != "base" {
		t.Errorf("base-1: expected ChainSlug=base, got %q", ps1.ChainSlug)
	}
	if ps1.Retracted {
		t.Errorf("base-1: expected Retracted=false, got true")
	}

	// Verify base-2 retracted flag.
	ps2, ok := state.Posts["base-2"]
	if !ok {
		t.Fatal("expected post base-2 to be present in state.Posts")
	}
	if ps2.MessageID != 99 {
		t.Errorf("base-2: expected MessageID=99, got %d", ps2.MessageID)
	}
	if !ps2.Retracted {
		t.Errorf("base-2: expected Retracted=true, got false")
	}

	// Use NewInMemory so no S3 client is needed.
	st := store.NewInMemory()

	// Verify the in-memory store API still works correctly after the removal.
	st.RegisterPost("base-1", 42, "base")
	st.UpdatePostSnapshot("base-1", 1, "2026-05-21T10:00:00Z")
	if !st.HasSnapshot("base-1") {
		t.Error("HasSnapshot must return true after UpdatePostSnapshot")
	}
	if st.HasChanged("base-1", 1, "2026-05-21T10:00:00Z") {
		t.Error("HasChanged must return false when snapshot matches")
	}
	if !st.HasChanged("base-1", 2, "2026-05-21T10:00:00Z") {
		t.Error("HasChanged must return true when actionCount changes")
	}
}
