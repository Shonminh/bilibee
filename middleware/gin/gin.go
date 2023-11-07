package gin

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	l "github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/Shonminh/bilibee/pkg/db"
	"github.com/Shonminh/bilibee/pkg/logger"
)

func UseMysql(d *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.Request
		ctx := req.Context()
		newCtx := db.BindDbContext(ctx, d)
		c.Request = req.WithContext(newCtx)
	}
}

var format = func(i interface{}) string { return fmt.Sprintf("%s;", i) }

const timeFormat = "2006-01-02T15:04:05.999999999"
const maxReqDataLength = 4096

func UseGinLogger() gin.HandlerFunc {
	output, err := logger.GetOutput()
	if err != nil {
		logger.LogPanic(err)
	}
	return l.SetLogger(l.WithLogger(func(ctx *gin.Context, z zerolog.Logger) zerolog.Logger {
		rawData, _ := ctx.GetRawData()
		data := rawData
		if len(data) > maxReqDataLength {
			data = rawData[:maxReqDataLength]
		}
		logger.LogInfof("%s", string(data))
		// 赋值，保证下次可以读取
		readCloser := io.NopCloser(bytes.NewReader(rawData))
		ctx.Request.Body = readCloser
		consoleWriter := zerolog.ConsoleWriter{
			Out:        output,
			NoColor:    true,
			TimeFormat: timeFormat,
			FormatLevel: func(i interface{}) string {
				return strings.ToUpper(fmt.Sprintf("|%s|", i))
			},
			FormatMessage:       func(i interface{}) string { return fmt.Sprintf("message=%s; req_body=%q;", i, string(data)) },
			FormatFieldValue:    format,
			FormatErrFieldValue: format,
		}
		return z.Output(consoleWriter).With().Logger()
	}))
}
