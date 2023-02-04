package collect

import bg "github.com/iyear/biligo"

// NewBilibiliClient 哔哩哔哩客户端工厂
func NewBilibiliClient() BilibiliClient {
	commClient := bg.NewCommClient(&bg.CommSetting{Client: nil, DebugMode: true})
	return &BilibiliClientImpl{commonCli: commClient}
}
