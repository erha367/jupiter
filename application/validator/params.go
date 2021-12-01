package validator

type UserInfo struct {
	UserName string `validator:"username" binding:"required"`
	Password string `validator:"password" binding:"required"`
}
