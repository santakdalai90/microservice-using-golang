package main

import (
        "context"
        "fmt"

        "go-micro.dev/v5"
        "microservice-using-golang/service_discovery/model"
)

type VegetableMarket struct{}

func (g *VegetableMarket) UpdatePrice(ctx context.Context, req *model.PriceUpdateRequest, rsp *model.PriceUpdateResponse) error {
        rsp.Message = fmt.Sprintf("Price of %s updated to %f", req.Name, req.Price)
        return nil
}

func main() {
        // Create a new service
        service := micro.NewService(
                micro.Name("VegetableMarket"),
                micro.Address(":8080")
        )

        // Initialize the service
        service.Init()

        // Register handler
        micro.RegisterHandler(service.Server(), new(VegetableMarket))

        // Start the service
        if err := service.Run(); err != nil {
                fmt.Println(err)
        }
}
