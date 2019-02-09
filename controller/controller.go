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

func resultJSON(ctx *gin.Context, detail interface{}) {
	if detail != nil {
		ctx.JSON(http.StatusOK, detail)
		return
	}
	ctx.JSON(http.StatusOK, &CodeMessage{
		Code:    0,
		Message: "success",
	})
}

func resultXML(ctx *gin.Context, detail interface{}) {
	if detail != nil {
		ctx.XML(http.StatusOK, detail)
		return
	}
	ctx.XML(http.StatusOK, &CodeMessage{
		Code:    0,
		Message: "success",
	})
}

// Success ...
func Success(ctx *gin.Context, detail interface{}) {
	switch ctx.NegotiateFormat(ContentTypeJSON, ContentTypeXML) {

	case ContentTypeXML:
		resultXML(ctx, detail)
	case ContentTypeJSON:
		fallthrough
	default:
		resultJSON(ctx, detail)
	}
}

// fail ...
func fail(ctx *gin.Context, code int, msg string) {
	accept := ctx.GetHeader(ContentType)
	log.Println(accept)
	switch accept {
	case ContentTypeJSON:
		ctx.JSON(http.StatusBadRequest, &CodeMessage{
			Code:    code,
			Message: msg,
		})
	case ContentTypeXML:
		ctx.XML(http.StatusBadRequest, &CodeMessage{
			Code:    code,
			Message: msg,
		})
	}
}
