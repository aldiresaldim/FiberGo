package Entity

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-" gorm:"column:password"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt" grom:"index,column:deletd_at"`
}
