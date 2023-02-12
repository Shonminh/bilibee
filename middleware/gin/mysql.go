package gin

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/Shonminh/bilibee/pkg/db"
)

func UseMysql(d *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.Request
		ctx := req.Context()
		newCtx := db.SetDbContext(ctx, d)
		c.Request = req.WithContext(newCtx)
	}
}
