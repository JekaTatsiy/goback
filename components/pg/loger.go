package pg

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// GormLogger struct
type GormLogger struct {
	Logger logrus.Entry
}

// Print - Log Formatter
func (g *GormLogger) Printf(s string, v ...interface{}) {
	fmt.Println(s, v)
	switch s {
	case "sql":
		g.Logger.WithFields(
			logrus.Fields{
				"module":        "gorm",
				"type":          "sql",
				"rows_returned": v[4],
				"src":           v[0],
				"values":        v[3],
				"duration":      v[1],
			},
		).Info(v[2])
	case "log":
		g.Logger.WithFields(logrus.Fields{"module": "gorm", "type": "log"}).Print(v[1])
	}
}
