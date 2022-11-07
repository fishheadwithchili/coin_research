package main

import (
	"fmt"
	"github.com/fishheadwithchili/coin_research/steps/01obtaining_ingredients"
	"github.com/fishheadwithchili/coin_research/steps/03processing_ingredients"
)

func main() {
	fmt.Println("start")
	data := obtaining_ingredients.BaseData()
	algo := processing_ingredients.Algo(data)
}
