package response

type AccessToken interface {
	accessToken()string
	expiresIn()int
	errCode()int
	errMsg()string
}