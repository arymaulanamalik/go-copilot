package main

import (
	"fmt"

	"github.com/arymaulanamalik/go-copilot/2/service"
)

func main() {

	// 1. create a new service
	// 2. create a new handler
	// 3. create a new library
	// 4. view list library

	// 1. create a new service
	serv := service.NewService()
	// or
	// service := service.Service{}
	// or
	// service := service.ServiceInterface{}

	// 2. create a new handler
	handler := service.NewHandler(*serv)

	// 3. create a new library
	request := service.Library{
		ID:   1,
		Name: "Book 1",
	}
	err := handler.CreateLibrary(request)
	if err != nil {
		fmt.Println(err)
	}

	// 4. view list library
	library := handler.ViewLibrary()
	fmt.Println(library)

	fmt.Println("vim-go")
}
