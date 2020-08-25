package config

import (
	"flag"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// YamlFile 配置文件路径
var YamlFile = flag.String("c", "", "配置文件路径")

// XlsxFile excel表格路径
var XlsxFile = flag.String("f", "", "xlsx文件路径")

// Sheet 需要处理的sheet
var Sheet = flag.String("s", "Sheet1", "指定要读取的Sheet")

// Start 忽略开头的行数
var Start = flag.Int("d", 1, "指定删除几行表头")

// Val 属性的值
type Val struct {
	Axis   string   `yaml:"axis"`
	Filter []string `yaml:"filter,omitempty"`
}

// Field 每个类的模型
type Field struct {
	Label string `yaml:"label"`
	Value Val    `yaml:"value"`
}

// Model 模型
type Model struct {
	Class  string  `yaml:"class"`
	Key    string  `yaml:"key"`
	Fields []Field `yaml:"fields"`
}

// Config 配置文件
type Config struct {
	Items []Model `yaml:"model"`
}

// Init 解析配置文件
func (c *Config) Init() error {
	flag.Parse()

	cfg, err := ioutil.ReadFile(*YamlFile)
	if err != nil {
		log.Printf("yamlFile.Get err #%v", err)
	}

	err = yaml.Unmarshal(cfg, c)

	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return err
}
