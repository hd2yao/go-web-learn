package main

import (
	"fmt"
	"gee"
)

func main() {
	r := gee.New()
	r.GET("/", func(context *gee.Context) {
		fmt.Println("test /")
	})

	r.Run(":8000")
}
