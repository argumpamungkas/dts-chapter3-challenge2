package main

import (
	"DTS/Chapter-3/chapter3-challenge2/repo"
	"DTS/Chapter-3/chapter3-challenge2/router"
)

func main() {

	repo.StartDB()

	router.StartServer().Run(":3000")

}
