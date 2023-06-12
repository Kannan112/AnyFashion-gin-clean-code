package req

type SuperAdmin struct {
	Email    string ` json:"email" validate:"required" binding:"email"`
	Password string ` json:"password" validate:"required"`
}

type CreateAdmin struct {
	ID       uint   `json:"id" gorm:"primaryKey;not null"`
	Name     string ` json:"name" validate:"required"`
	Email    string ` json:"email" validate:"required" binding:"email"`
	Password string ` json:"password" validate:"required"`
}
type BlockData struct {
	UserId uint   ` json:"userid" validate:"required"`
	Reason string ` json:"reason" validate:"required"`
}
type UserEmail struct {
	Email string `json:"email"`
}
