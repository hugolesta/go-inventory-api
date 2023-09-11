package models

type UserRole struct {
	ID     int64  `json:"id"`
	UserID int64  `json:"user_id"`
	RoleID int64  `json:"role_id"`
}
