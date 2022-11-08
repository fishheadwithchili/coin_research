package obtaining_ingredients

import (
	"fmt"
	"time"
)

func NewMission(missionOption missionEnumDefault) (mission MissionI) {
	switch fmt.Sprintf("%v", missionOption) {
	case "default":
		var missionI MissionI
		var missionO missionS
		missionI = &missionO
		return missionI
	default:
		panic("invalid missionOption")
	}
}

type MissionOptions struct {
	Default missionEnumDefault
}

type missionEnumDefault uint8

func (_ missionEnumDefault) String() string {
	return "default"
}

type MissionI interface {
	// Update 针对断线重启需要检查下之前有没有已完成的记录，更新任务需求
	Update()
	Set(subject string, source string, influence int, offset int, startTime time.Time, endTime time.Time)
	ToString() string
}
type missionS struct {
	Subject   string // 数据的主题
	Source    string // 数据来源,一般是sql,调接口拿
	Influence int    // 使用多少个原始数据进行计算
	Offset    int    // 偏移量，比如5，表示每个相邻数据相差5分钟
	StartTime time.Time
	EndTime   time.Time
}

func (m *missionS) ToString() string {
	return fmt.Sprintf("%v", *m) // todo 覆写下它系统的 time.time的 string 便于观看
}

func (m *missionS) Set(subject string, source string, influence int, offset int, startTime time.Time, endTime time.Time) {
	m.Subject = subject
	m.Source = source
	m.Influence = influence
	m.Offset = offset
	m.StartTime = startTime
	m.EndTime = endTime
}

func (m *missionS) Update() {
	fmt.Println("todo Update")
}
