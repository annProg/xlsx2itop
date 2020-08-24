package main

import (
	"bytes"
	"encoding/csv"
	"log"
	"os"
	"regexp"
	"strings"
	"xlsx2itop/internal/config"
	"xlsx2itop/internal/reader/xlsx"
	"xlsx2itop/internal/filters"
	//"reflect"
)

var cfg config.Config

// Data is excel data
var Data []map[string]string

type field struct {
	key    string
	filter []string
}

// Callback 回调 filter 函数，处理单元格数据
// type Callback func(val string, param string) string

// func filterCallback(val string, param string, callback Callback) string {
// 	return callback(val, param)
// }

// 以 = 开头的key是字符串值，直接赋值为 = 后面的值，不使用excel中的值
func isStrValue(key string) string {
	value := ""
	r := regexp.MustCompile(`^=(.*)$`)
	match := r.FindStringSubmatch(key)
	if len(match) > 1 {
		value = match[1]
	}
	return value
}

// 处理单元格数据
func doFilter(val string, filter []string, row map[string]string) string {
	// 删除空格，回车，TAB等
	val = strings.TrimSpace(val)
	for _, f := range filter {
		r := regexp.MustCompile(`^([a-z]+)\((.*)\)$`)
		match := r.FindStringSubmatch(f)
		if len(match) > 2 {
			switch match[1] {
			case "date":
				val = filters.Date(val, match[2], row)
			case "replace":
				return filters.Replace(val, match[2], row)
			case "join":
				return filters.Join(val, match[2], row)
			default:
				log.Fatalf("Filter function not supported: %v", match[1])
			}
		}
	}
	// 没有 filter 函数的保持不变
	return val
}

func main() {
	err := cfg.Init()
	if err != nil {
		log.Fatalf("Load Config Error: %v", err.Error())
	}

	Data = xlsx.GetData(*config.XlsxFile, *config.Sheet)
	//log.Printf("Excel Data: %v", Data)

	for _, conf := range cfg.Items {
		out := "import-" + conf.Class + ".csv"
		buf := new(bytes.Buffer)
		w := csv.NewWriter(buf)

		var label []string
		var fields []field
		for _, l := range conf.Fields {
			label = append(label, l.Label)
			fields = append(fields, field{l.Value.Axis, l.Value.Filter})
		}

		w.Write(label)
		w.Flush()

		for _, val := range Data {
			log.Printf("Excel item: %v", val)
			var item []string
			for _, f := range fields {
				v := isStrValue(f.key)
				if v != "" {
					item = append(item, v)
					continue
				}
				if _, ok := val[f.key]; !ok {
					log.Fatalf("key %v not found in Excel. Please check config", f.key)
				}

				v = doFilter(val[f.key], f.filter, val)

				item = append(item, v)
			}

			w.Write(item)
			w.Flush()
		}

		fout, err := os.Create(out)
		defer fout.Close()
		if err != nil {
			log.Fatalf("Create csv output failed: %v", err.Error())
		}

		fout.WriteString(buf.String())
	}
}
