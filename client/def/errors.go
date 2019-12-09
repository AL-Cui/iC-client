package def

import (
	"strconv"
)

// Error represents a general Error struct contains errcode and errmsg
type Error struct {
	Code  string
	Msg   string
	Level string
}

const (
	ErrPodModule 				= 1000
	ErrOK                       = 0    // EN: done   zh： 操作成功
	ErrGeneralUnauthorized    	= 9009 // EN： unauthorized   ZH： 未授权的操作
	ErrGeneralForbidden       	= 9010 // EN： forbidden request   ZH： 您无权查看
	ErrGeneralBadRequest      	= 9011 // EN： bad request   ZH： 错误的请求参数
	ErrInvalidLicense          	= 9998 // EN： invalid license  ZH： 无效的许可证
	ErrGeneralUnknown          	= 9999 // EN： unknown error  ZH： 未知错误
	 
)

const (
	ListPodsFromNamespace   	= "1001"
)

// Define standard error response with errcode and errmsg
var (
	Success      = Error{Code: strconv.Itoa(ErrOK), Msg: "success", Level: "Info"}
	Unknown      = Error{Code: strconv.Itoa(ErrGeneralUnknown), Msg: "ERR_UNKNOWN", Level: "Error"}
	Unauthorized = Error{Code: strconv.Itoa(ErrGeneralUnauthorized), Msg: "Unauthorized request", Level: "Error"}
	Forbidden    = Error{Code: strconv.Itoa(ErrGeneralForbidden), Msg: "You have no permission to access", Level: "Error"}
	BadRequest   = Error{Code: strconv.Itoa(ErrGeneralBadRequest), Msg: "You have send an error request", Level: "Error"}
	InvalidLic   = Error{Code: strconv.Itoa(ErrInvalidLicense), Msg: "License is invalid or expired", Level: "Error"}
)