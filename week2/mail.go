package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
* @Author: mill
* @Date: 2021/12/9 7:09 下午
 */

//1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
// 答： 应该Warp异常，因为可以上层可以scan整个栈信息。
func initDB() *gorm.DB {
	dsn := "admin:admin@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	return db
}
func main() {
	var n string
	n = "bonu"
	u, err := dao(n)
	if err != nil {
		fmt.Println(errors.WithMessage(err, "not fount"))
	} else {
		fmt.Println(u)
	}
}

type user struct {
	id   uint
	name string
}

func dao(n string) (u user, err error) {
	db := initDB()
	err = db.Raw("SELECT id,name"+
		"FROM user").Where("name like ?", n).Scan(&u).Error
	if err == sql.ErrNoRows {
		return u, errors.Wrapf(errors.New("not found"), fmt.Sprintf("no user: %s", n))
	}
	if err != nil {
		return u, errors.Wrapf(err, fmt.Sprintf("system error: %s", n))
	}
	return u, nil
}
