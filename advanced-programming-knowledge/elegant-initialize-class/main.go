package main

import "fmt"

type User struct {
	name string
	id   int
	age  int
	db   interface{}
}

// UserOption 代表可选参数
type UserOption func(user *User)

// WithName 代表Name为可选参数
func WithName(name string) UserOption {
	return func(user *User) {
		user.name = name
	}
}

// WithAge 代表age为可选参数
func WithAge(age int) UserOption {
	return func(user *User) {
		user.age = age
	}
}

// WithDB 代表db为可选参数
func WithDB(db interface{}) UserOption {
	return func(user *User) {
		user.db = db
	}
}

// NewUser 代表初始化
func NewUser(id int, options ...UserOption) *User {
	user := &User{
		name: "default",
		id:   id,
		age:  10,
		db:   nil,
	}
	for _, option := range options {
		option(user)
	}
	return user
}

func main() {
	// 参数分为两部分，一部分必填，一部分选填
	user := NewUser(1, WithAge(15), WithName("Alice"))
	WithDB("DB instance")(user)
	fmt.Println(*user)
}
