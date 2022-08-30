package goback_test

import (
	goback "github.com/JekaTatsiy/goback"
	"testing"
	"time"
)

type Conf struct {
	Port     int
	Host     string
	Database struct {
		Pass string
	}
}

func TestCreate(t *testing.T) {

	server := goback.Build().
		InitConfig(Conf{}, "config_test.yaml").
		InitDatabase("localhost").
		AddWorks("http", func(b goback.Back) error {
			time.Sleep(time.Second * 5)
			return nil
		})

	server.Start()

}
