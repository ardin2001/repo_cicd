package main

import (
	"unit_testing/routes"
)

func main() {

	e := routes.New()

	e.Logger.Fatal(e.Start(":8000"))
}
