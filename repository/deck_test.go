package repository

import (
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestCreateDeck(t *testing.T) {
	wantShuffled := true
	deck, _ := CreateDeck(wantShuffled, []string{})

	if deck.DrawCount != 0 {
		t.Errorf("CreateDeck().Shuffled == 0, got %d", deck.DrawCount)
	}

	if deck.Shuffled != wantShuffled {
		t.Errorf("CreateDeck().Shuffled == %v, got %v", wantShuffled, deck.Shuffled)
	}

	if _, err := uuid.FromString(deck.Id); err != nil {
		t.Errorf("CreateDeck().Id should be valid UUID, error %v", err)
	}
}

func TestOpenDeck(t *testing.T) {
	id := "id"
	want := &Deck{}
	decks[id] = want

	got, present := OpenDeck(id)
	if got != want {
		t.Errorf("OpenDeck(%q) == %v, got %v", id, want, got)
	}
	if !present {
		t.Errorf("OpenDeck(%q) shuld be present", id)
	}

	wrongId := "wrong"
	got, present = OpenDeck(wrongId)
	if got != nil {
		t.Errorf("OpenDeck(%q) == %v, got %v", wrongId, nil, got)
	}
	if present {
		t.Errorf("OpenDeck(%q) shuld be absent", wrongId)
	}
}
