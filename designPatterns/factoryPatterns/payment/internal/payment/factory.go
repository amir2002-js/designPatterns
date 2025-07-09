package payment

type FactoryPaymentAccount interface {
	CreatePayment(balance float64) PaymentMethod
}

type FactoryPaPal struct{}

func (f *FactoryPaPal) CreatePayment(balance float64) PaymentMethod {
	return CreateAccountPayPal(balance)
}

type FactoryCreditCard struct{}

func (f *FactoryCreditCard) CreatePayment(balance float64) PaymentMethod {
	return CreateAccountCreditCard(balance)
}
