# xlsx2itop

使用方法

```
$ ./bin/xlsx2itop -h
Usage of xlsx2itop:
  -c string
        配置文件路径
  -d int
        指定删除几行表头 (default 1)
  -f string
        xlsx 文件路径
  -s string
        指定要读取的 Sheet (default "Sheet1")
```

例如：
```
xlsx2itop -c config.yaml -f CI.xlsx -s SheetName
```

配置文件示例

```
---
date: 2020-07-05
model:
- class: Server
  key: serialnumber
  fields:
  - label: 名称
    value:
      axis: A
  - label: 组织->名称
    value:
      axis: =Demo
  - label: 状态
    value:
      axis: C
      filter:
      - replace(00=在线，01=已下线）
  - label: 机柜->全称
    value:
      axis: F
      filter:
      - join(E,F)
  - label: 投产日期
    value:
      axis: G
      filter:
      - date()
```

目前支持 3 个过滤函数

- join，两列内容拼接
- replace，简单的替换
- date，日期格式转换为 iTop 默认格式