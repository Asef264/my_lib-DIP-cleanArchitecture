/*
((((Khaled Moazedi))))
this snipet of codes is a sample for using clean architecture in golang
obviusly it can be more complicated than this but this is for well undrestandings
*/
package main

import (
	"my_lib/handller"
	"my_lib/repository"
)

func main() {
	repository.Connect()
	handller.Run()
}
