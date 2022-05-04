package main

import "github.com/tedoham/DeckCardsAPI/app"

var a app.App

func main() {
	a.Initialize()
	a.Run(":8080")
}
