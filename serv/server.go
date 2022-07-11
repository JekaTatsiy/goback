package golog

import (
	"os"

	"github.com/sirupsen/logrus"
)

type LogServerI interface {
	SetLoger(*logrus.Entry)
	GetLoger() *logrus.Entry
	SetCurrentFile(*os.File)
	GetCurrentFile() *os.File
}

type LogServer struct {
	Log            *logrus.Entry
	LogCurrentFile *os.File
}

func (l *LogServer) SetLoger(entry *logrus.Entry) {
	l.Log = entry
}
func (l *LogServer) GetLoger() *logrus.Entry {
	return l.Log
}
func (l *LogServer) SetCurrentFile(f *os.File) {
	l.LogCurrentFile = f
}
func (l *LogServer) GetCurrentFile() *os.File {
	return l.LogCurrentFile
}
