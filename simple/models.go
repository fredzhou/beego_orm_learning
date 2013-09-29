package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type User struct {
	Id      int
	Name    string
	Profile *Profile `orm:"rel(one);on_delete(cascade)"` // OneToOne relation
}

type Profile struct {
	Id   int
	Age  int16
	User *User `orm:"reverse(one)"` // 设置反向关系(可选)
}

func init() {
	orm.RegisterDataBase("default", "sqlite3", "data.db", -1)

	orm.RegisterModel(new(User), new(Profile))

	name := "default"

	force_drop_table := true

	verbose := true

	err := orm.RunSyncdb(name, force_drop_table, verbose)
	if err != nil {
		fmt.Println(err)
	}
}
