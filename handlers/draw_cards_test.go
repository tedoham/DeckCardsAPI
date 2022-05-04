package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/tedoham/DeckCardsAPI/decorators"
	"github.com/tedoham/DeckCardsAPI/repository"
)

func TestDrawCardsHandler(t *testing.T) {
	deck, err := repository.CreateDeck(false, []string{})
	if err != nil {
		t.Fatal(err)
	}

	count := 2
	url := fmt.Sprintf("/decks/%v/draw", deck.Id)
	var req *http.Request
	body := fmt.Sprintf(`{"count": %d}`, count)
	req, err = http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/decks/{id}/draw", DrawCardsHandler).Methods(http.MethodPost)
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	decorator := decorators.NewDrawCardsDecorator(deck, count)
	var expected []byte
	expected, err = json.Marshal(decorator)
	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), string(expected))
	}
}
