package models

import (
    "github.com/jinzhu/gorm"
    "github.com/plexmediamanager/micro-database/proto"
    "time"
)

type Genre struct {
    ID                          uint64              `json:"id" gorm:"primary_key"`
    Name                        string              `json:"name" gorm:"size:255;not null"`
    Timestamps
}

// Provide the table name for Gorm
func (genre *Genre) TableName() string {
    return "genres"
}

// Get primary key of the table
func (genre *Genre) PrimaryKey() uint64 {
    return genre.ID
}

// Action which will be executed before new model is added to the database
func (genre *Genre) BeforeCreate(transaction *gorm.DB) error {
    return nil
}

// Action which will be executed before model is updated
func (genre *Genre) BeforeUpdate(transaction *gorm.DB) error {
    return nil
}

// Action which will be executed before model entry is deleted from the database
func (genre *Genre) BeforeDelete(transaction *gorm.DB) error {
    return nil
}

// Convert model to proto
func (genre *Genre) ToProto() *proto.Genre {
    return &proto.Genre {
        Id:                     genre.ID,
        Name:                   genre.Name,
        CreatedAt:              genre.CreatedAt.Unix(),
        UpdatedAt:              genre.UpdatedAt.Unix(),
    }
}

// Extract model from proto
func (genre *Genre) FromProto(protoGenre *proto.Genre) {
    genre.ID                        =   protoGenre.Id
    genre.Name                      =   protoGenre.Name
    genre.CreatedAt                 =   time.Unix(protoGenre.CreatedAt, 0).UTC()
    genre.UpdatedAt                 =   time.Unix(protoGenre.UpdatedAt, 0).UTC()
}

const (
    KeyGenreID                      =   "id"
    KeyGenreName                    =   "name"
    KeyGenreCreatedAt               =   "created_at"
    KeyGenreUpdatedAt               =   "updated_at"
)

// Get list of updated fields
func (genre *Genre) UpdatedFields(updateID bool, updatedFields ...string) map[string]interface{} {
    fieldValueMap := make(map[string]interface{})
    for _, fieldName := range updatedFields {
        switch fieldName {
        case KeyGenreID:
            if updateID {
                fieldValueMap[fieldName] = genre.ID
            }
        case KeyGenreName:
            fieldValueMap[fieldName] = genre.Name
        case KeyGenreCreatedAt:
            fieldValueMap[fieldName] = genre.CreatedAt
        case KeyGenreUpdatedAt:
            fieldValueMap[fieldName] = genre.UpdatedAt
        }
    }
    return fieldValueMap
}