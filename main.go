package main

import (
	"github.com/mirzaRakha28/ninja_to/db"
	"github.com/mirzaRakha28/ninja_to/routes"
)

func main() {
	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
