package es

import (
	"github.com/Shonminh/bilibee/pkg/logger"
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
	"os"
	"strings"
)

func NewEsClient() *elasticsearch8.Client {
	output, err := logger.GetOutput()
	if err != nil {
		logger.LogPanic(err.Error())
	}
	cfg := elasticsearch8.Config{
		Addresses:         strings.Split(os.Getenv("ES_ADDRESS_LIST"), ","),
		Username:          "",
		Password:          "",
		MaxRetries:        3,
		EnableDebugLogger: true,
		Logger: &elastictransport.TextLogger{
			Output:             output,
			EnableRequestBody:  true,
			EnableResponseBody: true,
		},
	}
	es, err := elasticsearch8.NewClient(cfg)
	if err != nil {
		logger.LogPanic(err)
	}
	return es
}
