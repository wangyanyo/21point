package common

const (
	StatusSuccess      = 0
	UserNameEmptyError = 1
	PasswordEmptyError = 2
	CallDBError        = 3
	UserExistedError   = 4
	ProcessErr         = 5
)

const (
	UserNameEmptyMsg = "username is empty"
	PasswordEmptyMsg = "password is empty"
	UserExistedMsg   = "The user already exists"
)
