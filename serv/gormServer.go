package serv

import (
	"fmt"

	"github.com/JekaTatsiy/goback/err"
	"gorm.io/gorm"
)

type GormServer interface {
	SetGormInnerError(*err.Err)
	GetGormInnerError() *err.Err
	GetDsn() string
	SetGorm(*gorm.DB)
	Gorm() *gorm.DB
}

type GormSimpleServer struct {
	InnerError     *err.Err
	ORM            *gorm.DB
	DatabaseString string
	Database       struct {
		Host string
		Port string
		User string
		Pass string
		Base string
	}
}

func (g *GormSimpleServer) SetDsnString(dbstring string) {
	g.DatabaseString = dbstring
}

func (g *GormSimpleServer) SetDsn(host, port, user, pass, dbname string) {
	g.Database.Host = host
	g.Database.Port = port
	g.Database.User = user
	g.Database.Pass = pass
	g.Database.Base = dbname

	testDSN := "user=%s password=%s host=%s port=%s dbname=%s sslmode=disable"
	g.DatabaseString = fmt.Sprintf(
		testDSN,
		g.Database.User,
		g.Database.Pass,
		g.Database.Host,
		g.Database.Port,
		g.Database.Base,
	)

}
func (g *GormSimpleServer) SetGormInnerError(e *err.Err) {
	g.InnerError = e
}

func (g *GormSimpleServer) GetGormInnerError() *err.Err {
	return g.InnerError
}

func (g *GormSimpleServer) GetDsn() string {
	return g.DatabaseString
}

func (g *GormSimpleServer) SetGorm(orm *gorm.DB) {
	g.ORM = orm
}

func (g *GormSimpleServer) Gorm() *gorm.DB {
	return g.ORM
}
