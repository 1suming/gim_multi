package middleware

import (
	"bytes"
	uuid "github.com/satori/go.uuid"

	"io/ioutil"
	"time"

	"gim/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.RequestURI
		//ip := c.ClientIP()
		mothod := c.Request.Method
		uid := uuid.NewV4().String()
		c.Set("trace_id", uid)
		reqBody, _ := c.GetRawData()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))

		logger.Logger.Info("http_request", zap.String("path", path),
			zap.String("method", mothod),
			zap.Any("query", string(reqBody)),
			zap.String("uid", uid),
		)

		//
		//logger.Logger.Info("key "
		//	zap.Any("key", "req"),
		//	zap.Any("ip", ip),
		//	zap.Any("path", path),
		//	zap.String("method", mothod),
		//	zap.Any("query", string(reqBody)),
		//)

		c.Next()

		respBpdy, _ := c.Get("respBody")
		elapse := time.Since(start)
		//logger.Logger.Info(uid,
		//	zap.Any("key", "resp"),
		//	zap.String("ip", ip),
		//	zap.String("path", path),
		//	zap.String("method", mothod),
		//	zap.Duration("elapse", elapse),
		//	zap.Any("body", respBpdy),
		//)

		logger.Logger.Info("http_response", zap.String("path", path),
			zap.String("method", mothod),
			zap.Any("body", respBpdy),
			zap.Duration("elapse", elapse),
			zap.String("uid", uid),
		)

	}
}
