package errno

// https://open.weibo.com/wiki/Error_code
var (
	OK                   = &ErrNo{Code: 0, Message: "OK"}
	InternaleServerError = &ErrNo{Code: 10001, Message: "Internal server error"}
	BindError            = &ErrNo{Code: 10002, Message: "Error occurred while binding the request body to the struct"}
	UserNotFoundError    = &ErrNo{Code: 20102, Message: "The user not found "}
)
