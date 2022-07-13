package pg

import (
	"database/sql"
	"errors"
	"fmt"
	err "github.com/JekaTatsiy/goback/err"
	ser "github.com/JekaTatsiy/goback/serv"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetDBConnection(serv ser.GormServer) error {
	var e error

	base, e := sql.Open("postgres", serv.GetDsn())
	if e != nil {
		serv.SetInnerError(err.FromError(e))
		return e
	}

	errorWait := make(chan error)

	go func() {
		errorWait <- base.Ping()
	}()

	select {
	case e = <-errorWait:
		if e != nil {
			serv.SetInnerError(err.FromError(e))
		}
	case <-time.After(time.Second):
		msg := fmt.Sprintf("failed to connect to database after 1 second with dsn='%s'", serv.GetDsn())
		serv.SetInnerError(err.FromMsgf(msg))
		e = errors.New(msg)
	}

	if e != nil {
		return e
	}

	db, e := gorm.Open(postgres.New(postgres.Config{Conn: base}), &gorm.Config{})
	if e != nil {
		serv.SetInnerError(err.FromError(e))
		return e
	}
	serv.SetGorm(db)

	return nil
}
