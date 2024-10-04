package main

import (
	dp"day08/ex01/describeplant"
	"day08/ex01/model"
	"fmt"
)

func main(){
	un := model.UnknownPlant{
		FlowerType: "rose",
		LeafType: "mazda",
		Color: 228,
	}

	anUn := model.AnotherUnknownPlant{
		FlowerColor: 322,
		LeafType: "honda",
		Height: 52,
	}

	dp.DescribePlant(un)
	fmt.Println()
	dp.DescribePlant(anUn)
}