package client

type RepBlinds struct {
	Ready      bool //true если прочитаны
	Conflicts  Conflicts
	DefineNaps DefineNaps
	DefPhases  DefPhases
	RPU        OnePlan
	Plans      Plans
	Year       Year
	Weeks      Weeks
	Days       Days
	Status     [3]int
	SerialNum  [2]uint16
}

type Conflicts struct {
	Conflicts map[int]LineConflict
}
type LineConflict struct {
	Number int
	Line   []bool
}
type DefineNaps struct {
	DefineNaps map[int]DefNap
}
type DefNap struct {
	Number      int
	Type        int // 0 -не назначен 1 транспортный 2 пешеход 3 стрелка
	FlashGreen  int
	YellowFlash int
	Red         int
	RedYellow   int
	Keys        []bool
	Tiob        int // 0 выключен 1 постоянный 3 вызывной
}
type DefPhase struct {
	Number int
	Tmin   int
	Naps   []bool
}
type DefPhases struct {
	DefPhases map[int]DefPhase
}
type Plans struct {
	Plans map[int]OnePlan
}
type OnePlan struct {
	Number int
	Type   int //0 локальный план 1 координация
	Tcycle int
	Shift  int
	Lines  []Line
}
type Line struct {
	// SIMPLE = 0, 		//0 простая фаза
	// MGR,			//1 МГР фаза
	// TVP1,			//2 вызывная фаза 1
	// TVP2,			//3 вызывная фаза 2
	// TVP12,			//4 вызывная фаза 1 и 2
	// SUB_TVP1,		//5 замещающая вызывная фаза 1
	// SUB_TVP2,		//6 замещающая вызывная фаза 2
	// SUB_TVP12,		//7 замещающая вызывная фаза 1 и 2
	Type  int
	Phase int //Номер фазы
	Start int
	Stop  int
}
type Year struct {
	Year map[int]Month
}
type Month struct {
	Number int
	Days   []int
}
type Weeks struct {
	Weeks map[int]Week
}
type Week struct {
	Number int
	Days   []int
}
type Days struct {
	Days map[int]DayPlan
}

type DayPlan struct {
	Number int
	Nplans []Nplan
}
type Nplan struct {
	Start int
	Stop  int
	Plan  int
}
