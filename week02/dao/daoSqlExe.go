package dao

import (
	"database/sql"

	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

/*
当dao 层中当遇到一个 sql.ErrNoRows 的时候，直接向上抛还是如何
*/

func init() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	dbConn = db
}

func QueryProductList() ([]Product, error) {
	product := make([]Product, 0)
	tx := dbConn.Find(&product)
	if tx.Error == sql.ErrNoRows {
		//由于sql.ErrNoRows相当于没有查询到符合条件的数据 并不算错误
		return product, nil
	}
	if tx.Error != nil {
		return product, errors.Wrap(tx.Error, "获取产品数据集合调用数据库异常")
	}
	return product, nil
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}
