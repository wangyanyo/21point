package myerror

type PasswordWrongError struct{}

func (*PasswordWrongError) Error() string {
	return "PasswordWrongError"
}

type NoUserNameError struct{}

func (*NoUserNameError) Error() string {
	return "NoUserNameError"
}

type RepeatUsernameError struct{}

func (*RepeatUsernameError) Error() string {
	return "RepeatUsernameError"
}
