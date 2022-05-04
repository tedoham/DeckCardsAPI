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

func TestOpenDeckHandler(t *testing.T) {
	deck, err := repository.CreateDeck(false, []string{})
	if err != nil {
		t.Fatal(err)
	}

	url := fmt.Sprintf("/decks/%v", deck.Id)
	var req *http.Request
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/decks/{id}", OpenDeckHandler).Methods(http.MethodGet)
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	decorator := decorators.NewOpenDeckDecorator(deck)
	var expected []byte
	expected, err = json.Marshal(decorator)
	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), string(expected))
	}
}
