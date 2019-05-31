package object

type ErrTypeCode int

const (
	ErrTypeCodeNoError            ErrTypeCode = 0
	ErrTypeCodeWrongOprType       ErrTypeCode = 10001
	ErrTypeCodeWrongPassWord      ErrTypeCode = 10002
	ErrTypeCodeVerificationFailed ErrTypeCode = 10003
)

type ErrTypeMsg string

const (
	ErrTypeMsgNoError            ErrTypeMsg = ""
	ErrTypeMsgWrongOprType       ErrTypeMsg = ""
	ErrTypeMsgWrongPassWord      ErrTypeMsg = ""
	ErrTypeMsgVerificationFailed ErrTypeMsg = ""
)
