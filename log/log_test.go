package log

import (
	"fmt"
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	logger, _ := NewLog4jLogger("test.log", Warn, 0, 0)
	logger.Info("info message")
	logger.Debug("debug message")
	logger.Warn("warn message")
	logger.Error("error message")
	<-time.After(1 * time.Second)
	logger.Close()
}

func TestBufLogger(t *testing.T) {
	logCh := make(chan string, 100)

	go func() {
		for {
			log, ok := <-logCh
			if !ok {
				return
			}
			fmt.Printf(log)
		}
	}()

	logger := NewLog4jBufLogger(logCh, Info)
	defer logger.Close()
	logger.Info("buf log: info message")
	logger.Debug("buf log: debug message")
	logger.Warn("buf log: warn message")
	logger.Error("buf log: error message")
}

func TestTermLogger(t *testing.T) {
	InitLogger(Debug)
	defer CloseLogger()
	Fatalf("good boy")
}
