package pg

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	err "github.com/JekaTatsiy/goback/err"
	ser "github.com/JekaTatsiy/goback/serv"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetDBConnection(serv ser.GormServer) error {
	var e error

	base, e := sql.Open("postgres", serv.GetDsn())
	if e != nil {
		serv.SetGormInnerError(err.FromError(e).WithPos(1))
		return e
	}

	errorWait := make(chan error)

	go func() {
		errorWait <- base.Ping()
	}()

	select {
	case e = <-errorWait:
		if e != nil {
			serv.SetGormInnerError(err.FromError(e).WithPos(1))
		}
	case <-time.After(time.Second):
		msg := fmt.Sprintf("failed to connect to database after 1 second with dsn='%s'", serv.GetDsn())
		err.FromMsgf(msg).WithPos(1).Info(logrus.NewEntry(logrus.StandardLogger()))
	}

	if e != nil {
		return e
	}

	db, e := gorm.Open(postgres.New(postgres.Config{Conn: base}), &gorm.Config{})
	if e != nil {
		serv.SetGormInnerError(err.FromError(e).WithPos(1))
		return e
	}
	serv.SetGorm(db)

	return nil
}
