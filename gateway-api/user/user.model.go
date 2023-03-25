package user

type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`;
	Email string `json:"email" binding:"required,email"`;
	Password string `json:"password" binding:"required,min=6"`
}

type User struct {
	Id      uint `gorm:"primaryKey"`
	Name    string
	Email   string  `gorm:"unique;not null"`
	Is_Active bool `gorm:"default=true"`
	Password string `gorm:"unique;not null"`
	Password_changed_at  int64 `gorm:"autoUpdateTime:seconds"`
	Updated int64 `gorm:"autoUpdateTime:seconds"` // Use unix seconds as updating time
	Created int64 `gorm:"autoCreateTime:seconds"` // Use unix seconds as creating time
}