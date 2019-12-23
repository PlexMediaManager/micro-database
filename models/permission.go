package models

import (
    "github.com/jinzhu/gorm"
)

type Permission struct {
    ID                  uint64          `gorm:"primary_key"`
    Name                string          `gorm:"size:255"`
    Timestamps
}

// Provide the table name for Gorm
func (permission *Permission) TableName() string {
    return "permissions"
}

// Get primary key of the table
func (permission *Permission) PrimaryKey() uint64 {
    return permission.ID
}

// Action which will be executed before new model is added to the database
func (permission *Permission) BeforeCreate(transaction *gorm.DB) error {
    return nil
}

// Action which will be executed before model is updated
func (permission *Permission) BeforeUpdate(transaction *gorm.DB) error {
    return nil
}

// Action which will be executed before model entry is deleted from the database
func (permission *Permission) BeforeDelete(transaction *gorm.DB) error {
    return nil
}

const (
    KeyPermissionID                     =   "id"
    KeyPermissionName                   =   "name"
    KeyPermissionCreatedAt              =   "created_at"
    KeyPermissionUpdatedAt              =   "updated_at"
)

// Get list of updated fields
func (permission *Permission) UpdatedFields(updateID bool, updatedFields ...string) map[string]interface{} {
    fieldValueMap := make(map[string]interface{}, 0)
    for _, fieldName := range updatedFields {
        switch fieldName {
            case KeyPermissionID:
                if updateID {
                    fieldValueMap[fieldName] = permission.ID
                }
                break
            case KeyPermissionName:
                fieldValueMap[fieldName] = permission.Name
                break
            case KeyPermissionCreatedAt:
                fieldValueMap[fieldName] = permission.CreatedAt
                break
            case KeyPermissionUpdatedAt:
                fieldValueMap[fieldName] = permission.UpdatedAt
                break
        }
    }
    return fieldValueMap
}