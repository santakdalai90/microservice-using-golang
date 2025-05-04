package payment

import "context"

type Service interface {
	CreatePayment(ctx context.Context, payment Payment) error
	GetPayment(ctx context.Context, id string) (Payment, error)
	// ListPayments(ctx context.Context) ([]Payment, error)
	// UpdatePayment(ctx context.Context, id string, payment Payment) error
}

type service struct {
    repo Repository
    logger log.Logger
}

func NewService(repo Repository, logger log.Logger) Service {
    return &service{
        repo: repo,
        logger: logger,
    }
}

func (s *service) CreatePayment(ctx context.Context, payment Payment) error {
    payment.Id = uuid.New().String()
    payment.Status = "success"
    payment.PaymentTime = time.Now().Format(time.RFC3339)
    return s.repo.CreatePayment(ctx, payment)
}