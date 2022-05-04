package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tedoham/DeckCardsAPI/decorators"
	"github.com/tedoham/DeckCardsAPI/repository"
)

type drawCardsParams struct {
	Count int `json:"count,omitempty"`
}

func DrawCardsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	decoder := json.NewDecoder(r.Body)
	var params drawCardsParams
	err := decoder.Decode(&params)
	if err != nil {
		message := "This API endpoint should be called with 'count' param"
		respondWithError(w, http.StatusBadRequest, message)
		return
	}

	fmt.Print(params)

	deck, drawErr := repository.DrawCards(id, params.Count)
	if drawErr != nil {
		respondWithError(w, drawErr.Status(), drawErr.Error())
		return
	}

	decorator := decorators.NewDrawCardsDecorator(deck, params.Count)
	respondWithJSON(w, http.StatusOK, decorator)
}
