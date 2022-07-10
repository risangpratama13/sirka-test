package request

type DisplayUserRequest struct {
	Userid string `form:"userid" binding:"required"`
}
