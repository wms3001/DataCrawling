# DataCrawling
------

简单的数据抓取

------

1. 安装
```go
go get github.com/wms3001/datacrawling
```
2. 抓取百度热搜信息
```go
var bn BaiDuHotNews
res := bn.GetBaiDuHostNews()
log.Println(res)
```