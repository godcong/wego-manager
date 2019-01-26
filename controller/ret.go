package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// ContentTypeJSON ...
const ContentTypeJSON = "application/json"

// ContentTypeXML ...
const ContentTypeXML = "application/xml"

// ContentType ...
const ContentType = "Accept"

// RetDetail ...
type RetDetail struct {
	Code    int    `json:"code" example:"-1"`
	Message string `json:"message" example:"status bad request"`
}

// Ret ...
type Ret struct {
	RetDetail `json:",inline"`
	Detail    interface{}
}

func resultJSON(ctx *gin.Context, code int, msg string, detail interface{}) {
	if detail != nil {
		ctx.JSON(http.StatusOK, &Ret{
			RetDetail: RetDetail{
				Code:    code,
				Message: msg,
			},
			Detail: detail,
		})
		return
	}
	ctx.JSON(http.StatusOK, &RetDetail{
		Code:    code,
		Message: msg,
	})
}

func resultXML(ctx *gin.Context, code int, msg string, detail interface{}) {
	if detail != nil {
		ctx.XML(http.StatusOK, &Ret{
			RetDetail: RetDetail{
				Code:    code,
				Message: msg,
			},
			Detail: detail,
		})
		return
	}
	ctx.XML(http.StatusOK, &RetDetail{
		Code:    code,
		Message: msg,
	})
}

// result ...
func result(ctx *gin.Context, code int, msg string, detail interface{}) {
	accept := ctx.Request.Header.Get(ContentType)
	log.Println(accept)
	switch accept {
	case ContentTypeJSON:
		resultJSON(ctx, code, msg, detail)
	case ContentTypeXML:
		resultXML(ctx, code, msg, detail)
	}

}

// fail ...
func fail(ctx *gin.Context, msg string) {
	result(ctx, -1, msg, nil)
}

// success ...
func success(ctx *gin.Context, msg string) {
	result(ctx, 0, msg, nil)
}

// detail ...
func detail(ctx *gin.Context, d interface{}) {
	result(ctx, 0, "success", d)
}
