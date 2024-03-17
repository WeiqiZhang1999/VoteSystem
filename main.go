package main

import (
	"GinRanking/router"
)

func main() {
	r := router.Router()
	r.Run(":9999")
}
