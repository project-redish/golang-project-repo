package main

import (
	"fmt"
	handle "go-initializer/handler"
	"go-initializer/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	cleanUp := make(chan int, 1)

	webServer := server.Create(cleanUp)
	webServer.RegisterRoute()
	go webServer.Run(":8080")

	//this is dynamic part every cli will get registered through this TODO: create this kind of feature for rest also. for future releases
	go handle.RegisterCli()

	stopChan := make(chan os.Signal, 2)

	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	{
		fmt.Println(<-stopChan)
		fmt.Println("exiting ")

	}

}
