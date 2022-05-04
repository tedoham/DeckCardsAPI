package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/tedoham/DeckCardsAPI/decorators"
	"github.com/tedoham/DeckCardsAPI/repository"
)

func TestCreateDeckHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/decks", strings.NewReader("{}"))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/decks", CreateDeckHandler).Methods(http.MethodPost)
	r.ServeHTTP(rr, req)

	wantStatus := http.StatusCreated
	if status := rr.Code; status != wantStatus {
		t.Errorf("handler returned wrong status code: got %v want %v", status, wantStatus)
	}

	deck := repository.GetDecks()[0]
	decorator := decorators.NewCreateDeckDecorator(deck)
	var expected []byte
	expected, err = json.Marshal(decorator)
	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), string(expected))
	}
}
