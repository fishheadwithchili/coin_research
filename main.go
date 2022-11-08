package main

import (
	"fmt"
	"github.com/fishheadwithchili/coin_research/steps/obtaining_ingredients"
	"time"
)

func main() {
	var missionOption obtaining_ingredients.MissionOptions
	mission := obtaining_ingredients.NewMission(missionOption.Default)
	mission.Set("ETHUSDT", `{'0':'sql','1':'online'}`, 10, 10, time.Now(), time.Now())
	fmt.Println(mission.ToString())
	mission.Update()
}
