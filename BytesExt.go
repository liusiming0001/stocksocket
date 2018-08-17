// StockSocket project doc.go

/*
StockSocket document
*/
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
)

func ReverseBytes(input []byte) []byte {
	temp := input
	for from, to := 0, len(temp)-1; from < to; from, to = from+1, to-1 {
		temp[from], temp[to] = temp[to], temp[from]
	}
	return temp
}
func Int8ToBytes(n int8) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n)
	return bytesBuffer.Bytes()
}
func Int16ToBytes(n int16) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n)
	return bytesBuffer.Bytes()
}
func Int32ToBytes(n int32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n)
	return bytesBuffer.Bytes()
}
func Int64ToBytes(n int64) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n)
	return bytesBuffer.Bytes()
}

func BytesToInt8(b []byte) int8 {
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int8
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return int8(tmp)
}
func BytesToInt16(b []byte) int16 {
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int16
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return int16(tmp)
}
func BytesToInt32(b []byte) int32 {
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int32
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return int32(tmp)
}
func BytesToInt64(b []byte) int64 {
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int64
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return int64(tmp)
}

func IntToString(inter interface{}) string {
	return fmt.Sprintf("%d", inter)
}

func BytesCopy(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}
func PadRight(str string, tarLen int, padStr string) string {
	strLen := tarLen - len(str)
	if strLen < 0 {
		return str
	}
	for i := 0; i < strLen; i++ {
		str += padStr
	}
	return str
}
func CheckSum(model []byte) int32 {
	var sum int32 = 0
	for i := 0; i < len(model); i++ {
		sum += int32(model[i])
	}

	return int32(sum % 256)
}
func Contain(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}

	return false
}
