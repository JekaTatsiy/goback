package serv

import (
	"fmt"

	"github.com/JekaTatsiy/goback/err"
	"gorm.io/gorm"
)

type GormServer interface {
	SetInnerError(*err.Err)
	GetInnerError() *err.Err
	GetDsn() string
	SetGorm(*gorm.DB)
	Gorm() *gorm.DB
}

type GormSimpleServer struct {
	InnerError *err.Err
	ORM        *gorm.DB
	Database   struct {
		Host   string
		Port   string
		User   string
		Pass   string
		DBName string
	}
}

func (g *GormSimpleServer) SetInnerError(e *err.Err) {
	g.InnerError = e
}

func (g *GormSimpleServer) GetInnerError() *err.Err {
	return g.InnerError
}

func (g *GormSimpleServer) GetDsn() string {
	testDSN := "user=%s password=%s host=%s port=%s dbname=%s sslmode=disable"
	return fmt.Sprintf(
		testDSN,
		g.Database.User,
		g.Database.Pass,
		g.Database.Host,
		g.Database.Port,
		g.Database.DBName,
	)
}

func (g *GormSimpleServer) SetGorm(orm *gorm.DB) {
	g.ORM = orm
}

func (g *GormSimpleServer) Gorm() *gorm.DB {
	return g.ORM
}

