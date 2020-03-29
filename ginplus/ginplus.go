package ginplus

import (
	"context"
	"cappuccino/controller/schema"
	"cappuccino/errors"
	"cappuccino/utils"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	icontext "cappuccino/context"
)

// 定义上下文中的键
const (
	prefix = "gin-admin"
	// UserIDKey 存储上下文中的键(用户ID)
	UserIDKey = prefix + "/user-id"
	// TraceIDKey 存储上下文中的键(跟踪ID)
	TraceIDKey = prefix + "/trace-id"
	// ResBodyKey 存储上下文中的键(响应Body数据)
	ResBodyKey = prefix + "/res-body"
)

// NewContext 封装上下文入口
func NewContext(c *gin.Context) context.Context {
	parent := context.Background()

	if v := GetTraceID(c); v != "" {
		parent = icontext.NewTraceID(parent, v)
	}

	if v := GetUserID(c); v != "" {
		parent = icontext.NewUserID(parent, v)
	}

	return parent
}

// GetToken 获取用户令牌
func GetToken(c *gin.Context) string {
	var token string
	auth := c.GetHeader("Authorization")
	prefix := "Bearer "
	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}
	return token
}


// GetTraceID 获取追踪ID
func GetTraceID(c *gin.Context) string {
	return c.GetString(TraceIDKey)
}

// GetUserID 获取用户ID
func GetUserID(c *gin.Context) string {
	return c.GetString(UserIDKey)
}

// SetUserID 设定用户ID
func SetUserID(c *gin.Context, userID string) {
	c.Set(UserIDKey, userID)
}

// ParseJSON 解析请求JSON
func ParseJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		return errors.Wrap400Response(err, "解析请求参数发生错误")
	}
	return nil
}

// ParseQuery 解析Query参数
func ParseQuery(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		return errors.Wrap400Response(err, "解析请求参数发生错误")
	}
	return nil
}

// ParseForm 解析Form请求
func ParseForm(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindWith(obj, binding.Form); err != nil {
		return errors.Wrap400Response(err, "解析请求参数发生错误")
	}
	return nil
}


// ResList 响应列表数据
func ResList(c *gin.Context, v interface{}) {
	ResSuccess(c, schema.HTTPList{List: v})
}

// ResOK 响应OK
func ResOK(c *gin.Context) {
	ResSuccess(c, schema.HTTPStatus{Status: schema.OKStatusText.String()})
}

// ResSuccess 响应成功
func ResSuccess(c *gin.Context, v interface{}) {
	ResJSON(c, http.StatusOK, v)
}

// ResJSON 响应JSON数据
func ResJSON(c *gin.Context, status int, v interface{}) {
	buf, err := utils.JSONMarshal(v)
	if err != nil {
		panic(err)
	}
	c.Set(ResBodyKey, buf)
	c.Data(status, "application/json; charset=utf-8", buf)
	c.Abort()
}

// ResError 响应错误
func ResError(c *gin.Context, err error, status ...int) {
	var res *errors.ResponseError
	if err != nil {
		if e, ok := err.(*errors.ResponseError); ok {
			res = e
		} else {
			res = errors.UnWrapResponse(errors.Wrap500Response(err))
		}
	} else {
		res = errors.UnWrapResponse(errors.ErrInternalServer)
	}

	if len(status) > 0 {
		res.StatusCode = status[0]
	}

	eitem := schema.HTTPErrorItem{
		Code:    res.Code,
		Message: res.Message,
	}
	ResJSON(c, res.StatusCode, schema.HTTPError{Error: eitem})
}
