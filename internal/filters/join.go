package filters

import (
	"strings"
)

func Join(val string, param string, row map[string]string) string {
	p := strings.Split(param, ",")
	return row[p[0]] + "-" + row[p[1]]
}