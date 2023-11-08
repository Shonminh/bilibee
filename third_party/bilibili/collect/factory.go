package collect

import (
	"github.com/Shonminh/bilibee/pkg/logger"
	bg "github.com/iyear/biligo"
	"os"
)

// NewBiliBiliClient 哔哩哔哩客户端工厂

const defaultUserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36"

func NewBiliBiliClient() BiliBiliClient {
	commClient := bg.NewCommClient(&bg.CommSetting{Client: nil, DebugMode: true, UserAgent: defaultUserAgent})
	accountClient, err := bg.NewBiliClient(&bg.BiliSetting{
		Auth: &bg.CookieAuth{
			DedeUserID:      os.Getenv("DedeUserID"),
			SESSDATA:        os.Getenv("SESSDATA"),
			BiliJCT:         os.Getenv("BiliJCT"),
			DedeUserIDCkMd5: os.Getenv("DedeUserIDCkMd5"),
		},
		Client:    nil,
		DebugMode: true,
		UserAgent: defaultUserAgent,
	})
	if err != nil {
		logger.LogPanicf("NewBiliBiliClient failed: %v", err)
	}
	return &BilibiliClientImpl{commonCli: commClient, accountCli: accountClient}
}
