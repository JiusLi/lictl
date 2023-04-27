package response

import (
	"encoding/json"
	"fmt"
	"lictl/tools/errorx"
	"lictl/tools/process"
	"log"
	"net/http"
)

type body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var b body
	if err != nil {
		codeErr, ok := fromError(err)
		// ok == true 业务内部抛出的错误
		if ok {
			b.Code = codeErr.Code
			b.Msg = codeErr.Msg
		} else {
			// ok == false 服务端内部抛出错误
			errorJson(w, 500)
		}
		// 显示报错信息的详细信息
		if codeErr.Err != nil {
			log.Println(codeErr.Error())
		}
	} else {
		b.Code = http.StatusOK
		b.Msg = "执行成功"
	}
	b.Data = process.IF(resp == nil, []int{}, resp)
	okJson(w, b)

}

func BadRequestResponse(w http.ResponseWriter) {
	var b body
	b.Code = http.StatusBadRequest
	b.Msg = "错误请求"
	doWriteJson(w, http.StatusBadRequest, b)
}

func okJson(w http.ResponseWriter, v any) {
	doWriteJson(w, http.StatusOK, v)
}

func errorJson(w http.ResponseWriter, code int) {
	doWriteJson(w, http.StatusBadRequest, nil)
}

func doWriteJson(w http.ResponseWriter, code int, v any) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if v != nil {
		bs, err := json.Marshal(v)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return fmt.Errorf("marshal json failed, error: %w", err)
		}
		if n, err := w.Write(bs); err != nil {
			if err != http.ErrHandlerTimeout {
				return fmt.Errorf("write response failed, error: %w", err)
			}
		} else if n < len(bs) {
			return fmt.Errorf("actual bytes: %d, written bytes: %d", len(bs), n)
		}
	}

	return nil
}

func fromError(err error) (codeErr *errorx.ErrorInfo, ok bool) {
	if se, ok := err.(*errorx.ErrorInfo); ok {
		return se, true
	}
	return nil, false
}
