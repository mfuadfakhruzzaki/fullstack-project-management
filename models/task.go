// models/task.go
package models

import (
	"time"
)

type Task struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    ProjectID   uint      `gorm:"not null" json:"project_id"`
    Project     Project   `gorm:"foreignKey:ProjectID" json:"project"`
    Name        string    `gorm:"not null" json:"name"`
    Description string    `json:"description"`
    AssignedTo  uint      `json:"assigned_to"`
    Assignee    User      `gorm:"foreignKey:AssignedTo" json:"assignee"`
    Status      string    `gorm:"not null" json:"status"` // e.g., ongoing, completed
    Deadline    time.Time `json:"deadline"`
    Comments    []Comment `gorm:"foreignKey:TaskID" json:"comments,omitempty"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
