#配置表生成go代码和json文件

使用导表工具是[https://github.com/davyxu/tabtoy]

## 首先先把配置表copy到目录下

index.xlsx必须要有，并且和gen.sh参数-index保持一致，另外index.xlsx配置表结构是固定
定义参考[https://github.com/davyxu/tabtoy/blob/master/v3/model/index.go]

type.xlsx是类型定义，其中“种类”目前只有“表头”和“枚举”，该配置表的结构也是固定
定义参考[https://github.com/davyxu/tabtoy/blob/master/v3/model/type.go]

[=====以下都是游戏中的数值配置表=====]
kv.xlsx是全局常量表
item.xlsx是所有道具表

## 执行`gen.sh`: `bash gen.sh`

生成data_gen.json和data_gen.go文件
