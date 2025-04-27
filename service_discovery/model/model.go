package model

type PriceUpdateRequest struct{
    Price float64 `json:"price"`
    Name string `json:"name"`
}

type PriceUpdateResponse struct
    Message string `json:"message"`
}