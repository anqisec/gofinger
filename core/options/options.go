package options

type Options struct {
	Urls   []string
	Output string
	Thread int
	Proxy  string
	Level  int // 指纹库等级判定
	Stdin  bool
}
