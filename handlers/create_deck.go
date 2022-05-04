package handlers

import (
	"encoding/json"

	"net/http"

	"github.com/tedoham/DeckCardsAPI/decorators"
	"github.com/tedoham/DeckCardsAPI/repository"
)

type createDeckParams struct {
	Shuffled bool     `json:"shuffled,omitempty"`
	Cards    []string `json:"cards,omitempty"`
}

func CreateDeckHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var params createDeckParams
	err := decoder.Decode(&params)
	if err != nil {
		message := "This API endpoint should be called with 'shuffled' and 'cards' params"
		respondWithError(w, http.StatusBadRequest, message)
		return
	}

	deck, err := repository.CreateDeck(params.Shuffled, params.Cards)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	decorator := decorators.NewCreateDeckDecorator(deck)

	respondWithJSON(w, http.StatusCreated, decorator)
}
