package main

import (
	"fmt"
	"math"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	var userHeight = 1.8
	var userKg float64 = 100
	var IMT = userKg / math.Pow(userHeight, 2)
	fmt.Print(IMT)
}
