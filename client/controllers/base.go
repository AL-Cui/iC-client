package controllers

import (
	"fmt"
	"net/http"

	"iCenter-client/def"
	"iCenter-client/utils/errors"
	
	"github.com/astaxie/beego"
)
type BaseController struct {
	beego.Controller
}

// APIAbort is a function to prevent API request by panic.
func (b *BaseController) APIAbort(apiCode string, err error) {
	b.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
	msg := fmt.Sprintf("%v", err)
	errCode := apiCode + errors.ErrorCode(err)
	data := def.Error{Code: errCode, Msg: msg, Level: "Error"}
	b.Data["json"] = data
	b.ServeJSON()
	beego.Debug("api request aborted due to:", data)
	panic(err)
}

// unauthorizedAbort returns unauthorized error with HTTP code 401.
func (b *BaseController) unauthorizedAbort(err error) {
	b.Ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
	err401 := def.Unauthorized
	err401.Msg = err.Error()
	b.Data["json"] = err401
	b.ServeJSON()
	beego.Error("unauthorized request:", err401)
	panic(err401)
}

// forbiddenAbort returns forbidden error with HTTP code 403.
func (b *BaseController) forbiddenAbort() {
	b.Ctx.ResponseWriter.WriteHeader(http.StatusForbidden)
	err403 := def.Forbidden
	b.Data["json"] = err403
	b.ServeJSON()
	beego.Error("forbidden request:", err403)
	panic(err403)
}

// badRequestAbort returns request error with HTTP code 400, it usually means
// client send wrong parameters. The request could not be understood by the
// server due to malformed syntax. The client SHOULD NOT repeat the request
// without modifications.
func (b *BaseController) badRequestAbort(apiCode string, err error) {
	b.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	err400 := def.BadRequest
	err400.Code = apiCode + err400.Code
	err400.Msg = fmt.Sprintf("%s: %v", err400.Msg, err.Error())
	b.Data["json"] = err400
	b.ServeJSON()
	beego.Error("bad request:", err400)
	panic(err400)
}

// licenseInvalid returns forbidden error with HTTP code 403. Frontend should
// redirect to an independent page to show the infomation about license.
func (b *BaseController) licenseInvalid() {
	b.Ctx.ResponseWriter.WriteHeader(http.StatusForbidden)
	resp := def.InvalidLic
	b.Data["json"] = resp
	b.ServeJSON()
	beego.Error("invalid license:", resp)
	panic(resp)
}