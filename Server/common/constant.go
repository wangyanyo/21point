package common

const (
	StatusSuccess       = 0
	UserNameEmptyError  = 1
	PasswordEmptyError  = 2
	CallDBError         = 3
	UserExistedError    = 4
	ProcessErr          = 5
	NotFoundUserError   = 6
	PasswordWrongError  = 7
	TokenIsEmptyError   = 8
	CallRedisError      = 9
	ReqDataIsEmptyError = 10
)

const (
	UserNameEmptyMsg  = "username is empty"
	PasswordEmptyMsg  = "password is empty"
	UserExistedMsg    = "The user already exists"
	NotFoundUserMsg   = "username not found"
	PasswordWrongMsg  = "password is wrong"
	TokenIsEmptyMsg   = "token is empty"
	ReqDataIsEmptyMsg = "request Data id is empty"
)
