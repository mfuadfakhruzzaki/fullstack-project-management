// models/user.go
package models

import (
	"time"
)

type User struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
    Username  string         `gorm:"unique;not null" json:"username"`
    Email     string         `gorm:"unique;not null" json:"email"`
    Password  string         `gorm:"not null" json:"-"`
    Role      string         `gorm:"not null" json:"role"`
    Projects  []Project      `gorm:"foreignKey:OwnerID" json:"projects,omitempty"`
    Tasks     []Task         `gorm:"foreignKey:AssignedTo" json:"tasks,omitempty"`
    Comments  []Comment      `gorm:"foreignKey:AuthorID" json:"comments,omitempty"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
}
