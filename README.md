Navigate to the project and run `go run main.go` to start the project and navigate to `http://localhost:8080`

Avalable API methods are:
- `POST /decks` to create deck. Optional params (`shuffled` and `cards`) are expected as json in the request body.
- for `cards` on body provide json `{"cards": ["AS","KD","AC","2C","KH"] }`
- for `shuffled` on body provide json `{"shuffled": true}`
- `GET /decks/{id}` to open deck.
- `POST /decks/{id}/draw` to draw card(s). Parameter `count` is expected as json in the request body.
- for `draw` on body provide json `{"count":1}`


Run tests with `go test ./...`