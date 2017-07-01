package loger2

import (
    "sync"
    "testing"
    _ "time"
)

func TestLogger(t *testing.T) {
    Init2(3)
    defer Start(LogFilePath("./log"), EveryDay, AlsoStdout).Stop()
    defer Start2("loger_name1", LogFilePath("./log1"), EveryMinute, AlsoStdout).Stop2("loger_name1")
    defer Start2("loger_name2", LogFilePath("./log2"), EveryMinute, AlsoStdout).Stop2("loger_name2")
    Debugln2("loger_name1", "Ohch!")
    Infoln2("loger_name2", "Hello, Mike")
    Warnln("This might be painful but...")
    Errorln("You have to go through it until sunshine comes out")
    Infoln("Those were the days hard work forever pays")
}

func TestFileLoggerMultipleGoroutine(t *testing.T) {
    Init2(3)
    defer Start(LogFilePath("./log"), EveryDay, AlsoStdout).Stop()
    defer Start2("loger_name1", LogFilePath("./log1"), EveryMinute, AlsoStdout).Stop2("loger_name1")
    defer Start2("loger_name2", LogFilePath("./log2"), EveryMinute, AlsoStdout).Stop2("loger_name2")
    wg := sync.WaitGroup{}
    for i := 0; i < 10; i++ {
	wg.Add(1)
	go func() {
	    Debugln2("loger_name1", "Ohch!")
	    Infof2("loger_name2", "%s", "Wake up, Neo")
	    Warnf2("loger_name1", "%s", "The Matrix has you...")
	    Errorf2("loger_name2", "%s", "Follow the white rabbit")
	    Infof("%s", "Knock knock!")
	    wg.Done()
	}()
    }
    wg.Wait()
}
