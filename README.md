# gameserver

基于开源游戏框架leaf的初步尝试

## 编译并执行
cd app
go build -o gamesever


## 配置表生成对应的数据结构和json文件
直接执行data.sh：`bash data.sh`
生成的数据表json到app/data文件夹
生成的结构体到gamedata文件夹