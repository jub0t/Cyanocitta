package structs

import "gorm.io/gorm"

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
