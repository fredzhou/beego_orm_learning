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
	user := new(User)
	profile.Age = 30
    profile.User = user
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

	//AssertEqual(int16(30), user.Profile.Age, t)

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

func Test_Delete(t *testing.T) {
	user := &User{Id: 1}
	err := o.Read(user)
	AssertEqual(nil, err, t)
	profile := &Profile{Id: 1}
	err = o.Read(profile)
	AssertEqual(nil, err, t)

	o.Delete(profile)

	err = o.Read(profile)
	AssertEqual(true, nil != err, t)
    //verify on_delete(cascade)
	err = o.Read(user)
	AssertEqual(true, nil != err, t)
}
