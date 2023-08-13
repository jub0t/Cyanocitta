package structs

import "gorm.io/gorm"

var (
	JsLang int8 = 1
	Pylang int8 = 2
	GoLang int8 = 3
)

type User struct {
	gorm.Model

	Username string
	Password string
	Id       string
	Created  string
}

type Process struct {
	gorm.Model

	Name string
	Id   string
}

type Bot struct {
	gorm.Model

	Name     string
	BotId    string
	Language int8
}
