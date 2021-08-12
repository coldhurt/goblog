package main

import (
	"github.com/coldhurt/goblog/router"
)

func main() {
	r := router.InitRouter()
	r.Run(":8080")
}
