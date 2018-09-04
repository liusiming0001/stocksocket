package msg

import (
	"StockSocket/cache"
	"StockSocket/lib"
	"fmt"
	"github.com/pquerna/ffjson/ffjson"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	ZX    = 0  //最新
	RK    = 1  //日K
	FS    = 2  //分时
	CJMX  = 3  //成交明细
	ZK    = 4  //周K
	YK    = 5  //月K
	FZ1K  = 6  //1分钟K
	FZ5K  = 7  //5分钟K
	FZ15K = 8  //15分钟K
	FZ30K = 9  //30分钟K
	FZ60K = 10 //60分钟K
)

type MsgRedisPool struct {
	ZXRedis cache.RedisPool
	FSRedis cache.RedisPool
	RKRedis cache.RedisPool
	ZKRedis cache.RedisPool
	YKRedis cache.RedisPool
}
type TempCacheModel struct {
	lastMinuteTemp sync.Map
	ZKTemp         sync.Map
}
type ValidTimeModel struct {
	dtNow                   time.Time
	vaildTimeStart          time.Time
	vaildTimeEnd            time.Time
	vaildTimeAfternoonStart time.Time
	vaildTimeAfternoonEnd   time.Time
}
type HqModel struct {
	//时间戳
	UnixTime int64
	//接受日期
	RecTime time.Time
	//接受日期字符串
	RecTimeStr string
	/// 频道ID
	ChannelNo int16
	// 行情类别
	MDStreamID string
	//股票号
	SyNo string
	//股票来源
	SySource string
	//股票状态
	TradingPhaseCode string

	//昨日收盘
	ZRSP float64
	//成交笔数
	CJBS int64
	//成交数量
	CJSL float64
	//成交金额
	CJJE float64

	//买一价格 数量
	HQBJW1 float64
	HQBSL1 float64

	HQBJW2 float64
	HQBSL2 float64

	HQBJW3 float64
	HQBSL3 float64

	HQBJW4 float64
	HQBSL4 float64

	HQBJW5 float64
	HQBSL5 float64

	//卖一价格 数量
	HQSJW1 float64
	HQSSL1 float64

	HQSJW2 float64
	HQSSL2 float64

	HQSJW3 float64
	HQSSL3 float64

	HQSJW4 float64
	HQSSL4 float64

	HQSJW5 float64
	HQSSL5 float64

	//市盈率1
	HQSYL1 float64
	//市盈率2
	HQSYL2 float64

	//最近成交
	ZJCJ float64
	//开盘价
	KPJ float64
	//最高价
	ZGJ float64
	//最低价
	ZDJ float64

	//数据时间
	DataTime string
	//是否停牌·
	IsStop bool

	//现在时间
	DtNow string
	//上一分钟
	Upmin string
}
type FSModel struct {
	//月日+分时时间 MMddHHmm
	kstime string `json:"kstime"`
	//输出保存时间
	Dtnow string `json:"dtnow"`
	//股票代码
	SyNo string
	//股票名称
	Symbol string
	//分时时间  HHmm
	Time string `json:"time"`
	//本分钟开盘价
	HQJRKP string
	//本分钟最高价
	HQZGCJ string
	//本分钟最低价
	HQZDCJ string
	//今日开盘价
	HQJRKPDAY string
	//今日最高价
	HQZGCJDAY string
	//今日最低价
	HQZDCJDAY string
	//最新成交价
	HQZJCJ string
	//昨日收盘价
	HQZRSP string
	//总成交金额
	HQCJJE string
	//总成交数量
	HQCJSL string
	//本分钟成交金额差值
	HQCJJECZ string
	//本分钟成交数量差值
	HQCJSLCZ string
	//卖一价
	HQSJW1 string
	//买一价
	HQBJW1 string
}
type ZXModel struct {
	Symbol  string
	Time    string `json:"time"`
	Timere  string `json:"Timere"`
	SyNo    string
	Time_hi string `json:"time_hi"`
	HQJRKP  string
	HQZGCJ  string
	HQZDCJ  string
	HQJRSP  string
	HQZRSP  string
	HQCJJE  string
	HQCJSL  string

	HQUPSP    string
	HQCJSLCZ  string
	HQCJJECZ  string
	HQSEDSLCZ string
	HQSEDJECZ string

	//买一价格 数量
	HQBJW1 string
	HQBSL1 string

	HQBJW2 string
	HQBSL2 string

	HQBJW3 string
	HQBSL3 string

	HQBJW4 string
	HQBSL4 string

	HQBJW5 string
	HQBSL5 string

	//卖一价格 数量
	HQSJW1 string
	HQSSL1 string

	HQSJW2 string
	HQSSL2 string

	HQSJW3 string
	HQSSL3 string

	HQSJW4 string
	HQSSL4 string

	HQSJW5 string
	HQSSL5 string

	//市盈率1
	HQSYL1 string
	//市盈率2
	HQSYL2 string

	ISSTOP bool
}
type RKModel struct {
	Year   string
	Date   string
	SyNo   string
	HQJRKP string
	HQZGCJ string
	HQZDCJ string
	HQJRSP string
	HQZRSP string
	HQCJJE string
	HQCJSL string
	HQZDE  string
	HQZDL  string
}
type ZKModel struct {
	SyNo   string
	HQJRKP string
	HQZGCJ string
	HQZDCJ string
	HQJRSP string
	HQZRSP string
	HQCJJE string
	HQCJSL string
	HQZDE  string
	HQZDL  string
	Start  string `json:"start"`
	End    string `json:"end"`
}
type YKModel struct {
	SyNo   string
	HQJRKP string
	HQZGCJ string
	HQZDCJ string
	HQJRSP string
	HQZRSP string
	HQCJJE string
	HQCJSL string
	HQZDE  string
	HQZDL  string
	Start  string `json:"start"`
	End    string `json:"end"`
}

var (
	ChinaLocal     *time.Location
	validTimeModel ValidTimeModel
	RedisHost      = "127.0.0.1:6379"
	RedisAuth      = "123qwe"
	MsgRedis       = MsgRedisPool{}
)

func InitRedisPool() {
	MsgRedis.ZXRedis = cache.RedisPool{
		RedisHost: RedisHost,
		RedisAuth: RedisAuth,
		RedisDb:   ZX,
	}
	MsgRedis.FSRedis = cache.RedisPool{
		RedisHost: RedisHost,
		RedisAuth: RedisAuth,
		RedisDb:   FS,
	}
	MsgRedis.RKRedis = cache.RedisPool{
		RedisHost: RedisHost,
		RedisAuth: RedisAuth,
		RedisDb:   RK,
	}
	MsgRedis.ZKRedis = cache.RedisPool{
		RedisHost: RedisHost,
		RedisAuth: RedisAuth,
		RedisDb:   ZK,
	}
	MsgRedis.YKRedis = cache.RedisPool{
		RedisHost: RedisHost,
		RedisAuth: RedisAuth,
		RedisDb:   YK,
	}
	MsgRedis.ZXRedis.GetPool()
	MsgRedis.FSRedis.GetPool()
	MsgRedis.RKRedis.GetPool()
	MsgRedis.ZKRedis.GetPool()
	MsgRedis.YKRedis.GetPool()
}
func InitValidaTime() {
	ChinaLocal, _ = time.LoadLocation("Local")
	validTimeModel.dtNow = time.Now()
	validTimeModel.vaildTimeStart, _ = time.ParseInLocation("2006-01-02 15:04:05", validTimeModel.dtNow.Format("2006-01-02")+" 09:30:00", ChinaLocal)
	validTimeModel.vaildTimeEnd, _ = time.ParseInLocation("2006-01-02 15:04:05", validTimeModel.dtNow.Format("2006-01-02")+" 15:00:59", ChinaLocal)
	validTimeModel.vaildTimeAfternoonStart, _ = time.ParseInLocation("2006-01-02 15:04:05", validTimeModel.dtNow.Format("2006-01-02")+" 11:30:59", ChinaLocal)
	validTimeModel.vaildTimeAfternoonEnd, _ = time.ParseInLocation("2006-01-02 15:04:05", validTimeModel.dtNow.Format("2006-01-02")+" 13:00:00", ChinaLocal)
	fmt.Println(validTimeModel.vaildTimeStart.String())
	fmt.Println(validTimeModel.vaildTimeEnd.String())
	fmt.Println(validTimeModel.vaildTimeAfternoonStart.String())
	fmt.Println(validTimeModel.vaildTimeAfternoonEnd.String())
}

//解析300111号报文
func Save300111(input []byte) {
	bytes := input
	var hqModel HqModel
	hqModel.DtNow = time.Now().Format("2006-01-02 15:04:05")
	//0-8
	hqModel.RecTime = lib.GetTimeFromFormatIntByte(bytes[0:8])
	hqModel.UnixTime = hqModel.RecTime.Unix()
	if hqModel.UnixTime < 0 {
		return
	}
	if hqModel.RecTime.Before(validTimeModel.vaildTimeStart) {
		return
	}
	if hqModel.RecTime.After(validTimeModel.vaildTimeEnd) {
		hqModel.RecTime = validTimeModel.vaildTimeEnd
	}
	if hqModel.RecTime.After(validTimeModel.vaildTimeAfternoonStart) && hqModel.RecTime.Before(validTimeModel.vaildTimeAfternoonEnd) {
		hqModel.RecTime = validTimeModel.vaildTimeAfternoonStart
	}

	hqModel.DataTime = hqModel.RecTime.Format("1504")
	hqModel.RecTimeStr = hqModel.RecTime.Format("200601021504")
	m1, _ := time.ParseDuration("-1m")
	hqModel.Upmin = hqModel.RecTime.Add(m1).Format("01021504")

	//周六日排除，其实也是不会有的
	if hqModel.RecTime.Weekday() == time.Saturday || hqModel.RecTime.Weekday() == time.Sunday {
		return
	}

	//8:10
	//hqModel.ChannelNo = lib.BytesToInt16(bytes[8:10])
	//10:13
	//hqModel.MDStreamID = string(bytes[10:13])
	//13:21
	hqModel.SyNo = strings.TrimSpace(string(bytes[13:21]))
	/*if strings.Contains(hqModel.SyNo, "\u0000") {
		return
	}
	if strings.Contains(hqModel.SyNo, " ") {
		return
	}
	if len(hqModel.SyNo) != 6 {
		return
	}*/
	//21:25
	//hqModel.SySource = string(bytes[21:25])
	//25:33
	hqModel.TradingPhaseCode = strings.TrimSpace(string(bytes[25:33]))
	//33:41
	hqModel.ZRSP = float64(lib.BytesToInt64(bytes[33:41])) / 10000
	//41:49
	hqModel.CJBS = lib.BytesToInt64(bytes[41:49])
	//49:57
	hqModel.CJSL = float64(lib.BytesToInt64(bytes[49:57])) / 100
	//57:65
	hqModel.CJJE = float64(lib.BytesToInt64(bytes[57:65])) / 10000
	//65:69
	hqCount := lib.BytesToInt32(bytes[65:69])

	baseCount := 69

	for i := int32(0); i < hqCount; i++ {
		if (baseCount + 32) > len(bytes) {
			break
		}
		MDEntryType := strings.TrimSpace(string(bytes[baseCount : baseCount+2]))
		MDEntryPx := float64(lib.BytesToInt64(bytes[baseCount+2:baseCount+10])) / 1000000
		MDEntrySize := float64(lib.BytesToInt64(bytes[baseCount+10:baseCount+18])) / 100
		MDPriceLevel := lib.BytesToInt16(bytes[baseCount+18 : baseCount+20])
		//NumberOfOrders := BytesToInt64(bytes[baseCount+20 : baseCount+28])
		//NoOrders := BytesToInt32(bytes[baseCount+28 : baseCount+32])
		baseCount = baseCount + 32
		switch MDEntryType {
		case "0":
			switch MDPriceLevel {
			case 1:
				hqModel.HQBJW1 = MDEntryPx
				hqModel.HQBSL1 = MDEntrySize
				break
			case 2:
				hqModel.HQBJW2 = MDEntryPx
				hqModel.HQBSL2 = MDEntrySize
				break
			case 3:
				hqModel.HQBJW3 = MDEntryPx
				hqModel.HQBSL3 = MDEntrySize
				break
			case 4:
				hqModel.HQBJW4 = MDEntryPx
				hqModel.HQBSL4 = MDEntrySize
				break
			case 5:
				hqModel.HQBJW5 = MDEntryPx
				hqModel.HQBSL5 = MDEntrySize
				break
			default:
				break
			}
			break
		case "1":
			switch MDPriceLevel {
			case 1:
				hqModel.HQSJW1 = MDEntryPx
				hqModel.HQSSL1 = MDEntrySize
				break
			case 2:
				hqModel.HQSJW2 = MDEntryPx
				hqModel.HQSSL2 = MDEntrySize
				break
			case 3:
				hqModel.HQSJW3 = MDEntryPx
				hqModel.HQSSL3 = MDEntrySize
				break
			case 4:
				hqModel.HQSJW4 = MDEntryPx
				hqModel.HQSSL4 = MDEntrySize
				break
			case 5:
				hqModel.HQSJW5 = MDEntryPx
				hqModel.HQSSL5 = MDEntrySize
				break
			default:
				break
			}
			break
		case "2":
			hqModel.ZJCJ = MDEntryPx
			break
		case "4":
			hqModel.KPJ = MDEntryPx
			break
		case "7":
			hqModel.ZGJ = MDEntryPx
			break
		case "8":
			hqModel.ZDJ = MDEntryPx
			break
		case "x1":
			break
		case "x2":
			break
		case "x3":
			break
		case "x4":
			break
		case "x5":
			hqModel.HQSYL1 = MDEntryPx
			break
		case "x6":
			hqModel.HQSYL1 = MDEntryPx
			break
		case "x7":
			break
		case "x8":
			break
		case "x9":
			break
		case "xe":
			break
		case "xf":
			break
		case "xg":
			break
		default:
			break
		}
	}
	if strings.Contains(hqModel.TradingPhaseCode, "0") {
		if hqModel.TradingPhaseCode == "H0" {
			hqModel.IsStop = true
		} else {
			hqModel.IsStop = false
		}
	} else if strings.Contains(hqModel.TradingPhaseCode, "1") {
		hqModel.IsStop = true
	}
	/*js, _ := ffjson.Marshal(hqModel)
	jsStr := string(js)
	cache.StringRedisSet(hqModel.SyNo, jsStr, 15)*/
	var last1model *FSModel
	if hqModel.RecTime.Before(validTimeModel.vaildTimeStart) {
		last1model = nil
	} else if hqModel.RecTime.After(validTimeModel.vaildTimeEnd) {
		hqModel.Upmin = hqModel.RecTime.Format("0102") + "1459"
		last1model = getHashModel(hqModel.SyNo, hqModel.Upmin)
	} else if hqModel.RecTime.Equal(validTimeModel.vaildTimeAfternoonEnd) {
		hqModel.Upmin = hqModel.RecTime.Format("0102") + "1130"
		last1model = getHashModel(hqModel.SyNo, hqModel.Upmin)
	} else {
		last1model = getHashModel(hqModel.SyNo, hqModel.Upmin)
	}
	//if hqModel.TradingPhaseCode == "E0" || hqModel.TradingPhaseCode == "T0" || hqModel.TradingPhaseCode == "C0" || hqModel.TradingPhaseCode == "A0" {
	/*go FenShiDataInsert(hqModel, last1model)
	go FenShiDataInsert(hqModel, last1model)
	go ZuiXinDataInsert(hqModel, last1model)
	go RiKInsert(hqModel)
	go ZhouKInsert(hqModel, last1model)*/

	last3model := ZuiXinDataInsert(hqModel, last1model)
	FenShiDataInsert(hqModel, last3model)
	RiKInsert(hqModel)
	ZhouKInsert(hqModel, last3model)
	YueKInsert(hqModel, last3model)
}

func ZuiXinDataInsert(model HqModel, last1model *FSModel) *ZXModel {
	var TranModel ZXModel
	TranModel.SyNo = model.SyNo
	TranModel.Time = model.DtNow
	TranModel.Time_hi = model.DataTime
	TranModel.Timere = model.RecTimeStr
	TranModel.ISSTOP = model.IsStop

	//昨日收盘价
	TranModel.HQZRSP = floatToSting(model.ZRSP)
	//总成交量
	TranModel.HQCJSL = floatToSting(model.CJSL)
	//总成交金额
	TranModel.HQCJJE = floatToSting(model.CJJE)
	TranModel.SyNo = model.SyNo

	if TranModel.ISSTOP {
		TranModel.HQUPSP = floatToSting(model.ZRSP)
		TranModel.HQJRSP = floatToSting(model.ZRSP)
		TranModel.HQJRKP = floatToSting(model.ZRSP)
		TranModel.HQZGCJ = floatToSting(model.ZRSP)
		TranModel.HQZDCJ = floatToSting(model.ZRSP)
	} else {
		//今日开盘价
		TranModel.HQJRKP = floatToSting(model.KPJ)
		//最近成交价格
		TranModel.HQJRSP = floatToSting(model.ZJCJ)
		//最高成交价
		TranModel.HQZGCJ = floatToSting(model.ZGJ)
		//最低成交价
		TranModel.HQZDCJ = floatToSting(model.ZDJ)
	}

	TranModel.HQBJW1 = floatToDot2Sting(model.HQBJW1)
	TranModel.HQBSL1 = floatToDot2Sting(model.HQBSL1)
	//买二
	TranModel.HQBJW2 = floatToDot2Sting(model.HQBJW2)
	TranModel.HQBSL2 = floatToDot2Sting(model.HQBSL2)
	//买三
	TranModel.HQBJW3 = floatToDot2Sting(model.HQBJW3)
	TranModel.HQBSL3 = floatToDot2Sting(model.HQBSL3)
	//买四
	TranModel.HQBJW4 = floatToDot2Sting(model.HQBJW4)
	TranModel.HQBSL4 = floatToDot2Sting(model.HQBSL4)
	//买五
	TranModel.HQBJW5 = floatToDot2Sting(model.HQBJW5)
	TranModel.HQBSL5 = floatToDot2Sting(model.HQBSL5)
	//卖出
	TranModel.HQSJW1 = floatToDot2Sting(model.HQSJW1)
	TranModel.HQSSL1 = floatToDot2Sting(model.HQSSL1)
	//卖二
	TranModel.HQSJW2 = floatToDot2Sting(model.HQSJW2)
	TranModel.HQSSL2 = floatToDot2Sting(model.HQSSL2)
	//卖三
	TranModel.HQSJW3 = floatToDot2Sting(model.HQSJW3)
	TranModel.HQSSL3 = floatToDot2Sting(model.HQSSL3)
	//卖四
	TranModel.HQSJW4 = floatToDot2Sting(model.HQSJW4)
	TranModel.HQSSL4 = floatToDot2Sting(model.HQSSL4)
	//卖五
	TranModel.HQSJW5 = floatToDot2Sting(model.HQSJW5)
	TranModel.HQSSL5 = floatToDot2Sting(model.HQSSL5)

	//市盈率1
	TranModel.HQSYL1 = floatToDot2Sting(model.HQSYL1)
	//市盈率2
	TranModel.HQSYL2 = floatToDot2Sting(model.HQSYL2)

	last3model := jsonToZXModel(MsgRedis.ZXRedis.StringRedisGet(TranModel.SyNo))

	if last3model == nil {
		//和上3秒的成交数量（股）
		TranModel.HQSEDSLCZ = TranModel.HQCJSL
		//和上3秒的成交金额
		TranModel.HQSEDJECZ = TranModel.HQCJJE
	} else {
		date, _ := time.ParseInLocation("2006-01-02", last3model.Time, ChinaLocal)

		if date.Day() == model.RecTime.Day() {
			//和上3秒的成交数量（股）
			TranModel.HQSEDSLCZ = floatToDot2Sting(stringToFloat(TranModel.HQCJSL) - stringToFloat(last3model.HQCJSL))
			//和上3秒的成交金额
			TranModel.HQSEDJECZ = floatToDot2Sting(stringToFloat(TranModel.HQCJJE) - stringToFloat(last3model.HQCJJE))
		} else {
			//和上3秒的成交数量（股）
			TranModel.HQSEDSLCZ = TranModel.HQCJSL
			//和上3秒的成交金额
			TranModel.HQSEDJECZ = TranModel.HQCJJE
		}
	}

	if last1model != nil {
		//上分钟收盘价
		TranModel.HQUPSP = last1model.HQZJCJ
		//本分钟成交金额差值
		TranModel.HQCJJECZ = floatToDot2Sting(stringToFloat(TranModel.HQCJJE) - stringToFloat(last1model.HQCJJE))
		//本分钟成交数量差值
		TranModel.HQCJSLCZ = floatToDot2Sting(stringToFloat(TranModel.HQCJSL) - stringToFloat(last1model.HQCJSL))
	} else {
		//上分钟收盘价
		TranModel.HQUPSP = TranModel.HQJRSP
		TranModel.HQCJJECZ = TranModel.HQCJJE
		TranModel.HQCJSLCZ = TranModel.HQCJSL
	}
	js, _ := ffjson.Marshal(TranModel)
	MsgRedis.ZXRedis.StringRedisSet(TranModel.SyNo, string(js))
	return &TranModel
}

//分时数据
func FenShiDataInsert(hqModel HqModel, last3model *ZXModel) {
	var fsModel FSModel
	fsModel.kstime = hqModel.RecTime.Format("01021504")
	fsModel.Time = hqModel.DataTime
	fsModel.Dtnow = hqModel.RecTimeStr
	fsModel.HQZRSP = floatToSting(hqModel.ZRSP)
	fsModel.HQCJSL = floatToSting(hqModel.CJSL)
	fsModel.HQCJJE = floatToSting(hqModel.CJJE)
	fsModel.SyNo = hqModel.SyNo
	if hqModel.IsStop {
		fsModel.HQJRKP = fsModel.HQZRSP
		fsModel.HQJRKPDAY = fsModel.HQZRSP
		fsModel.HQZJCJ = fsModel.HQZRSP
		fsModel.HQZGCJDAY = fsModel.HQZRSP
		fsModel.HQZDCJDAY = fsModel.HQZRSP
		fsModel.HQZGCJ = fsModel.HQZRSP
		fsModel.HQZDCJ = fsModel.HQZRSP
	} else {
		fsModel.HQJRKPDAY = floatToSting(hqModel.KPJ)
		fsModel.HQZJCJ = floatToSting(hqModel.ZJCJ)
		fsModel.HQZGCJDAY = floatToSting(hqModel.ZGJ)
		fsModel.HQZDCJDAY = floatToSting(hqModel.ZDJ)
	}
	fsModel.HQBJW1 = floatToSting(hqModel.HQBJW1)
	fsModel.HQSJW1 = floatToSting(hqModel.HQSJW1)

	if last3model == nil {
		fsModel.HQZGCJ = fsModel.HQZJCJ
		fsModel.HQZDCJ = fsModel.HQZJCJ
	} else {
		if stringToFloat(fsModel.HQZJCJ) > stringToFloat(last3model.HQZGCJ) {
			fsModel.HQZGCJ = fsModel.HQZJCJ
		} else {
			fsModel.HQZGCJ = last3model.HQZGCJ
		}

		if stringToFloat(fsModel.HQZJCJ) < stringToFloat(last3model.HQZDCJ) {
			fsModel.HQZDCJ = fsModel.HQZJCJ
		} else {
			fsModel.HQZDCJ = last3model.HQZDCJ
		}
	}
	/*if last1model == nil {
		fsModel.HQJRKP = fsModel.HQZJCJ
		fsModel.HQCJJECZ = fsModel.HQCJJE
		fsModel.HQCJSLCZ = fsModel.HQCJSL
	} else {
		if last3model == nil {
			fsModel.HQJRKP = last1model.HQZJCJ
		} else {
			fsModel.HQJRKP = last3model.HQJRKP
		}
		//本分钟成交金额差值
		fsModel.HQCJJECZ = floatToDot2Sting(stringToFloat(fsModel.HQCJJE) - stringToFloat(last1model.HQCJJE))
		//本分钟成交数量差值
		fsModel.HQCJSLCZ = floatToDot2Sting(stringToFloat(fsModel.HQCJSL) - stringToFloat(last1model.HQCJSL))
	}*/
	//上分钟收盘价
	fsModel.HQJRKP = last3model.HQUPSP
	//本分钟成交金额差值
	fsModel.HQCJJECZ = last3model.HQCJJECZ
	//本分钟成交数量差值
	fsModel.HQCJSLCZ = last3model.HQCJSLCZ

	js, _ := ffjson.Marshal(fsModel)
	MsgRedis.FSRedis.MsgRedisHashSet(fsModel.SyNo, fsModel.kstime, string(js))

}
func RiKInsert(model HqModel) {
	var TranModel RKModel
	TranModel.Year = strconv.Itoa(model.RecTime.Year())
	TranModel.Date = model.RecTime.Format("2006-01-02")
	TranModel.SyNo = model.SyNo
	//昨日收盘价
	TranModel.HQZRSP = floatToSting(model.ZRSP)
	//总成交量
	TranModel.HQCJSL = floatToSting(model.CJSL)
	//总成交金额
	TranModel.HQCJJE = floatToSting(model.CJJE)
	if model.IsStop {
		TranModel.HQJRSP = floatToSting(model.ZRSP)
		TranModel.HQJRKP = floatToSting(model.ZRSP)
		TranModel.HQZGCJ = floatToSting(model.ZRSP)
		TranModel.HQZDCJ = floatToSting(model.ZRSP)
	} else {
		//今日开盘价
		TranModel.HQJRKP = floatToSting(model.KPJ)
		//最近成交价格
		TranModel.HQJRSP = floatToSting(model.ZJCJ)
		//最高成交价
		TranModel.HQZGCJ = floatToSting(model.ZGJ)
		//最低成交价
		TranModel.HQZDCJ = floatToSting(model.ZDJ)
	}
	zde := model.ZJCJ - model.ZRSP
	TranModel.HQZDE = floatToDot2Sting(zde)
	TranModel.HQZDL = floatToDot2Sting((zde / model.ZRSP) * 100)

	lastJson := MsgRedis.RKRedis.GetListIndexValue(TranModel.SyNo, 0)
	lastmodel := jsonToRKModel(lastJson)
	js, _ := ffjson.Marshal(TranModel)
	if lastmodel == nil {
		MsgRedis.RKRedis.ListLefPush(TranModel.SyNo, string(js))
	} else {
		if lastmodel.Date != TranModel.Date {
			MsgRedis.RKRedis.ListLefPush(TranModel.SyNo, *lastJson)
		}
		MsgRedis.RKRedis.SetListIndexValue(TranModel.SyNo, 0, string(js))
	}
}
func ZhouKInsert(model HqModel, last3model *ZXModel) {
	weekOfDay := int(model.RecTime.Weekday()) - 1
	monday := model.RecTime.AddDate(0, 0, -weekOfDay)
	friday := monday.AddDate(0, 0, 4)
	thisWeekModel := jsonToZKModel(MsgRedis.ZKRedis.GetListIndexValue(model.SyNo, 0))
	if thisWeekModel == nil {
		zde := model.ZJCJ - model.ZRSP
		zdl := floatToDot2Sting(zde / model.ZRSP * 100)
		//var TranModel ZKModel
		TranModel := ZKModel{
			model.SyNo,
			floatToSting(model.KPJ),
			floatToSting(model.ZGJ),
			floatToSting(model.ZDJ),
			floatToSting(model.ZJCJ),
			floatToSting(model.ZRSP),
			floatToSting(model.CJJE),
			floatToSting(model.CJSL),
			floatToDot2Sting(zde),
			zdl,
			monday.Format("20060102") + "0930",
			friday.Format("20060102") + "1500",
		}
		js, _ := ffjson.Marshal(TranModel)
		MsgRedis.ZKRedis.ListLefPush(TranModel.SyNo, string(js))
	} else {
		//chinaLocal, _ := time.LoadLocation("Local")
		startDate, _ := time.ParseInLocation("200601021504", thisWeekModel.Start, ChinaLocal)
		endDate, _ := time.ParseInLocation("200601021504", thisWeekModel.End, ChinaLocal)
		//如果在此区间，更新数据
		if model.RecTime.After(startDate) && model.RecTime.Before(endDate) {
			sp := stringToFloat(thisWeekModel.HQZRSP)
			zde := model.ZJCJ - sp
			zdl := floatToDot2Sting(zde / sp * 100)
			if model.ZGJ > stringToFloat(thisWeekModel.HQZGCJ) {
				thisWeekModel.HQZGCJ = floatToSting(model.ZGJ)
			}
			if model.ZDJ < stringToFloat(thisWeekModel.HQZDCJ) {
				thisWeekModel.HQZDCJ = floatToSting(model.ZDJ)
			}
			thisWeekModel.HQZDE = floatToDot2Sting(zde)
			thisWeekModel.HQZDL = zdl

			//var cjslCZ = model.CJSL - last3smodel.HQCJSL.StringToDeciaml();
			thisWeekModel.HQCJSL = floatToDot2Sting(stringToFloat(thisWeekModel.HQCJSL) + stringToFloat(last3model.HQSEDSLCZ))

			//var cjjeCZ = model.CJJE - last3smodel.HQCJJE.StringToDeciaml();
			thisWeekModel.HQCJJE = floatToDot2Sting(stringToFloat(thisWeekModel.HQCJJE) + stringToFloat(last3model.HQSEDJECZ))

			js, _ := ffjson.Marshal(thisWeekModel)
			MsgRedis.ZKRedis.SetListIndexValue(model.SyNo, 0, string(js))
		} else if startDate.Before(monday) {
			//此周第一条 读出第一条数据是上周的
			//把本周第一条数据弄进去
			zde := model.ZJCJ - model.ZRSP
			zdl := floatToDot2Sting(zde / model.ZRSP * 100)
			TranModel := ZKModel{
				model.SyNo,
				floatToSting(model.KPJ),
				floatToSting(model.ZGJ),
				floatToSting(model.ZDJ),
				floatToSting(model.ZJCJ),
				floatToSting(model.ZRSP),
				floatToSting(model.CJJE),
				floatToSting(model.CJSL),
				floatToDot2Sting(zde),
				zdl,
				monday.Format("20060102") + "0930",
				friday.Format("20060102") + "1500",
			}
			thisJs, _ := ffjson.Marshal(TranModel)
			MsgRedis.ZKRedis.ListLefPush(TranModel.SyNo, string(thisJs))
		}
	}
}
func YueKInsert(model HqModel, last3model *ZXModel) {
	monthStart := time.Date(model.RecTime.Year(), model.RecTime.Month(), 1, 0, 0, 0, 0, time.Local)
	monthEnd := monthStart.AddDate(0, 1, -1)

	for {
		if monthStart.Weekday() == time.Saturday || monthStart.Weekday() == time.Sunday {
			monthStart = monthStart.AddDate(0, 0, 1)
		} else {
			break
		}
	}
	for {
		if monthEnd.Weekday() == time.Saturday || monthEnd.Weekday() == time.Sunday {
			monthEnd = monthEnd.AddDate(0, 0, -1)
		} else {
			break
		}
	}
	start := monthStart.Format("20060102") + "0930"
	end := monthEnd.Format("20060102") + "1500"
	thisMonthModel := jsonToYKModel(MsgRedis.YKRedis.GetListIndexValue(model.SyNo, 0))
	if thisMonthModel == nil {
		zde := model.ZJCJ - model.ZRSP
		zdl := floatToDot2Sting(zde / model.ZRSP * 100)
		//var TranModel ZKModel
		TranModel := ZKModel{
			model.SyNo,
			floatToSting(model.KPJ),
			floatToSting(model.ZGJ),
			floatToSting(model.ZDJ),
			floatToSting(model.ZJCJ),
			floatToSting(model.ZRSP),
			floatToSting(model.CJJE),
			floatToSting(model.CJSL),
			floatToDot2Sting(zde),
			zdl,
			start,
			end,
		}
		js, _ := ffjson.Marshal(TranModel)
		MsgRedis.YKRedis.ListLefPush(TranModel.SyNo, string(js))
	} else {
		startDate, _ := time.ParseInLocation("200601021504", thisMonthModel.Start, ChinaLocal)
		endDate, _ := time.ParseInLocation("200601021504", thisMonthModel.End, ChinaLocal)
		//如果在此区间，更新数据
		if model.RecTime.After(startDate) && model.RecTime.Before(endDate) {
			sp := stringToFloat(thisMonthModel.HQZRSP)
			zde := model.ZJCJ - sp
			zdl := floatToDot2Sting(zde / sp * 100)
			if model.ZGJ > stringToFloat(thisMonthModel.HQZGCJ) {
				thisMonthModel.HQZGCJ = floatToSting(model.ZGJ)
			}
			if model.ZDJ < stringToFloat(thisMonthModel.HQZDCJ) {
				thisMonthModel.HQZDCJ = floatToSting(model.ZDJ)
			}
			thisMonthModel.HQZDE = floatToDot2Sting(zde)
			thisMonthModel.HQZDL = zdl

			//var cjslCZ = model.CJSL - last3smodel.HQCJSL.StringToDeciaml();
			thisMonthModel.HQCJSL = floatToDot2Sting(stringToFloat(thisMonthModel.HQCJSL) + stringToFloat(last3model.HQSEDSLCZ))

			//var cjjeCZ = model.CJJE - last3smodel.HQCJJE.StringToDeciaml();
			thisMonthModel.HQCJJE = floatToDot2Sting(stringToFloat(thisMonthModel.HQCJJE) + stringToFloat(last3model.HQSEDJECZ))

			js, _ := ffjson.Marshal(thisMonthModel)
			MsgRedis.YKRedis.SetListIndexValue(model.SyNo, 0, string(js))
		} else if startDate.Before(monthStart) {
			//此周第一条 读出第一条数据是上周的
			//把本周第一条数据弄进去
			zde := model.ZJCJ - model.ZRSP
			zdl := floatToDot2Sting(zde / model.ZRSP * 100)
			TranModel := ZKModel{
				model.SyNo,
				floatToSting(model.KPJ),
				floatToSting(model.ZGJ),
				floatToSting(model.ZDJ),
				floatToSting(model.ZJCJ),
				floatToSting(model.ZRSP),
				floatToSting(model.CJJE),
				floatToSting(model.CJSL),
				floatToDot2Sting(zde),
				zdl,
				start,
				end,
			}
			thisJs, _ := ffjson.Marshal(TranModel)
			MsgRedis.YKRedis.ListLefPush(TranModel.SyNo, string(thisJs))
		}
	}
}

func getHashModel(key string, field string) *FSModel {
	bytes := MsgRedis.FSRedis.MsgRedisHashGet(key, field)
	if lib.StringIsNullOrEmpty(bytes) {
		return nil
	} else {
		last3Model := FSModel{}
		fmtByte := []byte(*bytes)
		err := ffjson.Unmarshal(fmtByte, &last3Model)
		if err != nil {
			return nil
		} else {
			return &last3Model
		}
	}
}
func jsonToZXModel(str *string) *ZXModel {
	if lib.StringIsNullOrEmpty(str) {
		return nil
	} else {
		last3Model := ZXModel{}
		err := ffjson.Unmarshal([]byte(*str), &last3Model)
		if err != nil {
			return nil
		} else {
			return &last3Model
		}
	}
}
func jsonToRKModel(str *string) *RKModel {
	if lib.StringIsNullOrEmpty(str) {
		return nil
	} else {
		rkModel := RKModel{}
		err := ffjson.Unmarshal([]byte(*str), &rkModel)
		if err != nil {
			return nil
		} else {
			return &rkModel
		}
	}
}
func jsonToZKModel(str *string) *ZKModel {
	if lib.StringIsNullOrEmpty(str) {
		return nil
	} else {
		zkModel := ZKModel{}
		err := ffjson.Unmarshal([]byte(*str), &zkModel)
		if err != nil {
			return nil
		} else {
			return &zkModel
		}
	}
}
func jsonToYKModel(str *string) *YKModel {
	if lib.StringIsNullOrEmpty(str) {
		return nil
	} else {
		ykModel := YKModel{}
		err := ffjson.Unmarshal([]byte(*str), &ykModel)
		if err != nil {
			return nil
		} else {
			return &ykModel
		}
	}
}
func floatToSting(val float64) string {
	return strconv.FormatFloat(val, 'f', -1, 64)
}
func floatToDot2Sting(val float64) string {
	return fmt.Sprintf("% .2f", val)
}
func stringToFloat(val string) float64 {
	output, _ := strconv.ParseFloat(strings.TrimSpace(val), 64)
	return float64(output)
}
