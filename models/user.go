package models

import (
    "github.com/jinzhu/gorm"
    "time"
)

type User struct {
    ID                  uint64          `gorm:"primary_key"`
    Username            string          `gorm:"unique;size:255"`
    Email               string          `gorm:"unique;size:255"`
    EmailVerifiedAt     *time.Time      `gorm:"default:null"`
    Password            string          `gorm:"size:255"`
    RememberToken       string          `gorm:"size:100"`
    Timestamps
}

// Provide the table name for Gorm
func (user *User) TableName() string {
    return "users"
}

// Get primary key of the table
func (user *User) PrimaryKey() uint64 {
    return user.ID
}

// Action which will be executed before new model is added to the database
func (user *User) BeforeCreate(transaction *gorm.DB) error {
    return nil
}

// Action which will be executed before model is updated
func (user *User) BeforeUpdate(transaction *gorm.DB) error {
    return nil
}

// Action which will be executed before model entry is deleted from the database
func (user *User) BeforeDelete(transaction *gorm.DB) error {
    return nil
}

const (
    KeyUserID                       =   "id"
    KeyUserUsername                 =   "username"
    KeyUserEmail                    =   "email"
    KeyUserEmailVerifiedAt          =   "email_verified_at"
    KeyUserPassword                 =   "password"
    KeyUserRememberToken            =   "remember_token"
    KeyUserCreatedAt                =   "created_at"
    KeyUserUpdatedAt                =   "updated_at"
    //KeyUserDeletedAt                =   "deleted_at"
)

// Get list of updated fields
func (user *User) UpdatedFields(updateID bool, updatedFields ...string) map[string]interface{} {
    fieldValueMap := make(map[string]interface{}, 0)
    for _, fieldName := range updatedFields {
        switch fieldName {
        case KeyUserID:
            if updateID {
                fieldValueMap[fieldName] = user.ID
            }
            break
        case KeyUserUsername:
            fieldValueMap[fieldName] = user.Username
            break
        case KeyUserEmail:
            fieldValueMap[fieldName] = user.Email
            break
        case KeyUserEmailVerifiedAt:
            fieldValueMap[fieldName] = user.EmailVerifiedAt
        case KeyUserPassword:
            fieldValueMap[fieldName] = user.Password
            break
        case KeyUserRememberToken:
            fieldValueMap[fieldName] = user.RememberToken
        case KeyUserCreatedAt:
            fieldValueMap[fieldName] = user.CreatedAt
            break
        case KeyUserUpdatedAt:
            fieldValueMap[fieldName] = user.UpdatedAt
            break
        //case KeyUserDeletedAt:
        //    fieldValueMap[fieldName] = user.DeletedAt
        //    break
        }
    }
    return fieldValueMap
}