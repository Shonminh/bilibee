package http

import (
	"fmt"
	"strings"
)

func GetUrl(address string, path string, param ...interface{}) string {
	if strings.HasSuffix(address, "/") {
		address = address[:len(address)-1]
	}
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	return fmt.Sprintf(address+"/"+path, param...)
}
