package main

import (
	"./mathmet"
	"fmt"
	"math"
)

//функция для метода касательных,хорд
func f1(x float64) float64 {
	return (x*x - 17)
}

//функция для метода Лагранжа
func f2(x float64) float64 {
	return x*x*x + 3*x*x + 3*x + 1
}

func f3(x float64) float64 {
	return x
}

//функция для метода Эйлера
func f4(x, y float64) float64 {
	return 6*x*x + 5*x*y
}

//функция для метода Эйлера-Коши
func f5(x, y float64) float64 {
	return 3*math.Sin(2*y) + x
}

func main() {
	fmt.Println("Start")
	//a := mathmet.CreateMatr(3, 3)
	//fmt.Println(a)

	//fmt.Println(mathmet.MethNewton(f1, 5, 0.001))
	//fmt.Println(mathmet.MethChord(f1,-5,0,0.001))
	//fmt.Println(mathmet.MethIteration(f1, 4, 0.1))
	//fmt.Println(mathmet.MethInterpolateLagrangePolynomial(f2, 13, 10))
	/*
		x_values := []float64{7, 31, 61, 99, 129, 178, 209}
		y_values := []float64{13, 10, 9, 10, 12, 20, 26}
		mathmet.MethLeastSquare(x_values, y_values)
	*/
	//fmt.Println(mathmet.MethTrap(f3, 0, 5, 0.001))
	//fmt.Println(mathmet.MethTrap(f3, -5, 0, 0.001))
	//fmt.Println(mathmet.MethEuler(f4, 10, 0.01))
	//mathmet.MethEulerKoshi(f5, 0, 1, 0.1, 0, 2)
	//fmt.Println(mathmet.MethMinDih(f3,0,2,1))
	fmt.Println(mathmet.MethMinGold(f3,0,2,0.001))
}
