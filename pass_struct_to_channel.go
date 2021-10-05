package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type ResponseData struct {
	Data  []byte
	Error string
}

func Deploy(applianceProviderType, configData, outputFile string) ([]byte, string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*80))
	defer cancel()

	osSignalChannel := make(chan os.Signal, 1)
	createResponseChannel := make(chan ResponseData, 1)

	signal.Notify(osSignalChannel, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		kc, err := []byte("hello"), "new"
		if err == "new" {
			createResponseChannel <- ResponseData{Data: kc, Error: err}
		}
		createResponseChannel <- ResponseData{Data: []byte(""), Error: err}
	}()

	select {
	case <-ctx.Done():
		// fix error after rebase
		return []byte(""), "timed out"
	case createResponse := <-createResponseChannel:
		fmt.Println("printing inside operations")
		fmt.Println(createResponse)
		return createResponse.Data, createResponse.Error
	case returnSignal := <-osSignalChannel:
		return []byte(""), returnSignal.String()
	}
}

func main() {
	ans, err := Deploy("", "", "")
	fmt.Println("outside")
	fmt.Println(string(ans), err)
}
