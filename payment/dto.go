package payment

type (
    CreatePaymentRequest struct {
        Receiver      string  `json:"receiver"`
        Amount        float64 `json:"amount"`
        Currency      string  `json:"currency"`
        Description   string  `json:"description"`
        PaymentMethod string  `json:"payment_method"`
    }

    CreatePaymentResponse struct {
        Ok bool `json:"ok"`
        UniqueID string `json:"unique_id"`
        Status string `json:"status"`
    }
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
    return json.NewEncoder(w).Encode(response)
}

func decodeCreatePaymentRequest(ctx context.Context, r *http.Request) (interface{}, error) {
    var req CreatePaymentRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        return nil, err
    }
    return req, nil
}

