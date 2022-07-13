package serv

import (
	"os"

	"github.com/sirupsen/logrus"
)

type LogServer interface {
	SetLoger(*logrus.Entry)
	GetLoger() *logrus.Entry
	SetCurrentFile(*os.File)
	GetCurrentFile() *os.File
}

type LogSimpleServer struct {
	Log            *logrus.Entry
	LogCurrentFile *os.File
}

func (l *LogSimpleServer) SetLoger(entry *logrus.Entry) {
	l.Log = entry
}
func (l *LogSimpleServer) GetLoger() *logrus.Entry {
	return l.Log
}
func (l *LogSimpleServer) SetCurrentFile(f *os.File) {
	l.LogCurrentFile = f
}
func (l *LogSimpleServer) GetCurrentFile() *os.File {
	return l.LogCurrentFile
}
