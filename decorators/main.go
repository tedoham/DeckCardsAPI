package decorators

import (
	"github.com/tedoham/DeckCardsAPI/repository"
)

type CreateDeckDecorator struct {
	DeckId    string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

func NewCreateDeckDecorator(deck *repository.Deck) (decorator CreateDeckDecorator) {
	decorator.DeckId = deck.Id
	decorator.Shuffled = deck.Shuffled
	decorator.Remaining = len(deck.Cards) - deck.DrawCount
	return
}

type OpenDeckDecorator struct {
	DeckId    string          `json:"deck_id"`
	Shuffled  bool            `json:"shuffled"`
	Remaining int             `json:"remaining"`
	Cards     []CardDecorator `json:"cards"`
}

func NewOpenDeckDecorator(deck *repository.Deck) (decorator OpenDeckDecorator) {
	decorator.DeckId = deck.Id
	decorator.Shuffled = deck.Shuffled
	decorator.Remaining = len(deck.Cards) - deck.DrawCount
	decorator.Cards = NewCardsDecorator(deck.Cards[deck.DrawCount:])
	return
}

type DrawCardsDecorator struct {
	Cards []CardDecorator `json:"cards"`
}

func NewDrawCardsDecorator(deck *repository.Deck, count int) (decorator DrawCardsDecorator) {
	decorator.Cards = NewCardsDecorator(deck.Cards[deck.DrawCount-count : deck.DrawCount])
	return
}
