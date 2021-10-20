package main

import (
	"ginLearn/route"
)

func initialize() {

}

func main() {
	initialize()

	r := route.GetRoute()

	// TODO add middleware
	// TODO Custom Recovery behavior

	r.Run("localhost:9000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
