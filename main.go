package main

import "fmt"

func main() {
	pizza := &MeatMania{}

	pizzaWithMeat := &MeatTopping{
		pizza: pizza,
	}

	pizzaWithMeatAndSauce := &sauceTopping{
		pizza: pizzaWithMeat,
	}

	fmt.Println("Pizza with Meat and Sauce topping is", pizzaWithMeatAndSauce.getPrice())
}
