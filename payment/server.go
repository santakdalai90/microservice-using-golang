package payment

import (
	"context"
	"net/http"

    "github.com/gorilla/mux"
    httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("/payments").Handler(
		httpTransport.NewServer(
			endpoints.CreatePayment,
			decodeCreatePaymentRequest,
			encodeResponse,
		),
	)

	r.Methods("GET").Path("/payments/{id}").Handler(
		httpTransport.NewServer(
			endpoints.GetPayment,
			decodeGetPaymentRequest,
			encodeResponse,
		),
	)

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		},
	)
}
