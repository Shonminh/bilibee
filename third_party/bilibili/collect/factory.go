package collect

import bg "github.com/iyear/biligo"

// NewBilibiliClient 哔哩哔哩客户端工厂

const defaultUserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36"

func NewBilibiliClient() BilibiliClient {
	commClient := bg.NewCommClient(&bg.CommSetting{Client: nil, DebugMode: true, UserAgent: defaultUserAgent})
	return &BilibiliClientImpl{commonCli: commClient}
}
