package constants

type LoginStatus int

const (
	LOGIN_SUCCESS         LoginStatus = 0
	LOGIN_PASSWORD_ERROR  LoginStatus = -1
	LOGIN_ACCOUNT_EXISTED LoginStatus = -2
	LOGIN_EMPTY_INFO      LoginStatus = -3
)
