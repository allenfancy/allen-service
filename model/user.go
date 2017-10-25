package model

// UserInfo user basic infomartion
type UserInfo struct {
	ID       int64  `json:"id"`
	UserName string `json:"userName"`
	PassWord string `json:"password"`
	TelPhone string `json:"telPhone"`
	Email    string `json:"email"`
	Slat     string `json:"slat"`
}
