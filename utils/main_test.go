package utils

import "testing"

func TestContains(t *testing.T) {
	test := "AS"
	slice := []string{test, "TH"}

	got := Contains(slice, test)
	if got != true {
		t.Errorf("Contains(%v, %q) == true, got %v", slice, test, got)
	}

	test = "WRONG"
	got = Contains(slice, test)
	if got != false {
		t.Errorf("Contains(%v, %q) == false, got %v", slice, test, got)
	}
}
