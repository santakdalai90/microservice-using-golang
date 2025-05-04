package payment

type Payment struct {
    Id            string  `json:"id"`
    Receiver      string  `json:"receiver"`
    Amount        float64 `json:"amount"`
    Currency      string  `json:"currency"`
    Description   string  `json:"description"`
    PaymentMethod string  `json:"payment_method"`
    PaymentTime   string  `json:"payment_time"`
    Status        string  `json:"status"`
}