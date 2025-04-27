package main

import (
    "context"
    "fmt"
)

type TradeRequest struct {
    Price float64 `json:"price"`
}

type Sale struct{}

func (h *Sale) Sell(ctx context.Context, req *TradeRequest, rsp *Response) error {
    rsp.Message = fmt.Sprintf("Selling Price: %f", req.Price)
    return nil
}
