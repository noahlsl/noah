package dbx

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func (c *Cfg) NewDB() *sqlx.DB {

	dsn := c.DataSource()
	db, err := sql.Open("mysql",
		dsn)

	if err != nil {
		panic(err)
	}
	if c.MaxOpen == 0 {
		c.MaxOpen = 20
	}
	if c.MaxIdle == 0 {
		c.MaxIdle = 10
	}
	db.SetMaxOpenConns(c.MaxOpen)
	db.SetMaxIdleConns(c.MaxIdle)

	return sqlx.NewDb(db, "mysql")
}

func (c *Cfg) DataSource() string {

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		c.Username, c.Password, c.Host, c.Port, c.Db)
}

func MustDB(dsn string, poolSize ...int) *sqlx.DB {

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}
	var (
		maxOpen = 100
		maxIdle = 50
	)
	if len(poolSize) > 0 {
		if poolSize[0] > 0 {
			maxOpen = poolSize[0]
			maxIdle = poolSize[0] / 2
		}
	}

	db.SetMaxOpenConns(maxOpen)
	db.SetMaxIdleConns(maxIdle)

	return db
}
