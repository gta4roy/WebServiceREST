package util

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/magiconair/properties"
)

const (
	//Log Level ...
	LogLevel = "LOG_LEVEL"

	//Port ...
	Port = "PORT"

	//Host ...
	Host = "HOST"
)

var (
	propertyFile []string
	props        *properties.Properties
)

func readProperties(filepath string) {
	var propertyFile []string
	propertyFile = append(propertyFile, filepath)
	props, _ = properties.LoadFiles(propertyFile, properties.UTF8, true)
}

func init() {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	basepath := filepath.Dir(d)
	propertyFilePath := basepath + "/config/config.properties"
	readProperties(propertyFilePath)
}

func GetProperty(propertyName string, params ...string) string {
	msg, ok := props.Get(propertyName)
	if !ok {
		return props.MustGet("Property is not available in property file for the key " + propertyName)
	}

	placeHdrCnt := strings.Count(msg, "{")
	if placeHdrCnt == len(params) {
		for i, val := range params {
			replaceStr := fmt.Sprintf("%s%d%s", "{", i, "}")
			msg = strings.Replace(msg, replaceStr, val, -1)
		}
	}
	return msg
}
