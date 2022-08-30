package goback

import (
	"fmt"

	"github.com/JekaTatsiy/goback/components/conf"
	"github.com/JekaTatsiy/goback/components/data"
	"github.com/JekaTatsiy/goback/components/work"
)

type Back struct {
	Config   conf.Component
	Database data.Component
	Work     work.Component
}

func Build() *Back {
	return &Back{}
}

func (b *Back) InitConfig(t interface{}, name string) *Back {
	fmt.Println(t)

	return b
}
func (b *Back) InitDatabase(name string) *Back {

	return b
}
func (b *Back) AddWorks(name string, action func(Back) error) *Back {

	return b
}

func (b *Back) Start() error {
return nil
}
