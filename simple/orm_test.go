package main

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

type MyTestingT testing.T

var o orm.Ormer

func AssertEqual(exp interface{}, act interface{}, t *testing.T) {
    if exp != act {
        t.Error("expected ", exp, " actual ", act)
    }
}

func init() {
	o = orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
}

func Test_Insert(t *testing.T) {
	profile := new(Profile)
	profile.Age = 30

	user := new(User)
	user.Profile = profile
	user.Name = "slene"

	_, err := o.Insert(profile)
    AssertEqual(nil, err, t)
	_, err = o.Insert(user)
    AssertEqual(nil, err, t)
}

func Test_Read(t *testing.T) {
    user := User{Id: 1}
    err := o.Read(&user)
    AssertEqual(nil, err, t)
    AssertEqual("slene", user.Name, t)
}
