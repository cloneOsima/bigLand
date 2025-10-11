// handler package 전역에서 사용하는 각종 error check 및 logging fucntion을 모아둔 파일
package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	errdefs "github.com/cloneOsima/bigLand/backend/internal/errors"
	"github.com/cloneOsima/bigLand/backend/internal/utils"
	"github.com/gin-gonic/gin"
)

// request id 관련 로직을 통합해서 처리해주는 함수
func reqIdChecker(c *gin.Context) string {
	reqId := c.GetHeader("X-Request-ID")
	if reqId == "" {
		reqId = fmt.Sprint(utils.GenerateRequestId())
	}

	return reqId
}

// 공통적으로 발생할 수 있는 http response 관련 에러를 처리하는 함수
func commonErrorChecker(c *gin.Context, ctx context.Context, err error, reqId string) {
	if err != nil {

		// 에러가 잘못된 값으로 발생한 에러일 경우 400과 함께 내용 반환
		if apperr, ok := err.(*errdefs.ValueErr); ok {
			log.Printf("[Request failed] Request Id: %s, Error: %v, ErrorInfo: %v", reqId, err, apperr.ErrorInfo)
			c.JSON(http.StatusBadRequest, gin.H{
				"error":     "Request timeout",
				"requestID": reqId,
			})
			return
		}

		log.Printf("[Request failed] Request Id: %s, Error: %v", reqId, err)
		if errors.Is(err, context.DeadlineExceeded) || ctx.Err() == context.DeadlineExceeded {
			c.JSON(http.StatusRequestTimeout, gin.H{
				"error":     "Request timeout",
				"requestID": reqId,
			})
			return
		}
		if errors.Is(err, context.Canceled) || ctx.Err() == context.Canceled {
			c.JSON(http.StatusRequestTimeout, gin.H{
				"error":     "Request cancelled",
				"requestID": reqId,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":     err.Error(),
			"requestID": reqId,
		})
		return
	}
}
