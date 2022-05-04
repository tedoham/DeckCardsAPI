package decorators

var values = map[string]string{
	"A": "ACE",
	"2": "2",
	"3": "3",
	"4": "4",
	"5": "5",
	"6": "6",
	"7": "7",
	"8": "8",
	"9": "9",
	"T": "10",
	"J": "JACK",
	"Q": "QUEEN",
	"K": "KING",
}

var suits = map[string]string{
	"S": "SPADES",
	"D": "DIAMONDS",
	"C": "CLUBS",
	"H": "HEARTS",
}

type CardDecorator struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

func NewCardDecorator(card string) CardDecorator {
	return CardDecorator{
		Value: values[card[:1]],
		Suit:  suits[card[1:]],
		Code:  card,
	}
}

func NewCardsDecorator(cards []string) []CardDecorator {
	var decoratedCards = make([]CardDecorator, len(cards))
	for i, card := range cards {
		decoratedCards[i] = NewCardDecorator(card)
	}
	return decoratedCards
}
