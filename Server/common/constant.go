package common

const (
	StatusSuccess       = 0
	UserNameEmptyError  = 1
	PasswordEmptyError  = 2
	CallDBError         = 3
	UserExistedError    = 4
	SystemPanicErr      = 5
	NotFoundUserError   = 6
	PasswordWrongError  = 7
	TokenIsEmptyError   = 8
	CallRedisError      = 9
	ReqDataIsEmptyError = 10
	AskQueueError       = 11
	MatchOffCode        = 12
	RoomIDIsEmptyError  = 13
	RoomWrongError      = 14
)

const (
	UserNameEmptyMsg  = "username is empty"
	PasswordEmptyMsg  = "password is empty"
	UserExistedMsg    = "The user already exists"
	SystemPanicMsg    = "system panic error"
	NotFoundUserMsg   = "username not found"
	PasswordWrongMsg  = "password is wrong"
	TokenIsEmptyMsg   = "token is empty"
	ReqDataIsEmptyMsg = "request Data id is empty"
	MatchOffMsg       = "match off succeed"
	RoomIDIsEmptyMsg  = "roomID is empty"
	RoomIDIsWrongMsg  = "not found Player in the Room, Please check the roomID"
)

var CardItoaMap = map[int]string{
	0:  "A",
	1:  "2",
	2:  "3",
	3:  "4",
	4:  "5",
	5:  "6",
	6:  "7",
	7:  "8",
	8:  "9",
	9:  "10",
	10: "J",
	11: "Q",
	12: "K",
}

var CardAtoiMap = map[string]int{
	"A":  0,
	"2":  1,
	"3":  2,
	"4":  3,
	"5":  4,
	"6":  5,
	"7":  6,
	"8":  7,
	"9":  8,
	"10": 9,
	"J":  10,
	"Q":  11,
	"K":  12,
}
