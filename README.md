# xlsx2itop

使用方法

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
  - name: name
    label: 名称
    value:
      axis: A
  - name: org_id
    label: 组织->名称
    value:
      axis: =Demo
  - name: status
    label: 状态
    value:
      axis: C
      filter:
      - replace(00=在线，01=已下线）
  - name: rack_id
    label: 机柜->全称
    value:
      axis: F
      filter:
      - join(E,F)
  - name: move2production
    label: 投产日期
    value:
      axis: G
      filter:
      - date()
```

目前支持 3 个过滤函数

- join，两列内容拼接
- replace，简单的替换
- date，日期格式转换为 iTop 默认格式