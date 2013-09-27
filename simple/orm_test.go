package main

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

var o orm.Ormer

func init() {
	o = orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
}

func Test_Insert(t *testing.T) {
	profile := new(Profile)
	profile.Age = 30
	_, err := o.Insert(profile)
	AssertEqual(nil, err, t)

	user := new(User)
	user.Profile = profile
	user.Name = "slene"

	_, err = o.Insert(user)
	AssertEqual(nil, err, t)
}

func Test_Read(t *testing.T) {
	user := User{Id: 1}
	err := o.Read(&user)
	AssertEqual(nil, err, t)
	AssertEqual("slene", user.Name, t)

    profile := Profile{Id: 1}
    err = o.Read(&profile)
    AssertEqual(int16(30), profile.Age, t)

	//query with non-primary key
	user = User{Name: "slene"}
	err = o.Read(&user, "Name")
	AssertEqual(nil, err, t)
	AssertEqual(1, user.Id, t)
}

func Test_Update(t *testing.T) {
	user := User{Id: 1}
    err := o.Read(&user)
    AssertEqual(nil, err, t)
    AssertEqual("slene", user.Name, t)

    user.Name = "Fred"
    _, err = o.Update(&user)
    AssertEqual(nil, err, t)

    user1 := User{Id: 1}
    err = o.Read(&user1)
    AssertEqual("Fred", user1.Name, t)

    //teardown
    user1.Name = "slene"
    o.Update(&user1)
}
