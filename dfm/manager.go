package dfm

import (
	"disco/config"
	"disco/structs"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

var conf = config.GetConfig()

func MakeSpace(bot structs.Bot) bool {
	switch bot.Language {
	case structs.JsLang:
		{
			dir_path := filepath.Clean(fmt.Sprintf("%s%s", conf.StorePath, bot.BotId))
			if !Exists(dir_path) {
				if err := os.Mkdir(dir_path, fs.ModePerm); err != nil {
					fmt.Printf("Error Occured While Making Folder At %s\n", dir_path)
				}
			}

			return true
		}
	default:
		{
			return false
		}
	}
}
