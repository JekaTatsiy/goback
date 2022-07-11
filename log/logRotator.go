package golog

import (
	"fmt"
	"os"
	"time"

	ser "github.com/JekaTatsiy/goback/serv"
	"github.com/go-co-op/gocron"
	"github.com/sirupsen/logrus"
)

const (
	LogPath = "logs"
)

type CronExpr string

const (
	CRONMINUTE   CronExpr = "* * * * *"
	CRONHOUR     CronExpr = "0 * * * *"
	CRONMIDNIGHT CronExpr = "0 0 * * *"
)

func SetRotatebleLogger(serv ser.LogServerI, period CronExpr) error {
	go func() {
		s := gocron.NewScheduler(time.UTC)
		s.Cron(string(period)).Do(logRotator, serv)
		s.StartBlocking()
	}()

	return logRotator(serv)
}

func logRotator(serv ser.LogServerI) error {
	f, e := todayFile()
	if e != nil {
		if serv.GetLoger() != nil {
			serv.GetLoger().Info("new log file not created", e)
		} else {
			fmt.Println("new log file not created", e)
		}
		return e
	} else {
		if serv.GetLoger() != nil {
			serv.GetCurrentFile().Close()
		}
		serv.SetCurrentFile(f)
		serv.SetLoger(logrus.NewEntry(&logrus.Logger{
			Out:       f,
			Formatter: new(logrus.TextFormatter),
			Hooks:     make(logrus.LevelHooks),
			Level:     logrus.InfoLevel}))
		return nil
	}
}

func todayFile() (*os.File, error) {
	//date := time.Now().Format("2006-02-01 15:04:05")
	date := time.Now().Format("2006-01-02") // гггг-мм-дд
	fname := LogPath + string(os.PathSeparator) + date + ".log"

	f, e := os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)

	if e != nil {
		e = nil
		e = os.MkdirAll(LogPath, 0777)
		if e != nil {
			return nil, e
		}
		f, e = os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
		if e != nil {
			return nil, e
		}
	}
	return f, e
}
