package xlsx

import (
	"log"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// Index2Axis 将数字形式的数组下标转为 excel 字母形式坐标
// index 是下标，从 0 开始
func index2Axis(index int) (axis string) {
	first := index / 26
	last := index % 26
	var prefix rune
	var postfix rune
	// ASCII 码 A->65
	postfix = rune(last + 65)
	if first > 0 {
		// 大于 26 列时才加前缀，此时 first 大于 1，转换为字符时应少加 1，即 64
		prefix = rune(first + 64)
		axis = string(prefix) + string(postfix)
	} else {
		axis = string(postfix)
	}
	return axis
}

// GetData 将excel表格每一行转换为以 axis 为key的map
func GetData(xlsxFile string, sheet string, start int) (data []map[string]string) {
	f, err := excelize.OpenFile(xlsxFile)
	if err != nil {
		log.Fatalf("OpenExcel Failed: %v", err.Error())
	}

	rows, err := f.GetRows(sheet)
	// 删除表头
	rows = rows[start:]
	//fmt.Println(reflect.TypeOf(rows))
	for _, row := range rows {
		var item map[string]string
		item = make(map[string]string)
		for col, colCell := range row {
			axis := index2Axis(col)
			item[axis] = colCell
		}
		data = append(data, item)
	}
	return data
}
