package main

import "time"

func GetTimeFromFormatIntByte(bytes []byte) time.Time {
	Str := IntToString(BytesToInt64(bytes))
	OrigTime, err := time.Parse("20060102150405000", Str)
	if err != nil {
		return OrigTime
	} else {
		return time.Now()
	}
}
