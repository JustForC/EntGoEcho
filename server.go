package main

import (
	"CompanyAPI/route"
)

func main() {
	e := route.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
