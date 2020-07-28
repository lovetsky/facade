package main

import (
	"facade/internal/account"
	"facade/internal/messages"
	"facade/internal/order"
	"facade/internal/stock"
)

func main() {
	var productUid = map[string]int{"coffee": 1, "pizza": 2}

	a := account.NewAccount("user1")
	m := messages.NewMessage()
	s := stock.NewStock(productUid)

	orderFacade := order.NewOrderFacade(a, m, s)

	// для примера бронируем и возвращаем товар
	orderFacade.Reserved("coffee")
	orderFacade.Reserved("pizza")
	orderFacade.Returned("coffee")
	orderFacade.Reserved("coffee")
}
