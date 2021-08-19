package form

type UserInfo struct {
	UserName string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
