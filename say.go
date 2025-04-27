package main

import "context"

type Request struct {
    Name string `json:"name"`
}

type Response struct {
    Message string `json:"message"`
}

type Say struct{}

func (h *Say) Hello(ctx context.Context, req *Request, rsp *Response) error {
    rsp.Message = "Hello " + req.Name
    return nil
}

func (h *Say) Bye(ctx context.Context, req *Request, rsp *Response) error {
    rsp.Message = "Bye " + req.Name
    return nil
}
