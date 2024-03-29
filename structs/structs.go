package structs

import (
	"gorm.io/gorm"
)

var (
	JsLang int = 1
	PyLang int = 2
	GoLang int = 3

	Langauges = []int{JsLang, PyLang, GoLang}
)

type (
	AnyData map[string]Any
	Any     interface{}
)

type Bot struct {
	gorm.Model

	Name        string
	OwnerId     uint
	BotId       string
	AutoRestart bool
	MaxRestarts int8
	Language    int
	StartFile   string
}

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
	RestartOnStop bool
	MaxRestarts   int8
	Name          string
	CheckInterval int
	IndexFile     string
}
