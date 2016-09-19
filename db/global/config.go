package global

import (
	"github.com/Unknwon/goconfig"
)

var ConfigFile *goconfig.ConfigFile

func LoadConfigFile(fileName string) *goconfig.ConfigFile {
	c, _ := goconfig.LoadConfigFile(fileName)
	return c
}

func GetSectionKeyValue(c *goconfig.ConfigFile, section, key, defaultvalue string) string {
	return c.MustValue(section, key, defaultvalue)
}
