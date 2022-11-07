package obtaining_ingredients

import "time"

type MissionI interface {
	// UpdateMissions 针对断线重启需要检查下之前有没有已完成的记录，更新任务需求
	UpdateMissions(missions []Mission) (updatedMissions []Mission)
}

type Mission struct {
	Subject   string // 数据的主题
	Source    string // 数据来源,一般是sql，调接口拿
	Range     int    // 使用多少个原始数据进行计算
	Offset    int    // 偏移量，比如5，表示每个相邻数据相差5分钟
	StartTime time.Time
	EndTime   time.Time
}

type Data struct {
	Time  time.Time
	Value int
}
type BaseData interface {
	// FetchData get data from sql or net；
	//获取数量取得是公倍数&&不超过maxDataCount的最大值为map的期望长度；
	//每次执行获取后map的指针会放到一个属性里，再次执行该方法后该属性持有的数据量大于map的期望长度，则该属性将指向另一个map，为了gc掉原来那个map，反之直接给map扩充
	FetchData(missions []Mission, maxDataCount int) (baseData map[time.Time]int)
	// GetData 从那个缓存属性里获得data
	GetData(mission Mission) (baseData []Data)
}

type ResultData interface {
	// FetchData 只能从我的数据库里拿resultData
	//获取数量取得是公倍数&&不超过maxDataCount的最大值为map的期望长度；
	//每次执行获取后map的指针会放到一个属性里，再次执行该方法后该属性持有的数据量大于map的期望长度，则该属性将指向另一个map，为了gc掉原来那个map，反之直接给map扩充
	FetchData(missions []Mission, maxDataCount int) (resultData map[Mission]map[time.Time]int)
	// GetData 从那个缓存属性里获得data
	GetData(mission Mission) (resultData []Data)
}
