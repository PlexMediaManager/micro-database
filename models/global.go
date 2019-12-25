package models

import "time"

type Timestamps struct {
    CreatedAt           time.Time           `json:"created_at"`
    UpdatedAt           time.Time           `json:"updated_at"`
}

type SoftDeleteTimestamps struct {
    Timestamps
    DeletedAt           *time.Time           `json:"deleted_at",gorm:"default:null"`
}

func (timestamp *Timestamps) BeforeSave() error {
    if timestamp.CreatedAt.IsZero() {
        timestamp.CreatedAt = time.Now()
    }
    timestamp.UpdatedAt = time.Now()
    return nil
}

func (timestamp *Timestamps) BeforeUpdate() error {
    timestamp.UpdatedAt = time.Now()
    return nil
}