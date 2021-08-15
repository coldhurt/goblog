package main

import (
	_ "github.com/coldhurt/goblog/config"
	"github.com/coldhurt/goblog/router"
)

func main() {
	r := router.InitRouter()
	r.Run(":8080")
}
