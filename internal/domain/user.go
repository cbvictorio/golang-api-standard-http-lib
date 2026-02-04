package domain

import "time"

type UserRole string

const (
	RoleAdmin    UserRole = "admin"
	RoleCustomer UserRole = "customer"
	RoleVendor   UserRole = "vendor"
)

type User struct {
	ID        string `json:"id" gorm:"primaryKey;type:varchar(36)"`
	CreatedAt time.Time
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime:false"`
	Name      string     `json:"name"`
	Email     string     `json:"email" gorm:"unique"`
	Password  string     `json:"password"`
	Role      UserRole   `json:"role" gorm:"type:user_role"`
}
