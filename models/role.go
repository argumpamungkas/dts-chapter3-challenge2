package models

type Role struct {
	ID       uint   `gorm:"primaryKey" json:"role_id"`
	RoleName string `gorm:"not null" json:"role_name"`
}

func (r *Role) TableName() string {
	return "tb_role"
}
