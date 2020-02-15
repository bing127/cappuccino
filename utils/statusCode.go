package utils

const (
	UserNotExist           		= 200001
	UserAlreadyExist       		= 200002
)

var statusText = map[int]string{
	UserNotExist:   			"用户不存在",
	UserAlreadyExist:   		"用户已存在",
}

func StatusText(code int) string {
	return statusText[code]
}