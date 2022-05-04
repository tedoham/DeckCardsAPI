package decorators

import (
	"github.com/tedoham/DeckCardsAPI/repository"
	"testing"
)

func TestNewCreateDeckDecorator(t *testing.T) {
	id := "id"
	shuffled := true
	drawCount := 1
	cards := []string{"AC", "1C", "2C"}
	deck := repository.Deck{
		Id: id,
		Shuffled: shuffled,
		DrawCount: drawCount,
		Cards: cards,
	}

	got := NewCreateDeckDecorator(&deck)

	if got.DeckId != id {
		t.Errorf("NewCreateDeckDecorator(&deck).DeckId == %q, got %q", id, got.DeckId)
	}

	if got.Shuffled != shuffled {
		t.Errorf("NewCreateDeckDecorator(&deck).Shuffled == %v, got %v", shuffled, got.Shuffled)
	}

	if remaining := len(cards) - drawCount; got.Remaining != remaining {
		t.Errorf("NewCreateDeckDecorator(&deck).Remaining == %d, got %d", remaining, got.Remaining)
	}
}
