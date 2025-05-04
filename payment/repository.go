package payment

type repo struct {
    db *sql.DB
    logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) Repository {
    return &repo{
        db: db,
        logger: log.With(logger, "repo", "sql"),
    }
}

func (r *repo) CreatePayment(ctx context.Context, payment Payment) error {
    sql := `INSERT INTO payments (id, receiver, amount, currency, description, payment_method, payment_time, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
    _, err := r.db.ExecContext(ctx, sql, payment.Id, payment.Receiver, payment.Amount, payment.Currency, payment.Description, payment.PaymentMethod, payment.PaymentTime, payment.Status)
    if err != nil {
        return err
    }
    return nil
}

func (r *repo) GetPayment(ctx context.Context, id string) (Payment, error) {
    var payment Payment
    sql := `SELECT id, receiver, amount, currency, description, payment_method, payment_time, status FROM payments WHERE id = $1`
    row := r.db.QueryRowContext(ctx, sql, id)
    err := row.Scan(&payment.Id, &payment.Receiver, &payment.Amount, &payment.Currency, &payment.Description, &payment.PaymentMethod, &payment.PaymentTime, &payment.Status)
    if err != nil {
        return Payment{}, err
    }
    return payment, nil
}