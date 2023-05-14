package common

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type HTTPRsp struct {
	Status  string      `json:"status,omitempty"`
	Success bool        `json:"success,omitempty"`
	Error   *Error      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type Error struct {
	No     string `json:"no,omitempty"`
	Msg    string `json:"msg,omitempty"`
	OriMsg string `json:"oriMsg,omitempty"`
}

func (e Error) Error() string {
	return fmt.Errorf("error: No.%s, %s. the oriMsg is: %s", e.No, e.Msg, e.OriMsg).Error()
}

type HandlerFunc struct {
	Method string
	Path   string
	app.HandlerFunc
}

type PageData struct {
	Current  int         `json:"current,omitempty"`
	PageSize int         `json:"pageSize,omitempty"`
	Total    int         `json:"total"`
	List     interface{} `json:"list,omitempty"`
}

func NewError(err Error, oriMsg error) *Error {
	if oriMsg != nil {
		err.OriMsg = oriMsg.Error()
	}
	return &err
}

func WriteFailHTTP(r *app.RequestContext, rsp *HTTPRsp, err error) {
	rsp.Success = false

	switch e := err.(type) {
	case *Error:
		rsp.Error = e
	default:
		rsp.Error = &Error{Msg: e.Error()}
	}

	r.JSON(consts.StatusOK, rsp)
}

func WriteSuccessHTTP(r *app.RequestContext, rsp *HTTPRsp) {
	rsp.Success = true
	r.JSON(consts.StatusOK, rsp)
}
