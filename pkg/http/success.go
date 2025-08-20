package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func SimpleSuccessResponse(c *gin.Context) {
	SuccessResponse(c, nil)
}

func SuccessResponse(c *gin.Context, data interface{}) {
	// set trace id to response header
	// key is X-Trace-ID
	c.JSON(http.StatusOK, BaseResponse{
		Success:   true,
		ErrorCode: nil,
		Data:      data,
	})
}

func SuccessExportExcel(c *gin.Context, data []byte, name string) {
	unix := time.Now().Unix()
	fileName := fmt.Sprintf("%s_%d.xlsx", name, unix)
	// set trace id to response header
	// key is X-Trace-ID
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", fmt.Sprintf("attachment;filename=%s", fileName))
	c.Header("File-Name", fileName)
	c.Header("Content-Length", strconv.Itoa(len(data)))

	// Trả file Excel về client
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", data)
}
