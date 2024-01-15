package db

import (
	"time"

	"gorm.io/gorm/logger"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"

	"context"
	"fmt"
)

type Sqllogger struct {
	logger.Interface
}

func (l Sqllogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n===================================================]\n", sql)
}

func Postgres_init() *gorm.DB {
	
	dsn := "host=localhost user=postgres password=P@ssw0rd dbname=banktransfer port=5432 sslmode=disable TimeZone=Asia/Bangkok"

	dial := postgres.Open(dsn)

	db_Postgr, err := gorm.Open(dial, &gorm.Config{
		Logger: &Sqllogger{},
		// DryRun: true,
	})

	if err != nil {
		fmt.Println("Postgre Error: ", err)
		panic("Postgr Can not connect")
	}

	return db_Postgr
}
