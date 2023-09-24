# gofinger

gofinger 是一款 web 指纹识别工具。

```shell
Usage:
  gofinger [flags]

Flags:
  -f, --file string     -f targets.txt
  -h, --help            help for gofinger
  -l, --level int       -l 1-3 (default 1)
  -o, --output string   -o results.csv
  -p, --proxy string    -p http://127.0.0.1:8080
      --stdin           --stdin true
  -t, --thread int      -t 25 (default 50)
  -u, --url string      -u https://www.baidu.com
  -v, --version         version for gofinger
  
```

`level` 指的是用于匹配的指纹数量，1 级误报几乎没有，但由于 chunsou 中有些指纹都是单条规则匹配的，并的数量不多，可能会有误报。这里的数量是指 CMS 的数量，会有重复，可以自己修改指纹去掉。

| level | 数量  | 来源                                |
| ----- | ----- | ----------------------------------- |
| 1     | 7288  | Goby + chunsou 的 icon_hash 部分    |
| 2     | 8541  | Goby + chunsou CMS 有多条规则的部分 |
| 3     | 12830 | Goby + chunsou 合并后的             |

`--stdin` 就是接收外部的 url 数据，比如 `cat urls.txt | gofinger.exe --stdin true`，可以配合向其他工具使用。

感谢 Goby 和 icon_hash 的指纹 ！
