package dfm

import (
	"disco/config"
	"disco/structs"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func MakeSpace(bot structs.Bot) bool {
	dir_path := filepath.Clean(fmt.Sprintf("%s%s", config.C.StorePath, bot.BotId))
	println(dir_path)

	if !Exists(dir_path) {
		if err := os.Mkdir(dir_path, fs.ModePerm); err != nil {
			fmt.Printf("Error Occured While Making Folder At %s: %s\n", dir_path, err)
		}
	}

	return true
}

func PrepareSpace() {
	dir_path := filepath.Clean(config.C.StorePath)
	if !Exists(dir_path) {
		if err := os.Mkdir(dir_path, 0755); err != nil {
			fmt.Printf("Space Folder Allocation Failed [%s]: %s\n", dir_path, err)
		} else {
			fmt.Printf("Space Folder Created: %s\n", dir_path)
		}
	}
}

func StartFileByLanguage(lang int) string {
	switch lang {
	case structs.JsLang:
		{
			return "index.js"
		}
	case structs.GoLang:
		{
			return "main.go"
		}
	case structs.PyLang:
		{
			return "main.py"
		}
	default:
		{
			return ""
		}
	}
}
