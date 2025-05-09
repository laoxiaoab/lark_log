package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age int32
	Score int64
	Addr string
}

func main() {
	err := new(LarkLog).Init("./log/", "lark_log")
	if nil != err {
		fmt.Printf("init lark log failed err:%+v\n", err)
		panic(err)
	}
	fmt.Printf("init lark log success:%+v\n", err)

	/*
	startTs := time.Now().UnixMilli()
	logCnt := int(3000000)
	for i:=0; i<logCnt; i++ {
		str := fmt.Sprintf("are you ok today %d", i)
		DebugLogBufIo(str)
	}
	endTs := time.Now().UnixMilli()
	fmt.Printf("total log count: %+v, useTime:[%d]ms, write one log time: %+v[ms]\n", logCnt, endTs-startTs, float64(endTs-startTs)/float64(logCnt))
	*/

	/*
	year := time.Now().Year()
	mon := time.Now().Month()
	name := "nancy"
	startTs := time.Now().UnixMilli()
	logCnt := int(10)
	var person Person
	for i:=0; i<logCnt; i++ {
		person.Name = fmt.Sprintf("name:%+v", i)
		person.Addr = fmt.Sprintf("addr:%+v", i)
		person.Age = int32(i)
		person.Score = int64(rand.Int())
		str := fmt.Sprintf("are you ok today %d", i)
		DebugLogMultiVariable("multi variable log, year:%+v, month:%+v, name:%+v, personInfo:%+v, str:%+v", year, mon, name, person, str)
	}
	endTs := time.Now().UnixMilli()
	fmt.Printf("total log count: %+v, useTime:[%d]ms, write one log time: %+v[ms]\n", logCnt, endTs-startTs, float64(endTs-startTs)/float64(logCnt))
	*/

	//不带参数的日志
	DebugLogMultiVariable("multi variable log, log-1")
	DebugLogMultiVariable("multi variable log, log-2")
	DebugLogMultiVariable("multi variable log, log-3")

	/*
	//多个goroutine写入
	startTs := time.Now().UnixMilli()
	var wg sync.WaitGroup
	wg.Add(3)
	logCnt := int(1000000)
	go func() {
		defer wg.Done()
		for i:=0; i<logCnt; i++ {
			str := fmt.Sprintf("are you ok today [one] %d", i)
			DebugLogBufIo(str)
		}
	}()
	go func() {
		defer wg.Done()
		for i:=0; i<logCnt; i++ {
			str := fmt.Sprintf("are you ok today [two] %d", i)
			DebugLogBufIo(str)
		}
	}()
	go func() {
		defer wg.Done()
		for i:=0; i<logCnt; i++ {
			str := fmt.Sprintf("are you ok today [three] %d", i)
			DebugLogBufIo(str)
		}
	}()
 	wg.Wait()
	endTs := time.Now().UnixMilli()
	fmt.Printf("total log count: %+v, useTime:[%d]ms, write one log time: %+v[ms]\n", logCnt, endTs-startTs, float64(endTs-startTs)/float64(logCnt))
	*/
}
