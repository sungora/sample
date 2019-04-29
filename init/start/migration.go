package start

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sungora/app/connect"
	"gopkg.in/webnice/migrate.v1/goose"
)

// миграция БД из командной строки
// https://github.com/webnice/migrate
// todo gsmigrate --dir="." --drv="postgres" --dsn="host=localhost port=5432 user=postgres password=postgres dbname=teleport sslmode=disable" create add_some_column sql
// Восстановление БД из текстового дампа
// /usr/bin/psql -h "localhost" -p "5432" -U "postgres" -w -d "test" -f "/home/.../test.sql"

// Migration миграция бд
func Migration(path string, version int64) (ver int64, err error) {
	if version == 0 {
		return goose.GetDBVersion(connect.GetDB().DB())
	}
	// UP
	if version > 0 {
		if version == 1 {
			err = goose.Up(connect.GetDB().DB(), path)
		} else {
			err = goose.UpTo(connect.GetDB().DB(), path, version)
		}
		// DOWN
	} else {
		if version == -1 {
			err = goose.Down(connect.GetDB().DB(), path)

		} else {
			err = goose.DownTo(connect.GetDB().DB(), path, version*-1)
		}

	}
	if err != nil {
		return 0, err
	}
	return goose.GetDBVersion(connect.GetDB().DB())
}
