package main

import (
    "go-micro.dev/v5"
)

func main() {
    // create the service
    service := micro.NewService(
        micro.Name("helloworld"),
        micro.Address(":8080"),
    )

    // register handlers
    service.Handle(new(Say))
    service.Handle(new(Sale))

    // run the service
    service.Run()
}
