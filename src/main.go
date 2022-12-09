package main

import "fmt"

type Car struct {
	brand string
	year  int
	print func(brand string, yeaer int) string
}

func main() {
	myCar := Car{
		brand: "toiota",
		year:  2013,
		print: func(brand string, year int) string {
			return fmt.Sprintf("waos %v %v", brand, year)
		},
	}
	fmt.Println(myCar.print(myCar.brand, myCar.year))
}

func fact(n int) int {
	if n == 1 {
		return n
	}
	return n * fact(n-1)
}
