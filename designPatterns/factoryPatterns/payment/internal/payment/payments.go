package paymentSystem

type PaymentMethod interface {
	Pay(amount float64) error
	Deposit(amount float64)
	ShowBalance()
}
