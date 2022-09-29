package main

type MeatTopping struct {
	pizza IPizza
}

func (add *MeatTopping) getPrice() int {
	pizzaPrice := add.pizza.getPrice()
	return pizzaPrice + 50
}
