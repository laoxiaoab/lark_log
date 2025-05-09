package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

type LarkLog struct {
	logFilePath string
	logFileName string
	logFile     *os.File
}

var (
	logEntity *LarkLog
	logMux    sync.Mutex
)

func (m *LarkLog) Init(path string, fileName string) error {
	y, mon, d, h, min, sec, milli := CurrentTimeDetail()
	logFileName := fmt.Sprintf("%s%s-%04d%02d%02d-%02d-%02d-%02d-%04d.log",
		path, fileName, y, mon, d, h, min, sec, milli)
	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if nil != err {
		fmt.Printf("[error] log_file Init failed err:%+v\n", err)
		return err
	}
	fmt.Printf("[info] log_file Init success err:%+v\n", err)

	logEntity = &LarkLog{}
	logEntity.logFilePath = path
	logEntity.logFileName = fileName
	logEntity.logFile = file
	return nil
}

func (m *LarkLog) WriteLog(content string) {
	logMux.Lock()
	defer logMux.Unlock()

	file := logEntity.logFile
	_, err := file.WriteString(content)
	if nil != err {
		fmt.Printf("WriteLog write string failed err:%+v\n", err)
		return
	}
	fmt.Printf("\n")
}

func (m *LarkLog) WriteLogBufIo(content string) {
	logMux.Lock()
	defer logMux.Unlock()

	writer := bufio.NewWriter(logEntity.logFile)
	_, err := writer.WriteString(content)
	if nil != err {
		fmt.Printf("WriteLogBufIo write string failed err:%+v\n", err)
		return
	}
	writer.Flush()
}

func DebugLog(content string) {
	y, mon, d, h, min, sec, milli := CurrentTimeDetail()
	str := fmt.Sprintf("[debug] %04d-%02d-%02d %02d:%02d:%02d.%d %s\n", y, mon, d, h, min, sec, milli, content)
	logEntity.WriteLog(str)
}

func InfoLog(content string) {
	y, mon, d, h, min, sec, milli := CurrentTimeDetail()
	str := fmt.Sprintf("[debug] %04d-%02d-%02d %02d:%02d:%02d.%d %s\n", y, mon, d, h, min, sec, milli, content)
	logEntity.WriteLog(str)
}

func ErrorLog(content string) {
	y, mon, d, h, min, sec, milli := CurrentTimeDetail()
	str := fmt.Sprintf("[debug] %04d-%02d-%02d %02d:%02d:%02d.%d %s\n", y, mon, d, h, min, sec, milli, content)
	logEntity.WriteLog(str)
}

func DebugLogBufIo(content string) {
	y, mon, d, h, min, sec, milli := CurrentTimeDetail()
	str := fmt.Sprintf("[debug] %04d-%02d-%02d %02d:%02d:%02d.%d %s\n", y, mon, d, h, min, sec, milli, content)
	logEntity.WriteLogBufIo(str)
}

/**
**@brief: DebugLogMultiVariable, 变参形式的日志写入
**@test:
**/

func DebugLogMultiVariable(fmtStr string, param ...interface{}) {
	content := fmt.Sprintf(fmtStr, param...)
	y, mon, d, h, min, sec, milli := CurrentTimeDetail()
	str := fmt.Sprintf("[debug] %04d-%02d-%02d %02d:%02d:%02d.%d %s\n", y, mon, d, h, min, sec, milli, content)
	logEntity.WriteLogBufIo(str)
}
