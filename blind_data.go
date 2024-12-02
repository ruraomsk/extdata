package client

type RepBlinds struct {
	Ready      bool       //true если прочитаны
	Conflicts  Conflicts  //Таблица конфликтов
	DefineNaps DefineNaps //Описание направлений
	DefPhases  DefPhases  //Описание фаз
	RPU        OnePlan    //Резервный план управления (который работает всегда когда есть ошибки в лругих планах)
	Plans      Plans      //Описание планов управления
	Year       Year       //Годовая карта
	Weeks      Weeks      //Недельные карты
	Days       Days       //Дневные карты
	Status     [3]int
	SerialNum  [2]uint16
}

type Conflicts struct {
	Conflicts map[int]LineConflict
}
type LineConflict struct {
	Number int    // Номер направления
	Line   []bool // Маска конфликтов (true если с этим направлением конфликт)
}
type DefineNaps struct {
	DefineNaps map[int]DefNap
}
type DefNap struct {
	Number      int
	Type        int    // 0 -не назначен 1 транспортный 2 пешеход 3 стрелка
	FlashGreen  int    // время мигания зеленым
	YellowFlash int    // время мигания желтым
	Red         int    // время красного
	RedYellow   int    // время красно-желтого
	Keys        []bool // Ключи
	Tiob        int    // 0 выключен 1 постоянный 3 вызывной
}
type DefPhase struct {
	Number int    //Номер фазы
	Tmin   int    // Минимальное время фазы
	Naps   []bool // Направления задействованные в фазе
}
type DefPhases struct {
	DefPhases map[int]DefPhase
}
type Plans struct {
	Plans map[int]OnePlan
}
type OnePlan struct {
	Number int    // номер фазы
	Type   int    //0 локальный план 1 координация
	Tcycle int    //Время цикла
	Shift  int    //Cдвиг начала цикла для типа 1
	Lines  []Line // Лписание стадий плана
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
	Start int //время начала фазы от старта цикла
	Stop  int //время завершения фазы
}
type Year struct {
	Year map[int]Month
}
type Month struct {
	Number int
	Days   []int // Номера недельной карта на каждый день
}
type Weeks struct {
	Weeks map[int]Week
}
type Week struct {
	Number int
	Days   []int //Номера суточных карт на каждый день недели
}
type Days struct {
	Days map[int]DayPlan
}

type DayPlan struct {
	Number int
	Nplans []Nplan
}
type Nplan struct {
	Start int //Время начала дествия плана (минуты от начала суток)
	Stop  int //Время окончания действия  плана
	Plan  int //Номер плана на этот период
}
