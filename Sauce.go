package main

type sauceTopping struct {
	pizza IPizza
}

func (add *sauceTopping) getPrice() int {
	pizzaPrice := add.pizza.getPrice()
	return pizzaPrice + 20
}
