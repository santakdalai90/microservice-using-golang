package account

import (
    "context"

    "github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
    CreatePayment endpoint.Endpoint
    GetPayment    endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
    return Endpoints{
        CreatePayment: makeCreatePaymentEndpoint(s),
        GetPayment:     makeGetPaymentEndpoint(s),
    }
}

func makeCreatePaymentEndpoint(s Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(CreatePaymentRequest)
        ok, err := s.CreatePayment(ctx, req.Receiver, req.Amount, req.Currency, req.Description, req.PaymentMethod, req.PaymentTime)
        return CreatePaymentResponse{Ok: ok, UniqueID: req.Id, Status: req.Status)
    }
}

