package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"net"
	"strings"
	"sync"
	"time"
)

var RecLength int64 = 0
var ErrLength int64 = 0

type HeadModel struct {
	MsgType   int32
	MsgLength int32
}
type TailModel struct {
	CheckSum int32
}
type MsgModel struct {
	MsgHead     []byte
	MsgContent  []byte
	MsgCheckSum []byte
}
type BaseMsgModel struct {
	MsgType    int32
	MsgLength  int32
	MsgContent []byte
}
type LogOnModel struct {
	HeadModel
	SenderCompID     string
	TargetCompID     string
	HeartBtInt       int32
	Password         string
	DefaultApplVerID string
	TailModel
}
type redisModel struct {
	zxRedis redis.Conn
	rkRedis redis.Conn
	fsRedis redis.Conn

	zkRedis redis.Conn
	ykRedis redis.Conn
}

var redisConn redisModel
var MsgTypes []int32 = []int32{
	1,
	3,
	8,
	390013,
	390019,
	390012,
	309011,
	309111,
	300111,
	300192,
	300191,
	300611,
	300592,
	300792,
	300591,
	300791,
	306311,
	390090,
	390095,
	390094,
	390093,
}
var err error
var redisLock sync.Mutex
var _redis redis.Conn

func main() {
	pwOption := redis.DialPassword("123qwe")
	dbOption := redis.DialDatabase(0)
	redisConn.fsRedis, err = redis.Dial("tcp", "127.0.0.1:6379", pwOption, dbOption)
	CheckError(err)
	_client, err := net.Dial("tcp", ":8088")
	CheckError(err)
	defer _client.Close()
	go handMsgAsync(_client)
	for {
		fmt.Println("-------------------")
		str := fmt.Sprintf("读取数据数量 %d", RecLength)
		fmt.Println(str)
		str = fmt.Sprintf("错误数据数量 %d", ErrLength)
		fmt.Println(str)
		time.Sleep(2 * time.Second)
	}
}
func CheckError(err error) {
	if err != nil {
		fmt.Printf(err.Error())
	}
}

//socket接受函数
func handMsgAsync(con net.Conn) {
	buffer := make([]byte, 4096000)
	con.Write(SendLogOn())
	go sendHeartBt(con)
	var unCompleteBytes []byte
	for {
		readLength, err := con.Read(buffer)
		CheckError(err)
		var readBuf []byte = buffer[0:readLength]
		//获取上次为解析的半包并装载到此次循环中解析
		if len(unCompleteBytes) != 0 {
			readBuf = BytesCopy(unCompleteBytes, readBuf)
		}
		startTime := time.Now()
		msgList := readAMessage(&readBuf) //拆包分解，并把半包数据留在下次循环处理
		go translateMsgAsync(&msgList)    //同步没有问题 异步就会出现数组数据解析出现乱码
		endTime := time.Since(startTime)
		fmt.Println("数据处理时间: ", endTime)
		RecLength += int64(len(msgList))
		unCompleteBytes = readBuf
	}
}

//返回发送登陆的数据
func SendLogOn() []byte {
	var logon LogOnModel
	var logonBytes MsgModel
	logon.MsgType = 1
	logon.SenderCompID = PadRight("admin", 20, " ")
	logon.TargetCompID = PadRight("customer", 20, " ")
	logon.HeartBtInt = 3
	logon.Password = PadRight("123qwe", 16, " ")
	logon.DefaultApplVerID = PadRight("1.01", 32, " ")

	// intbyte := int32ToBytes(logon.MsgType)

	content := BytesCopy(
		[]byte(logon.SenderCompID),
		[]byte(logon.TargetCompID),
		Int32ToBytes(logon.HeartBtInt),
		[]byte(logon.Password),
		[]byte(logon.DefaultApplVerID))

	logon.MsgLength = int32(len(content))
	logonBytes.MsgHead = BytesCopy(Int32ToBytes(logon.MsgType), Int32ToBytes(logon.MsgLength))
	logonBytes.MsgContent = content
	logon.CheckSum = CheckSum(BytesCopy(logonBytes.MsgHead, logonBytes.MsgContent))
	logonBytes.MsgCheckSum = Int32ToBytes(logon.CheckSum)
	// byteMsgType := reverseBytes(intbyte)
	outputbytes := BytesCopy(
		logonBytes.MsgHead,
		logonBytes.MsgContent,
		logonBytes.MsgCheckSum)
	return outputbytes
}

//拆包，半包留着下次解析
func readAMessage(input *[]byte) []BaseMsgModel {
	var output []BaseMsgModel
	for {
		byteArray := *input
		//头+长度+校验 最小12个字节，不足跳出
		if len(byteArray) < 12 {
			break
		}
		mt := BytesToInt32(byteArray[0:4])
		//获取数据头
		if Contain(mt, MsgTypes) {
			ml := BytesToInt32(byteArray[4:8])
			//数据长度
			if ml >= 0 && ml < 5000 {
				if len(byteArray) < 12+int(ml) {
					break
				}
				//获取报文content
				mc := byteArray[8 : 12+int(ml)]
				sum := BytesToInt32(mc[len(mc)-4 : len(mc)])
				checksum := CheckSum(byteArray[0 : 8+int(ml)])
				//检查校验
				if checksum == sum {
					var model BaseMsgModel
					model.MsgType = mt
					model.MsgLength = ml
					model.MsgContent = mc
					output = append(output, model)
					*input = byteArray[12+int(ml):]
				} else {
					*input = byteArray[12+int(ml):]
				}
			} else {
				//错误长度移除
				*input = byteArray[8:]
			}
		} else {
			//错误头部移除
			*input = byteArray[4:]
		}
	}
	return output
}

//发送心跳包
func sendHeartBt(con net.Conn) {
	var model MsgModel
	model.MsgHead = BytesCopy(Int32ToBytes(3), Int32ToBytes(0))
	model.MsgCheckSum = Int32ToBytes(CheckSum(model.MsgHead))
	outputbytes := BytesCopy(
		model.MsgHead,
		model.MsgCheckSum)
	for {
		con.Write(outputbytes)
		time.Sleep(3 * time.Second)
	}
}

//分发数据
func translateMsgAsync(list *[]BaseMsgModel) {
	modelList := *list
	msg300111Count := 0
	for i := 0; i < len(modelList); i++ {
		index := i
		switch modelList[i].MsgType {
		case 300111:
			go Save300111(&modelList[index])
			msg300111Count++
			break
		default:
			break
		}
	}

	fmt.Println("300111数据量: ", msg300111Count)
}

//解析300111号报文
func Save300111(input *BaseMsgModel) {
	//model := input
	bytes := input.MsgContent
	var hqModel HqModel
	dtNow := time.Now()
	hqModel.DtNow = dtNow.Format("2006-01-02 15:04:05")

	//0-8
	hqModel.RecTime = GetTimeFromFormatIntByte(bytes[0:8])
	hqModel.DataTime = hqModel.RecTime.Format("1504")
	hqModel.RecTimeStr = hqModel.RecTime.Format("200601021504")
	m1, _ := time.ParseDuration("-1m")
	hqModel.Upmin = hqModel.RecTime.Add(m1).Format("200601021504")
	hqModel.UnixTime = hqModel.RecTime.Unix()
	if hqModel.UnixTime < 0 {
		ErrLength++
		return
	}
	//8:10
	hqModel.ChannelNo = BytesToInt16(bytes[8:10])
	//10:13
	hqModel.MDStreamID = string(bytes[10:13])
	//13:21
	hqModel.SyNo = strings.TrimSpace(string(bytes[13:21]))
	//21:25
	hqModel.SySource = string(bytes[21:25])
	//25:33
	hqModel.TradingPhaseCode = string(bytes[25:33])
	//33:41
	hqModel.ZRSP = float32(BytesToInt64(bytes[33:41])) / 10000
	//41:49
	hqModel.CJBS = BytesToInt64(bytes[41:49])
	//49:57
	hqModel.CJSL = float32(BytesToInt64(bytes[49:57])) / 100
	//57:65
	hqModel.CJJE = float32(BytesToInt64(bytes[57:65])) / 10000
	//65:69
	hqCount := BytesToInt32(bytes[65:69])

	baseCount := 69

	for i := int32(0); i < hqCount; i++ {
		if (baseCount + 32) > len(bytes) {
			break
		}
		MDEntryType := strings.TrimSpace(string(bytes[baseCount : baseCount+2]))
		MDEntryPx := float32(BytesToInt64(bytes[baseCount+2:baseCount+10])) / 1000000
		MDEntrySize := float32(BytesToInt64(bytes[baseCount+10:baseCount+18])) / 100
		MDPriceLevel := BytesToInt16(bytes[baseCount+18 : baseCount+20])
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

	js, _ := json.Marshal(hqModel)
	jsstr := string(js)
	//fmt.Println(jsstr)
	zxRedisInsert(hqModel.SyNo, jsstr)
}

//redis存储
func zxRedisInsert(key string, val string) {
	redisLock.Lock()
	redisConn.fsRedis.Do("SET", key, val)
	redisLock.Unlock()
}
