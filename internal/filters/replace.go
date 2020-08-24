package filters

import (
	"strings"
)

func Replace(val string, param string, row map[string]string) string {
	rules := strings.Split(param, ",")
	for _, rule := range rules {
		r := strings.Split(rule, "=")
		if val == r[0] {
			return r[1]
		}
	}
	return val
}