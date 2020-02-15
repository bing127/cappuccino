package dto

type CreateUserRequest struct {
	LoginName    string `json:"loginName" binding:"required,excludes= "`
	NickName     string `json:"nickName" binding:"required,excludes= "`
}

type UpdateUserRequest struct {
	ID string `json:"id",binding:"required"`
	NickName     string `json:"nickName" binding:"required,excludes= "`
	Phone     string `json:"phone" binding:"required,excludes= "`
}

type CreateUserResponse struct {
	ID string `json:"id" binding:"required"`
}

type UserInfo struct {
	LoginName    string `json:"loginName" binding:"required,excludes= "`
	NickName     string `json:"nickName" binding:"required,excludes= "`
}
