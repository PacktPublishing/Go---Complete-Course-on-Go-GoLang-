package main

import "fmt"

type PaymentMethod interface {
	Pay(amount float64) string
}

type CardInfoProvider interface {
	CardInfo() string
}

type GiftCard struct {
	Code    string
	Balance float64
}

func (g GiftCard) Pay(amount float64) string {
	if amount > g.Balance {
		return "Not enough balance!"
	}
	return fmt.Sprintf("Paid $%.2f using gift card", amount)
}

func (g GiftCard) CardInfo() string {
	return fmt.Sprintf("Gift card code: %s | Balance: $%.2f", g.Code, g.Balance)
}

func main() {
	card := GiftCard{Code: "GC0001", Balance: 125.0}
	anotherCard := GiftCard{Code: "GC0002", Balance: 78.0}

	var pay PaymentMethod = card
	var info CardInfoProvider = card

	fmt.Println(info.CardInfo())
	fmt.Println(pay.Pay(35.50))

	pay = anotherCard
	info = anotherCard

	fmt.Println(info.CardInfo())
	fmt.Println(pay.Pay(18.00))
}
