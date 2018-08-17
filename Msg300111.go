package main

import (
	"time"
)

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
	ZRSP float32
	//成交笔数
	CJBS int64
	//成交数量
	CJSL float32
	//成交金额
	CJJE float32

	//买一价格 数量
	HQBJW1 float32
	HQBSL1 float32

	HQBJW2 float32
	HQBSL2 float32

	HQBJW3 float32
	HQBSL3 float32

	HQBJW4 float32
	HQBSL4 float32

	HQBJW5 float32
	HQBSL5 float32

	//卖一价格 数量
	HQSJW1 float32
	HQSSL1 float32

	HQSJW2 float32
	HQSSL2 float32

	HQSJW3 float32
	HQSSL3 float32

	HQSJW4 float32
	HQSSL4 float32

	HQSJW5 float32
	HQSSL5 float32

	//市盈率1
	HQSYL1 float32
	//市盈率2
	HQSYL2 float32

	//最近成交
	ZJCJ float32
	//开盘价
	KPJ float32
	//最高价
	ZGJ float32
	//最低价
	ZDJ float32

	//数据时间
	DataTime string
	//是否停牌·
	IsStop bool

	//现在时间
	DtNow string
	//上一分钟
	Upmin string
}
