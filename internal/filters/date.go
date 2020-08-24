package filters

import (
	"strings"
)

// Date 转换时间格式为 iTop 标准 yyyy-mm-dd
func Date(val string, param string, row map[string]string) string {
	dt := strings.Split(val, "-")
	return "20" + dt[2] + "-" + dt[0] + "-" + dt[1]
}
