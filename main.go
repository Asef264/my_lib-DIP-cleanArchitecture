package main

import (
	"my_lib/handller"
	"my_lib/repository"
)

func main() {
	repository.Connect()
	handller.Run()
}
