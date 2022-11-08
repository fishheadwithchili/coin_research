package obtaining_ingredients

import "time"

type ResultData interface {
	// FetchData 只能从我的数据库里拿resultData
	//获取数量取得是公倍数&&不超过maxDataCount的最大值为map的期望长度；
	//每次执行获取后map的指针会放到一个属性里，再次执行该方法后该属性持有的数据量大于map的期望长度，则该属性将指向另一个map，为了gc掉原来那个map，反之直接给map扩充
	FetchData(missions []MissionI, maxDataCount int) (resultData map[string]map[time.Time]int)
	// GetData 从那个缓存属性里获得data
	GetData(mission MissionI) (resultData []data)
}
