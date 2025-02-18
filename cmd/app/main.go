package main

import "github.com/hell077/DiabetesHealthBot/db"

var (
	err error
)

func main() {
	err = db.Migrate()
	if err != nil {
		panic(err)
	}

}
