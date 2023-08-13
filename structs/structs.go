package structs

import (
	"gorm.io/gorm"
)

var (
	JsLang int8 = 1
	Pylang int8 = 2
	GoLang int8 = 3
)

type Bot struct {
	gorm.Model

	Name        string
	OwnerId     string
	BotId       string
	AutoRestart bool
	Language    int8
}

type Response struct {
	Success bool
	Message string  `json:",omitempty"`
	Data    AnyData `json:",omitempty"`
}

type AnyData map[string]interface{}

type User struct {
	gorm.Model

	Username string
	Password string
	Token    string
}

type NodeInstance struct {
	IndexFile     string
	RestartOnStop bool
	MaxRestarts   int8
	Name          string
	CheckInterval int
}
