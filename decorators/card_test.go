package decorators

import "testing"

func TestNewCardDecorator(t *testing.T) {
	value := "T"
	suit := "D"
	code := value + suit
	got := NewCardDecorator(code)

	if got.Value != values[value] {
		t.Errorf("NewCardsDecorator(%q).Value == %q, got %q", code, values[value], got.Value)
	}

	if got.Suit != suits[suit] {
		t.Errorf("NewCardsDecorator(%q).Suit == %q, got %q", code, suits[suit], got.Suit)
	}

	if got.Code != code {
		t.Errorf("NewCardsDecorator(%q).Code == %q, got %q", code, code, got.Code)
	}
}
