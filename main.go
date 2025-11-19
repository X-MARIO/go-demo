package main

import (
	"fmt"
	"math"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	var userKg, userHeight float64 = getUserInput()
	IMT := calculateIMT(userHeight, userKg)
	outputResult(IMT)
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Printf(result)
}

func calculateIMT(userKg float64, userHeight float64) float64 {
	const IMTPower = 2
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT
}

func getUserInput() (float64, float64) {
	var userHeight, userKg float64
	fmt.Println("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Println("Введите свой вес: ")
	fmt.Scan(&userKg)
	return userHeight, userKg
}
