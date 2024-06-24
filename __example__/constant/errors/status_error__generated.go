// Code generated by tools. DO NOT EDIT!
package errors

import (
    "fmt"
	"github.com/shrewx/ginx/pkg/statuserror"
	"strconv"
	"strings"
)

func (v StatusError) StatusErr(args ...interface{}) statuserror.CommonError {
	return &statuserror.StatusErr{
		Key:       v.Key(),
		ErrorCode:      v.Code(),
		Message:   fmt.Sprintf(v.ZhMessage(), args...),
		ZHMessage: fmt.Sprintf(v.ZhMessage(), args...),
		ENMessage: fmt.Sprintf(v.EnMessage(), args...),
	}
}

func (v StatusError) I18n(language string) statuserror.CommonError {
	e := &statuserror.StatusErr{
		Key:       v.Key(),
		ErrorCode: v.Code(),
		ZHMessage: v.ZhMessage(),
		ENMessage: v.EnMessage(),
	}
	language = strings.ToLower(language)
	switch language {
	case "zh":
		e.Message = v.ZhMessage()
	case "en":
		e.Message = v.EnMessage()
	}

	return e
}


func (v StatusError) StatusCode() int {
	strCode := fmt.Sprintf("%d", v.Code())
	if len(strCode) < 3 {
		return 400
	}
	statusCode, _ := strconv.Atoi(strCode[:3])
	return statusCode
}

func (v StatusError) Key() string {
	switch v { 
	case InternalServerError:
		return "InternalServerError"
	case BadRequest:
		return "BadRequest"
	case UploadFileIsNotExist:
		return "UploadFileIsNotExist"
	case Unauthorized:
		return "Unauthorized"
	case Forbidden:
		return "Forbidden"
	case NotFound:
		return "NotFound"
	case Conflict:
		return "Conflict"
	}
	return "UNKNOWN"
}

func (v StatusError) Code() int {
	return int(v)
}

func (v StatusError) Error() string {
	return fmt.Sprintf("[%s][%d] zh:(%s), en:(%s)", v.Key(), v.StatusCode(), v.ZhMessage(), v.EnMessage())
}

func (v StatusError) ZhMessage() string {
	switch v { 
	case InternalServerError:
		return "内部处理错误"
	case BadRequest:
		return "请求参数错误"
	case UploadFileIsNotExist:
		return "上传文件为找到"
	case Unauthorized:
		return "未授权，请先授权"
	case Forbidden:
		return "禁止操作"
	case NotFound:
		return "资源未找到"
	case Conflict:
		return "资源冲突"
	}
	return ""
}

func (v StatusError) EnMessage() string {
	switch v { 
	case InternalServerError:
		return ""
	case BadRequest:
		return ""
	case UploadFileIsNotExist:
		return ""
	case Unauthorized:
		return ""
	case Forbidden:
		return ""
	case NotFound:
		return ""
	case Conflict:
		return ""
	}
	return ""
}
