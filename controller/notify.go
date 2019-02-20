package controller

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego"
	"github.com/godcong/wego/cipher"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"net/http"
)

// Notify ...
type Notify struct {
	//property  *wego.Property
	payment   *wego.Payment
	RefundKey string
}

// NewNotify ...
func NewNotify(payment *wego.Payment) *Notify {
	return &Notify{
		payment:   payment,
		RefundKey: "",
	}
}

// HandleRefunded ...
func (n *Notify) HandleRefunded(f ServeNotify) Notifier {
	return &paymentRefunded{
		cipher: cipher.New(cipher.AES256ECB, &cipher.Option{
			Key: n.RefundKey,
		}),
		ServeNotify: f,
	}
}

// HandleRefundedNotify ...
func (n *Notify) HandleRefundedNotify(f ServeNotify) ServeHTTPFunc {
	return n.HandleRefunded(f).ServeHTTP
}

// HandleScannedNotify ...
func (n *Notify) HandleScannedNotify(f ServeNotify) Notifier {
	return &paymentScanned{
		Notify:      n,
		ServeNotify: f,
	}
}

// HandleScanned ...
func (n *Notify) HandleScanned(f ServeNotify) ServeHTTPFunc {
	return n.HandleScannedNotify(f).ServeHTTP
}

// HandlePaidNotify ...
func (n *Notify) HandlePaidNotify(f ServeNotify) Notifier {
	return &paymentPaid{
		Notify:      n,
		ServeNotify: f,
	}
}

// Notifier ...
type Notifier interface {
	ServeHTTP(ctx *gin.Context)
}

// ServeNotify ...
type ServeNotify func(p util.Map) (util.Map, error)

// ServeHTTPFunc ...
type ServeHTTPFunc func(ctx *gin.Context)

/*Notifier 监听 */
type paymentRefunded struct {
	cipher cipher.Cipher
	ServeNotify
}

// ServeHTTP ...
func (obj *paymentRefunded) ServeHTTP(ctx *gin.Context) {
	var err error
	rlt := SUCCESS()
	defer func() {
		err = XMLResponse(ctx.Writer, rlt.ToXML())
		log.Error(err)
	}()
	maps, err := core.RequestToMap(ctx.Request)
	//wrong request will do nothing
	if err != nil {
		log.Error(err)
		rlt = FAIL(err.Error())
	} else {
		reqInfo := maps.GetString("req_info")
		maps.Set("reqInfo", obj.DecodeReqInfo(reqInfo))
		if obj.ServeNotify == nil {
			log.Error(xerrors.New("null notify callback"))
			return
		}
		_, err = obj.ServeNotify(maps)
		if err != nil {
			rlt = FAIL(err.Error())
		}
	}

}

// DecodeReqInfo ...
func (obj *paymentRefunded) DecodeReqInfo(info string) util.Map {
	maps := util.Map{}
	dec, _ := obj.cipher.Decrypt(info)
	e := xml.Unmarshal(dec, &maps)
	if e != nil {
		log.Error(e)
	}
	return maps
}

/*Notifier 监听 */
type paymentScanned struct {
	*Notify
	ServeNotify
}

// ServeHTTP ...
func (obj *paymentScanned) ServeHTTP(ctx *gin.Context) {
	var e error
	rlt := SUCCESS()
	defer func() {
		e = XMLResponse(ctx.Writer, rlt.ToXML())
		log.Error(e)
	}()
	var p util.Map
	maps, e := core.RequestToMap(ctx.Request)

	if e != nil {
		log.Error(xerrors.Errorf("paymentScanned RequestToMap:%w", e))
		rlt = FailDes(e.Error())
	} else {
		if util.ValidateSign(maps, obj.payment.GetKey()) {
			if obj.ServeNotify == nil {
				log.Error(xerrors.New("null notify callback"))
				return
			}
			p, e = obj.ServeNotify(maps)
			if e != nil {
				rlt = FailDes(e.Error())
			}
			if !p.Has("prepay_id") {
				log.Error("null prepay_id")
				rlt = FailDes("null prepay_id")
			} else {
				//公众账号ID	appid	String(32)	是	wx8888888888888888	微信分配的公众账号ID
				//商户号	mch_id	String(32)	是	1900000109	微信支付分配的商户号
				//随机字符串	nonce_str	String(32)	是	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	微信返回的随机字符串
				//预支付ID	prepay_id	String(64)	是	wx201410272009395522657a690389285100	调用统一下单接口生成的预支付ID
				//业务结果	result_code	String(16)	是	SUCCESS	SUCCESS/FAIL
				//错误描述	err_code_des	String(128)	否		当result_code为FAIL时，商户展示给用户的错误提
				//签名	sign	String(32)	是	C380BEC2BFD727A4B6845133519F3AD6	返回数据签名，签名生成算法
				rlt.Set("appid", obj.payment.AppID)
				rlt.Set("mch_id", obj.payment.MchID)
				rlt.Set("nonce_str", util.GenerateNonceStr())
				rlt.Set("prepay_id", p.Get("prepay_id"))
				rlt.Set("sign", util.GenSign(maps, obj.payment.GetKey()))

			}

		}
	}
}

/*Notifier 监听 */
type paymentPaid struct {
	*Notify
	ServeNotify
}

// ServerHttp ...
func (n *paymentPaid) ServeHTTP(ctx *gin.Context) {
	var e error
	rlt := SUCCESS()
	defer func() {
		e = XMLResponse(ctx.Writer, rlt.ToXML())
		log.Error(e)
	}()
	maps, e := wego.BuildRequester(ctx.Request).Result()

	if e != nil {
		log.Error(e)
		rlt = FAIL(e.Error())
	} else {
		if util.ValidateSign(maps, n.payment.GetKey()) {
			if n.ServeNotify == nil {
				log.Error(xerrors.New("null notify callback "))
				return
			}
			_, e = n.ServeNotify(maps)
			if e != nil {
				log.Error(xerrors.Errorf(" paymentPaid ServeNotify error:%w", e))
				rlt = FAIL(e.Error())
			}
		}
	}

}

// XMLResponse ...
func XMLResponse(w http.ResponseWriter, data []byte) error {
	w.WriteHeader(http.StatusOK)
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{"application/xml; charset=utf-8"}
	}
	_, err := w.Write(data)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// JSONResponse ...
func JSONResponse(w http.ResponseWriter, data []byte) error {
	w.WriteHeader(http.StatusOK)
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{"application/json; charset=utf-8"}
	}
	_, err := w.Write(data)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// SUCCESS ...
func SUCCESS() util.Map {
	return util.Map{
		"return_code": "SUCCESS",
		"return_msg":  "OK",
	}
}

// FAIL ...
func FAIL(msg string) util.Map {
	return util.Map{
		"return_code": "FAIL",
		"return_msg":  msg,
	}
}

// FailDes ...
func FailDes(msg string) util.Map {
	return util.Map{
		"return_code":  "FAIL",
		"err_code_des": msg,
	}
}
