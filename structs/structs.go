package structs

import "gorm.io/gorm"

var (
	JsLang int8 = 1
	Pylang int8 = 2
	GoLang int8 = 3
)

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

type Response struct {
	Data    AnyData
	Success bool
	Message string
}

type AnyData map[string]interface{}

type User struct {
	gorm.Model

	Username string
	Password string
	Token    string
}
