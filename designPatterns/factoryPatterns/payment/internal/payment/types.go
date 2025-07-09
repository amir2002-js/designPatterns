package payment

import (
	"fmt"
)

type account struct {
	balance float64
}

type PayPal struct {
	account
}

func (p *PayPal) Pay(amount float64) error {
	if amount <= p.balance {
		p.balance = p.balance - amount
		return nil
	}
	return fmt.Errorf("not enough balance")
}
func (p *PayPal) Deposit(amount float64) {
	p.balance = p.balance + amount
}
func (p PayPal) ShowBalance() {
	fmt.Println(p.balance)
}
func CreateAccountPayPal(balance float64) PaymentMethod {
	return &PayPal{account{balance}}
}

type CreditCard struct {
	account
}

func (c *CreditCard) Pay(amount float64) error {
	if amount <= c.balance {
		c.balance = c.balance - amount
		return nil
	}
	return fmt.Errorf("not enough balance")
}
func (c *CreditCard) Deposit(amount float64) {
	c.balance = c.balance + amount
}
func (c CreditCard) ShowBalance() {
	fmt.Println(c.balance)
}
func CreateAccountCreditCard(balance float64) PaymentMethod {
	return &CreditCard{account{balance}}
}
