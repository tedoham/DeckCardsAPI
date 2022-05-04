package repository

import (
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"
	"sync"
)

type Deck struct {
	Id       string
	Shuffled bool
	Cards
	DrawCount int
}

// in memory repository
// can be changed to database storage if needed
var decks = make(map[string]*Deck)
var mux = sync.Mutex{}

func GetDecks() (slice []*Deck) {
	slice = make([]*Deck, 0, len(decks))
	for _, value := range decks {
		slice = append(slice, value)
	}
	return
}

func CreateDeck(shuffled bool, cards []string) (*Deck, error) {
	if len(cards) != 0 && !ValidateCards(cards) {
		return nil, errors.New("invalid cards list")
	}

	id := fmt.Sprint(uuid.NewV4())

	deck := Deck{
		Id:       id,
		Shuffled: shuffled,
		Cards:    CreateCards(shuffled, cards),
	}

	decks[id] = &deck

	return &deck, nil
}

func OpenDeck(id string) (*Deck, bool) {
	deck, present := decks[id]
	return deck, present
}

type drawError struct {
	status int    // http status
	err    string // error description
}

func (e *drawError) Error() string {
	return e.err
}

func (e *drawError) Status() int {
	return e.status
}

func DrawCards(id string, count int) (*Deck, *drawError) {
	deck, present := OpenDeck(id)
	if !present {
		return deck, &drawError{http.StatusNotFound, "Deck not found"}
	}

	if count+deck.DrawCount > len(deck.Cards) {
		return deck, &drawError{http.StatusBadRequest, "Not enough cards are remaining"}
	}

	mux.Lock()
	deck.DrawCount += count
	mux.Unlock()

	return deck, nil
}
