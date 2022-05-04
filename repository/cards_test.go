package repository

import (
	"testing"
)

func TestValidateCards(t *testing.T) {
	input := []string{"AC", "TS"}
	got := ValidateCards(input)
	if got == false {
		t.Errorf("ValidateCards(%v) == true, got %v", input, got)
	}

	input = []string{"WRONG"}
	got = ValidateCards(input)
	if got == true {
		t.Errorf("ValidateCards(%v) == false, got %v", input, got)
	}
}

func TestCreateCards(t *testing.T) {
	got := CreateCards(false, []string{})
	if len(got) != 52 {
		t.Errorf("len(CreateCards(false, []string{})) == 52, got %v", len(got))
	}

	want := []string{"AC", "4H"}
	got = CreateCards(false, want)
	if len(want) != len(got) || want[0] != got[0] || want[1] != got[1] {
		t.Errorf("CreateCards(false, %v) == %v, got %v", want, want, got)
	}
}
