# gofinger

## 工具介绍

gofinger 是一款由 golang 开发的 web 指纹识别工具，指纹识别速度快，指纹库较全。默认配置识别 1000 个 url 大概 3 分钟左右。默认指纹 3400 多条。

且本工具会对站点 进行自动去重，根据其 IP 和响应体 hash 进行去重，去重重复站点。另外可以进行站点截图，需安装有 `chrome` 浏览器，截图也不慢，3 min 100 个左右，默认只开一个浏览器 15 个标签页，防止内存占用过大。

只进行指纹识别会自动保存为 csv 格式，指纹识别 + 截图为 html：

![image-20231218215319530](https://gallery-1310215391.cos.ap-beijing.myqcloud.com/img/image-20231218215319530.png)

![image-20231218215356635](https://gallery-1310215391.cos.ap-beijing.myqcloud.com/img/image-20231218215356635.png)

## 工具使用

```shell
gofinger -u https://www.baidu.com
gofinger -f targets.txt 	
gofinger -f targets.txt -s	// 进行站点截图
gofinger -u https://www.baidu.com -l 2 // 使用 10379 的指纹库(多但是可能不准确)
gofinger -u https://www.baidu.com -p http://127.0.0.1:8080 // 代理
cat urls.txt | gofinger --stdin true -s // 进行指纹识别+截图(配合其他工具使用)
gofinger -h 
```

## 其他介绍

> PS：欢迎大家使用并对有问题的地方提提建议呀，收到后会尽快优化 ~

支持的指纹识别模式：

- title、header、body
- cert
- icon_hash

指纹格式：

```json
{
    "cms": "thinkphp",
    "rule": "((header=\"thinkphp\") && header!=\"couchdb\" && header!=\"st: upnp:rootdevice\") || body=\"href=\\\"http://www.thinkphp.cn\\\">thinkphp</a><sup>\" || ((header=\"thinkphp\") && header!=\"couchdb\" && header!=\"st: upnp:rootdevice\") || icon_hash=\"1165838194\""
},
```


