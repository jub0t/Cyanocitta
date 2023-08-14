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
	OwnerId     uint
	BotId       string
	AutoRestart bool
	MaxRestarts int8
	Language    int8
}

type (
	AnyData map[string]Any
	Any     interface{}
)

type ResponseAny struct {
	Success bool
	Message string `json:",omitempty"`
	Data    Any    `json:",omitempty"`
}

type Response struct {
	Success bool
	Message string  `json:",omitempty"`
	Data    AnyData `json:",omitempty"`
}

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
