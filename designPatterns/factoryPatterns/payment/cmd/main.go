// در قسمت payment سعی شد نسبت بهanimal کد خوانا تر و نزدیکتر به استاندارد هاس رایج اکوسیستم گولنگ

package main

import (
	"fmt"
	"practice/designPatterns/factoryPatterns/payment/internal/payment"
)

func main() {
	newAcc := payment.FactoryCreditCard{}
	amirAcc := newAcc.CreatePayment(120)
	err := amirAcc.Pay(120)
	if err != nil {
		fmt.Println(err)
	} else {
		amirAcc.ShowBalance()
	}
}
