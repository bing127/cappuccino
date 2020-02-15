package dto


type GetAccessTokenRequest struct {
	LoginName string `json:"login_name" binding:"required,excludes= "`
	Password string `json:"password" binding:"required"`
}

type GetAccessTokenResponse struct {
	AccessToken string `json:"accessToken" binding:"required,excludes= "`
}

type UserLoginInfo struct {
	ID            uint64 `json:"id" binding:"required"`
	LoginName      string `json:"username" binding:"required,excludes= "`
	Token         string `json:"token"`
	NickName         string `json:"nickname"`
}
