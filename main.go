package main

import (
	"StockSocket/cache"
	"StockSocket/lib"
	"StockSocket/msg"
	"fmt"
	"net"
	"runtime"
	"sync"
	"time"
)

var dataLock sync.Mutex

type SystemParams struct {
	RecLength       int64
	ErrLength       int64
	msg300111Length int64
	Params          sync.Mutex
}

var systemParams SystemParams

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
var msgChannel chan []byte = make(chan []byte, 100000)
var socketStateChan chan bool
var (
	msgTempDb  = 15
	msgTempKey = "msgQueue300111"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//创建redis连接
	//cache.InitRedis()
	//cache.MsgRedisListRemove(msgTempKey, msgTempDb)

	//创建socket连接
	_client, err := net.Dial("tcp", "127.0.0.1:6666")
	CheckError(err, _client)
	defer _client.Close()

	//初始化有效日期
	msg.InitValidaTime()
	msg.InitRedisPool()
	//轮询读取socket数据HQSEDSLCZ
	go handMsgAsync(_client)
	//go translateMsg300111(0)
	for {

		fmt.Println("-------------------")
		str := fmt.Sprintf("读取数据数量 %d", systemParams.RecLength)
		fmt.Println(str)
		str = fmt.Sprintf("300111数据个数 %d", systemParams.msg300111Length)
		fmt.Println(str)
		systemParams.Params.Lock()
		systemParams.RecLength = 0
		systemParams.ErrLength = 0
		systemParams.msg300111Length = 0
		systemParams.Params.Unlock()
		time.Sleep(1 * time.Second)

	}
}

func CheckError(err error, con net.Conn) {
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Socket连接错误，正在重连")
		con.Close()
		con, err = net.Dial("tcp", "127.0.0.1:6666")
		time.Sleep(time.Second)
		CheckError(err, con)
	}
}

//socket接受函数
func handMsgAsync(con net.Conn) {
	buffer := make([]byte, 4096000)
	con.Write(SendLogOn())
	go sendHeartBt(con)
	var unCompleteBytes []byte

	for {
		//msg300111Count := 0
		readLength, err := con.Read(buffer)
		CheckError(err, con)
		var readBuf []byte = buffer[0:readLength]
		//获取上次为解析的半包并装载到此次循环中解析
		if len(unCompleteBytes) != 0 {
			readBuf = lib.BytesCopy(unCompleteBytes, readBuf)
		}
		//startTime := time.Now()
		msgList := readAMessage(&readBuf) //拆包分解，并把半包数据留在下次循环处理

		/*systemParams.Params.Lock()
		systemParams.RecLength += int64(len(*msgList)) //接受长度
		systemParams.Params.Unlock()*/

		//var msgQueue *[]BaseMsgModel
		//msgQueue = msgList
		translateMsg300111(msgList)

		/*for _, item := range *msgList {
			switch item.MsgType {
			case 300111:
				//cache.Pipeline("RPUSH", msgTempKey, item.MsgContent, msgTempDb)
				cache.MsgRedisListRightAppend(msgTempKey, item.MsgContent, msgTempDb)
				systemParams.msg300111Length++
				msg300111Count++
				break
			default:
				break
			}
		}
		if msg300111Count != 0 {
			go translateMsgQueue(msg300111Count)
		}*/
		unCompleteBytes = readBuf
		//translateMsgAsync(msgList) //同步没有问题 异步就会出现数组数据解析出现乱码
		//go translateMsgQueue(msg300111Count)
		//go translateMsg300111(msg300111Count)
		//msg300111Count = 0
		//endTime := time.Since(startTime)
		//fmt.Println("读取并数据处理时间: ", endTime)

	}
}

func readMsgAsync(con net.Conn, unCompleteBytes []byte) {
	buffer := make([]byte, 4096000)
	con.Write(SendLogOn())
	readLength, err := con.Read(buffer)
	CheckError(err, con)
	var readBuf = buffer[0:readLength]
	//获取上次为解析的半包并装载到此次循环中解析
	if len(unCompleteBytes) != 0 {
		readBuf = lib.BytesCopy(unCompleteBytes, readBuf)
	}

	msgList := readAMessage(&readBuf)              //拆包分解，并把半包数据留在下次循环处理
	systemParams.RecLength += int64(len(*msgList)) //接受长度
	readMsgAsync(con, unCompleteBytes)
}

//返回发送登陆的数据
func SendLogOn() []byte {
	var logon LogOnModel
	var logonBytes MsgModel
	logon.MsgType = 1
	logon.SenderCompID = lib.PadRight("admin", 20, " ")
	logon.TargetCompID = lib.PadRight("customer", 20, " ")
	logon.HeartBtInt = 3
	logon.Password = lib.PadRight("123qwe", 16, " ")
	logon.DefaultApplVerID = lib.PadRight("1.01", 32, " ")

	// intbyte := int32ToBytes(logon.MsgType)

	content := lib.BytesCopy(
		[]byte(logon.SenderCompID),
		[]byte(logon.TargetCompID),
		lib.Int32ToBytes(logon.HeartBtInt),
		[]byte(logon.Password),
		[]byte(logon.DefaultApplVerID))

	logon.MsgLength = int32(len(content))
	logonBytes.MsgHead = lib.BytesCopy(lib.Int32ToBytes(logon.MsgType), lib.Int32ToBytes(logon.MsgLength))
	logonBytes.MsgContent = content
	logon.CheckSum = lib.CheckSum(lib.BytesCopy(logonBytes.MsgHead, logonBytes.MsgContent))
	logonBytes.MsgCheckSum = lib.Int32ToBytes(logon.CheckSum)
	// byteMsgType := reverseBytes(intbyte)
	outputbytes := lib.BytesCopy(
		logonBytes.MsgHead,
		logonBytes.MsgContent,
		logonBytes.MsgCheckSum)
	return outputbytes
}

//拆包，半包留着下次解析
func readAMessage(input *[]byte) *[]BaseMsgModel {
	var output []BaseMsgModel
	for {
		byteArray := *input
		//头+长度+校验 最小12个字节，不足跳出
		if len(byteArray) < 12 {
			break
		}
		mt := lib.BytesToInt32(byteArray[0:4])
		//获取数据头
		if lib.Contain(mt, MsgTypes) {
			ml := lib.BytesToInt32(byteArray[4:8])
			//数据长度
			if ml >= 0 && ml < 5000 {
				if len(byteArray) < 12+int(ml) {
					break
				}
				//获取报文content
				mc := byteArray[8 : 12+int(ml)]
				sum := lib.BytesToInt32(mc[len(mc)-4 : len(mc)])
				checksum := lib.CheckSum(byteArray[0 : 8+int(ml)])
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
	return &output
}

//发送心跳包
func sendHeartBt(con net.Conn) {
	var model MsgModel
	model.MsgHead = lib.BytesCopy(lib.Int32ToBytes(3), lib.Int32ToBytes(0))
	model.MsgCheckSum = lib.Int32ToBytes(lib.CheckSum(model.MsgHead))
	outputbytes := lib.BytesCopy(
		model.MsgHead,
		model.MsgCheckSum)
	for {
		con.Write(outputbytes)
		time.Sleep(3 * time.Second)
	}
}

//分发数据
func translateMsgQueue(msgCount int) {
	for i := 0; i < msgCount; i++ {
		msgContent := cache.MsgRedisListLeftPop(msgTempKey, msgTempDb)
		if msgContent != nil {
			go msg.Save300111(*msgContent)
		}
	}
	/*for {
		item := <-msgChannel
		msg.Save300111(item)
		//fmt.Println("数据长度%d", item)
	}*/
}

//分发数据
func translateMsg300111(msgList *[]BaseMsgModel) {
	for _, item := range *msgList {
		switch item.MsgType {
		case 300111:
			msg.Save300111(item.MsgContent)
			systemParams.Params.Lock()
			systemParams.msg300111Length++
			systemParams.Params.Unlock()
			break
		default:
			break
		}
		systemParams.Params.Lock()
		systemParams.RecLength++
		systemParams.Params.Unlock()
	}
}
