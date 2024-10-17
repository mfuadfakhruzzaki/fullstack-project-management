// models/comment.go
package models

import (
	"time"
)

type Comment struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    TaskID    uint      `gorm:"not null" json:"task_id"`
    Task      Task      `gorm:"foreignKey:TaskID" json:"task"`
    AuthorID  uint      `gorm:"not null" json:"author_id"`
    Author    User      `gorm:"foreignKey:AuthorID" json:"author"`
    Content   string    `gorm:"not null" json:"content"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
