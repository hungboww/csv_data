package models

import "time"

type User struct {
	Id          int       `json:"id"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	LastLogin   time.Time `json:"last_login"`
	IsSuperuser bool      `json:"is_superuser"`
	UserName    string    `json:"user_name"`
	FirstName   string    `json:"first_name"`
	StartDate   time.Time `json:"start_date"`
	About       string    `json:"about"`
	Image       string    `json:"image"`
	ISActive    bool      `json:"is_active"`
	IsStaff     bool      `json:"is_staff"`
	RoleId      int       `json:"role_id"`
}
type Role struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (Role) TableName() string {
	return "tbl_role"
}
func (User) TableName() string {
	return "tbl_user"
}
