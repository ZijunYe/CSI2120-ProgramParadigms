//2020 winter
package main

import "fmt"

type Dessert struct {
	Name  string
	Price float32
}

type MainCourse struct {
	Name  string
	Price float32
}

type Meal struct {
	MainCourse
	Dessert
	Total float32
}

func (m *Meal) printMeal() {
	fmt.Printf("Main: %s at %.2f\n", m.MainCourse.Name, m.MainCourse.Price)
	fmt.Printf("Main: %s at %.2f\n", m.Dessert.Name, m.Dessert.Price)
	fmt.Printf("Total:%.2f\n", m.Total)
}

func main() {
	m := Meal{MainCourse{"Schnitzel", 15.50}, Dessert{"Pumpkin Pie", 5.60}, 0.0}
	//Calculate the total price
	m.Total = m.MainCourse.Price + m.Dessert.Price
	m.printMeal()

}
