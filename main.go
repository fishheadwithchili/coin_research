package main

import (
	"fmt"
	"github.com/fishheadwithchili/coin_research/steps/01obtaining_ingredients"
	"github.com/fishheadwithchili/coin_research/steps/03processing_ingredients"
	"time"
)

func main() {
	fmt.Println("start")
	data := obtaining_ingredients.BaseData.FetchData()
	algo := processing_ingredients.Algo(data)
	var A obtaining_ingredients.BaseData
	var a AS
	A = a
	data1 := A.GetData()

}

type AS struct {
}

func (A AS) FetchData(missions []obtaining_ingredients.Mission, maxDataCount int) (baseData map[time.Time]int) {
	//TODO implement me
	panic("implement me")
}

func (A AS) GetData(mission obtaining_ingredients.Mission) (baseData []obtaining_ingredients.Data) {
	//TODO implement me
	panic("implement me")
}
