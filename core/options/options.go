package options

type Options struct {
	Urls       []string
	Thread     int
	Proxy      string
	Level      int // 指纹库等级判定
	Stdin      bool
	Timeout    int
	Screenshot bool
}
