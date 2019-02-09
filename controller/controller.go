package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/model"
	"log"
	"net/http"
	"strings"
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
	case ContentTypeJSON:
		resultJSON(ctx, detail)
	case ContentTypeXML:
		resultXML(ctx, detail)
	}

}

// ServerBack ...
func ServerBack(ver string) gin.HandlerFunc {
	//var back model.UserCallback
	//backs, e := back.Callbacks()
	//if e != nil {
	//	return nil
	//}

	return func(ctx *gin.Context) {
		paths := strings.Split(ctx.Request.URL.Path, "/")
		var back model.UserCallback
		back.Ver = paths[2]
		back.Sign = paths[3]
		back.URI = paths[4]
		back.BackType = paths[5]
		model.FindWhere(back)
		//p := model.UserProperty{}
		//
		//n := notify.NewNotify(wego.NewPayment())
		log.Println(ctx.HandlerName())
		log.Println(ctx.Request.URL.RawPath)
		log.Println(ctx.Request.URL.Path)
		log.Println(ctx.Request.URL.EscapedPath())
		for _, val := range backs {
			path := strings.Join([]string{"/api", val.Ver, val.Sign, val.URI}, "/")
			if path == ctx.Request.URL.Path {
				log.Println("founded", path)
			}
		}

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
