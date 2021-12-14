package dao

import (
	"database/sql"
	"fmt"
	"strings"

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

//opaque写法
var notFoundCode = 40001
var systemErr = 50001

func QueryProductList(query string) error {

	tx := dbConn.Find(query)
	if tx.Error == sql.ErrNoRows {
		//在这一步封装好查询参数，这样DEBUG就能知道请求什么数据，没找到
		//同时带上了堆栈信息方便定位
		//我们没有仔细区别err是什么 就是告诉上游 出错了
		return fmt.Errorf("%d, not found", notFoundCode)
	}
	if tx.Error != nil {
		return fmt.Errorf("%d, not found", systemErr)
	}
	// go something
	return nil
}

func IsNoRow(err error) bool {
	return strings.HasPrefix(err.Error(), fmt.Sprintf("%d", notFoundCode))
}
