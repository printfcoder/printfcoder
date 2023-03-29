package common

import (
	"fmt"
	"net/http"

	log "github.com/stack-labs/stack/logger"
	"github.com/stack-labs/stack/service/web"
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

func WriteFailHTTP(w http.ResponseWriter, rsp *HTTPRsp, err error) {
	rsp.Success = false
	rsp.Error = err.(*Error)
	_, e := web.HTTPJSON(w, rsp)
	if e != nil {
		log.Errorf("return json error, %s", err)
		return
	}
}

func WriteSuccessHTTP(w http.ResponseWriter, rsp *HTTPRsp) {
	rsp.Success = true
	_, err := web.HTTPJSON(w, rsp)
	if err != nil {
		log.Errorf("return json error, %s", err)
		return
	}
}
