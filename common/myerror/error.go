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

type GetScoreError struct{}

func (*GetScoreError) Error() string {
	return "GetScoreError"
}

type RankListError struct{}

func (*RankListError) Error() string {
	return "RankListError"
}

type UserCountError struct{}

func (*UserCountError) Error() string {
	return "UserCountError"
}
