package response

import "github.com/gin-gonic/gin"

type Response struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Total   int         `json:"total,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

type ResponseError struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

func ResponseJSON(ctx *gin.Context, res Response) Response {
	ctx.JSON(res.Status, res)
	return res
}

func ResponseWithError(ctx *gin.Context, res Response) Response {
	ctx.JSON(res.Status, res)
	return res
}

func ResponseWithPagination(ctx *gin.Context, res Response) Response {
	ctx.JSON(res.Status, res)
	return res
}
