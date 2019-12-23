package models

import "time"

type Timestamps struct {
    CreatedAt           time.Time
    UpdatedAt           time.Time
    DeletedAt           *time.Time           `gorm:"default:null"`
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