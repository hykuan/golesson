package main

import (
	"github.com/hykuan/golesson/route"
)

func main() {
	r := router.Router

	router.SetRouter()

	r.Run(":1234")
}
