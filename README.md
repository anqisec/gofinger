# gofinger

## 工具介绍

gofinger 是一款由 golang 开发的批量 url 处理工具, 可一键对 url 进行指纹识别、存活验证、站点去重、站点截图。

当拿到大量 url 时可快速对其资产有一定认识，寻找薄弱资产。



```shell
Usage:
  gofinger [flags]

Flags:
  -f, --file string     -f targets.txt
  -h, --help            help for gofinger
  -l, --level int       -l 1-2 (default 1)
  -o, --output string   -o results.csv
  -p, --proxy string    -p http://127.0.0.1:8080
      --stdin           --stdin true
  -t, --thread int      -t 25 (default 50)
  -u, --url string      -u https://www.baidu.com
  -v, --version         version for gofinger
  
```
level 指的是用于匹配的指纹数量，1 级误报几乎没有，但由于 chunsou 中有些指纹都是单条规则匹配的，并的数量不多，可能会有误报。这里的数量是指 CMS 的数量，会有重复，可以自己修改指纹去掉。

| level | 数量  | 来源                              |
| ----- | ----- |---------------------------------|
| 1     | 3499  | Goby软件指纹 + chunsou的icon_hash 部分 |
| 2     | 10379  | Goby软件指纹 + chunsou所有规则          |

--stdin 就是接收外部的 url 数据，比如 cat urls.txt | gofinger.exe --stdin true，可以配合向其他工具使用。