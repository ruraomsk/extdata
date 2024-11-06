package client

import (
	"time"
)

/*
Протокол обмена
При запуске я ожидаю соединения на порту допустим 8888 (как сервер).
Обмен в режиме вопрос ответ. Обмен json в текстовом виде признак конца сообщения \n
При любом разрыве клиент переподключается

Запросы от внешней программы        	Ответы от меня
GetStateHardware - просто Message		структура StateHardware
SetCommand - заполненная структура		сообщение Ok или сообщение с ошибкой
GetSetup - просто Message				структура SetupSubsystem
SetSetup - структура SetupSubsystem 	сообщение Ok или сообщение с ошибкой затем будет перезагружена моя программа
	(потребуется переподключение)
GetStatistics 							структура RepStatistics
*/

const Endline byte = '\n'

type Message struct {
	Message string `json:"message"`
}
type GetStatistics struct {
	Type string `json:"type"` //all - все что хранится в бд устройства last - последняя запись
	//period - указать время начала и время конца
	TimeStart time.Time `json:"start"`
	TimeEnd   time.Time `json:"end"`
}
type RepStatistics struct {
	Counts  []Counts `json:"counts"`  //Проехало ТС
	Ocupaes []Counts `json:"ocupaes"` // Окупация
}

type Counts struct {
	Time   time.Time `json:"time"`
	Values []int     `json:"values"`
}

type Setup struct {
	Name        string      `toml:"name"`
	LogPath     string      `toml:"logpath"`
	Id          int         `toml:"id"`
	DbPath      string      `toml:"dbpath"`
	Modbus      Modbus      `toml:"modbus" json:"modbus"`
	Utopia      Utopia      `toml:"utopia" json:"utopia"`
	SNMP        SNMP        `toml:"snmp" json:"snmp"`
	TrafficData TrafficData `toml:"trafficdata" json:"trafficdata"`
	ModbusRadar ModbusRadar `toml:"modbusradar" json:"modbusradar"`
	Elistar     Elistar     `toml:"elistar" json:"elistar"`
	Micro       Micro       `toml:"micro" json:"micro"`
	Energy      Energy      `toml:"energy" json:"energy"`
	Mgr         Mgr         `toml:"mgr" json:"mgr"`
	Tunel       Tunel       `toml:"tunel" json:"tunel"`
	Comsignal   Comsignal   `toml:"comsignal" json:"comsignal"`
}
type Mgr struct {
	Run      bool  `toml:"run" json:"run"`           //true Run MGR
	Work     bool  `toml:"work" json:"work"`         //true work MGR
	Step     int   `toml:"step" json:"step"`         //Время в секундах которое будет посылаться для инициализации MGR
	Interval int   `toml:"interval" json:"interval"` //Время в секундах за которое будет собираться статистика
	Chanels  []int `toml:"chanels" json:"chanels"`   //Номера каналов из которых будет собираться статистика
	Limits   []int `toml:"limits" json:"limits"`     // Лимиты  на каждый канал
}

// Настройки тунеля TCP-IP to Modbus
type Tunel struct {
	Run  bool `toml:"run" json:"run"`   //true Work Tunel
	Port int  `toml:"port" json:"port"` //Порт для внешнего доступа
	Log  bool `toml:"log" json:"log"`   //Логировать ли входящие запросы?
}
type Energy struct {
	Work     bool   `toml:"work" json:"work"`
	Device   string `toml:"device" json:"device"`
	BaudRate int    `toml:"baudrate" json:"baudrate"`
	Parity   string `toml:"parity" json:"parity"`
	UId      int    `toml:"uid" json:"uid"`
}

type Micro struct {
	Run  bool   `toml:"run" json:"run"`
	Host string `toml:"host" json:"host"`
	Port int    `toml:"port" json:"port"`
	Log  bool   `toml:"log" json:"log"`
	ID   int    `toml:"id" json:"id"`
}

type SNMP struct {
	Run    bool   `toml:"run" json:"run"`
	Listen string `toml:"listen" json:"listen"`
}
type Elistar struct {
	Run bool `toml:"run" json:"run"`
}

// Настройка для приема статистики
type ModbusRadar struct {
	Work    bool   `toml:"work" json:"work"` //True запускать прием
	Master  bool   `toml:"master" json:"master"`
	Debug   bool   `toml:"debug" json:"debug"`
	Host    string `toml:"host" json:"host"`       //Host комплекса
	Port    int    `toml:"port" json:"port"`       //Его порт
	ID      int    `toml:"id" json:"id"`           //Uid
	Chanels int    `toml:"chanels" json:"chanels"` //Кол-во датчиков
}
type Comsignal struct {
	Run      bool
	Device   string `toml:"device" json:"device"`
	BaudRate int    `toml:"baudrate" json:"baudrate"`
	Parity   string `toml:"parity" json:"parity"`
	Debug    bool   `toml:"debug" json:"debug"`
}

type Utopia struct {
	Run         bool
	Device      string `toml:"device" json:"device"`     //Устройство
	BaudRate    int    `toml:"baudrate" json:"baudrate"` //Скорость
	Parity      string `toml:"parity" json:"parity"`     //Паритет
	UId         int    `toml:"uid" json:"uid"`           //Iid
	Debug       bool   `toml:"debug" json:"debug"`
	LostControl int    `toml:"lostControl" json:"lostControl"`
	Recode      bool   `toml:"recode" json:"recode"`
	Replay      bool   `toml:"replay" json:"replay"`
}
type TrafficData struct {
	Work    bool   `toml:"work" json:"work"`
	Debug   bool   `toml:"debug" json:"debug"`
	Host    string `toml:"host" json:"host"`
	Port    int    `toml:"port" json:"port"`
	Listen  int    `toml:"listen" json:"listen"`
	Chanels int    `toml:"chanels" json:"chanels"`
	Diaps   int    `toml:"diaps" json:"diaps"`
	Diap    int    `toml:"diap" json:"diap"`
}
type Modbus struct {
	Device   string `toml:"device" json:"device"`
	BaudRate int    `toml:"baudrate" json:"baudrate"`
	Parity   string `toml:"parity" json:"parity"`
	UId      int    `toml:"uid" json:"uid"`
	Debug    bool   `toml:"debug" json:"debug"`
	Log      bool   `toml:"log" json:"log"`
	Tmin     int    `toml:"tmin" json:"min"` //Минимальная длительность фазы
	Old      bool   `toml:"old" json:"old"`
	TypeKey3 bool   `toml:"key3" json:"key3"` //faile - Нормально замкнутый true -нормально разомкнутый
}

type SetupSubsystem struct {
	Setup Setup
}
type StateHard struct {
	Central       bool      //true управляение центром false - локальное управление
	LastOperation time.Time //время последней операции обмена
	Connect       bool      //true если есть связь с КДМ
	Dark          bool      //true если Режим ОС
	AllRed        bool      //true если Режим Кругом Красный
	Flashing      bool      //true если Режим Желтый Мигающий
	SourceTOOB    bool      //true если Источник времени отсчета внешний
	VPU           bool      //true если включен ВПУ
	WatchDog      uint16    //Текущий Тайм аут управления
	Plan          int       //Номер исполняемого плана контроллером КДМ
	Phase         int       //Номер исполняемой фазы контроллером КДМ
	LastPhase     int       //Номер предыдущей фазы контроллера если промтакт
	WeekCard      int       // Текущая недельная карта
	DayCard       int       // Текущая суточная карта
	ElapTimePk    int       //Оставшееся время до конца фазы ПК
	ElapTimeCoord int       //Оставшееся время до конца входа в координацию

	//Код управления (0 локальное управление
	//1 - внешнее управление по группам
	//2 - внешнее управление по планам
	//3 - ручное управление)
	Source int

	// typedef enum {					//Идентификаторы событий в логе аварий и в регистре событий
	// 	ALL_IS_GOOD = 0,			//Все хорошо, нет предупреждений
	// 	LOW_CURRENT_RED_LAMP,			//Ток через открытый ключ меньше минимального - лампа сгорела, применяется при контроле красных
	// 	NOT_ALLOWED_VOLTAGE_GREEN_OUT, 	//Обнаружено напряжение на закрытом ключе, применяется при контроле зеленых
	// 	NO_CLOCK,				//Нет ответа от микросхемы аппаратных часов
	// 	NO_GPS,					//Нет сигнала от GPS приемника
	// 	NO_POWER_BOARD,			//Нет ответа от платы силовых ключей
	// 	NO_IO_BOARD,				//Нет ответа от платы ввода-вывода
	// 	SHORT_CIRQUIT_KVP			//КЗ цепи кнопки КВП
	// 	WRONG_FILE_VER,		//версия файла конфигурации в ПЗУ не соответствует требуемой
	// 	WRONG_FILE_CRC			//контрольная сумма файла конфигурации в ПЗУ показывает ошибку
	// 	DIRECTIONS_CONFLICT		//обнаружен конфликт направлений
	// 	DC_DIRECTIONS_CONFLICT		//при вызове направлений по сети обнаружен конфликт направлений, вызов отклонен
	// 	NOT_ENTERING_COORDINATION		//не вхождение в координацию
	// }EventId;

	// для событий LOW_CURRENT_RED_LAMP и NOT_ALLOWED_VOLTAGE_GREEN_OUT, S1 содержит номер платы, S2 – номер ключа на плате;
	// для событий NO_POWER_BOARD и NO_IO_BOARD, S1 содержит номер платы;
	// для события SHORT_CIRQUIT_KVP, S1 содержит номер кнопки;
	// для событий DIRECTIONS_CONFLICT и DC_DIRECTIONS_CONFLICT, S1 содержит номер конфликтующего направления
	// для других событий описания не используются.
	Status     []int     //Статус КДМ в его кодировке
	StatusDirs [32]uint8 //Статусы состояния по направлениям
	//   OFF = 0, //все сигналы выключены
	//   DEACTIV_YELLOW=1, //направление перешло в неактивное состояние, желтый после зеленого
	//   DEACTIV_RED=2, //направление перешло в неактивное состояние, красный
	//   ACTIV_RED=3, //направление перешло в активное состояние, красный
	//   ACTIV_REDYELLOW=4, //направление перешло в активное состояние, красный c желтым
	//   ACTIV_GREEN=5, //направление перешло в активное состояние, зеленый
	//   UNCHANGE_GREEN=6, //направление не меняло свое состояние, зеленый
	//   UNCHANGE_RED=7, //направление не меняло свое состояние, красный
	//   GREEN_BLINK=8, //зеленый мигающий сигнал
	//   ZM_YELLOW_BLINK=9, //желтый мигающий в режиме ЖМ
	//   OS_OFF=10,	//сигналы выключены в режиме ОС
	//   UNUSED=11 //неиспользуемое направление
	Tmin         int      //Последнее заданное Тмин вызвать направления
	MaskCommand  uint32   //Последняя маска
	RealWatchDog uint16   //Остаток watchdog
	TOOBs        []uint16 //Счетчики по направлениям
	// MyStatus     int      //Статус в переданной команде центра
	TimeData   []uint16 //текущее временя в контроллере
	HoursAdd   int16    //Смещение часового пояса
	TypeDevice []int    // Ответ от кдм на serverid
	DeviceID   []uint16 //Номер устройства
	Key3       bool     //Состояние концевого выключателя (дверь)
}
type StateHardware struct {
	Message       string    `json:"message"` //StateHardware
	StateHardware StateHard //Состояние контроллера
}

// type SetCommand struct {
// }

type CommandForDevice struct {
	Plan     int  `json:"plan"`     // 0 - снять управление
	Phase    int  `json:"phase"`    // 0 - снять управление
	Dark     bool `json:"dark"`     // false  - отменить true-включить
	AllRed   bool `json:"allred"`   // false  - отменить true-включить
	Flashing bool `json:"flashing"` // false  - отменить true-включить
}
