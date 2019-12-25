package models

import (
    "github.com/jinzhu/gorm"
    "github.com/plexmediamanager/micro-database/proto"
    "time"
)

type Language struct {
    ID                          uint64              `json:"id" gorm:"primary_key"`
    Iso                         string              `json:"iso_639_1" gorm:"size:50;not null"`
    EnglishName                 string              `json:"english_name" gorm:"size:255;not null"`
    Name                        string              `json:"name" gorm:"size:255;not null"`
    Timestamps
}

// Provide the table name for Gorm
func (language *Language) TableName() string {
    return "languages"
}

// Get primary key of the table
func (language *Language) PrimaryKey() uint64 {
    return language.ID
}

// Action which will be executed before new model is added to the database
func (language *Language) BeforeCreate(transaction *gorm.DB) error {
    return nil
}

// Action which will be executed before model is updated
func (language *Language) BeforeUpdate(transaction *gorm.DB) error {
    return nil
}

// Action which will be executed before model entry is deleted from the database
func (language *Language) BeforeDelete(transaction *gorm.DB) error {
    return nil
}

// Convert model to proto
func (language *Language) ToProto() *proto.Language {
    return &proto.Language {
        Id:                     language.ID,
        Iso:                    language.Iso,
        EnglishName:            language.EnglishName,
        Name:                   language.Name,
        CreatedAt:              language.CreatedAt.Unix(),
        UpdatedAt:              language.UpdatedAt.Unix(),
    }
}

// Extract model from proto
func (language *Language) FromProto(protoLanguage *proto.Language) {
    language.ID                        =   protoLanguage.Id
    language.Iso                       =   protoLanguage.Iso
    language.EnglishName               =   protoLanguage.EnglishName
    language.Name                      =   protoLanguage.Name
    language.CreatedAt                 =   time.Unix(protoLanguage.CreatedAt, 0).UTC()
    language.UpdatedAt                 =   time.Unix(protoLanguage.UpdatedAt, 0).UTC()
}

const (
    KeyLanguageID                      =   "id"
    KeyLanguageIso                     =   "iso"
    KeyLanguageEnglishName             =   "english_name"
    KeyLanguageName                    =   "name"
    KeyLanguageCreatedAt               =   "created_at"
    KeyLanguageUpdatedAt               =   "updated_at"
)

// Get list of updated fields
func (language *Language) UpdatedFields(updateID bool, updatedFields ...string) map[string]interface{} {
    fieldValueMap := make(map[string]interface{})
    for _, fieldName := range updatedFields {
        switch fieldName {
            case KeyLanguageID:
                if updateID {
                    fieldValueMap[fieldName] = language.ID
                }
            case KeyLanguageIso:
                fieldValueMap[fieldName] = language.Iso
            case KeyLanguageEnglishName:
                fieldValueMap[fieldName] = language.EnglishName
            case KeyLanguageName:
                fieldValueMap[fieldName] = language.Name
            case KeyLanguageCreatedAt:
                fieldValueMap[fieldName] = language.CreatedAt
            case KeyLanguageUpdatedAt:
                fieldValueMap[fieldName] = language.UpdatedAt
        }
    }
    return fieldValueMap
}