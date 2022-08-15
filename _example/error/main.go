package main

import (
	"errors"
	"os"

	"github.com/JekaTatsiy/goback/err"
	"github.com/sirupsen/logrus"
)

func main() {
	entry := logrus.NewEntry(&logrus.Logger{
		Out:       os.Stderr,
		Level:     logrus.DebugLevel,
		Formatter: new(logrus.TextFormatter),
	})

	err.FromMsg("i'm stoped").Err(entry)                             // yes. with message
	err.FromMsg("i'm stoped with code").WithCode(1300).Err(entry)    // yes. with message and code
	err.FromCode(1300).WithMsg("i'm stoped with code").Err(entry)    // no.  if the code is first, then code shoud be from standart codes
	err.FromCode(2000).WithMsg("Database not responding").Err(entry) // yes. with standart code
	err.FromError(errors.New("i'm fall")).Err(entry)                 // yes. with error
}
