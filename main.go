package main

import (
	"launchtime/api"
)

func main() {

	err := api.Run(":5000")
	if err != nil {
		panic(err)
	}
}
