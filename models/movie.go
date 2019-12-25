package models

import (
    "encoding/json"
    "github.com/jinzhu/gorm"
)

type Movie struct {
    ID                          uint64              `gorm:"primary_key"`
    Title                       string              `gorm:"size:255;not null"`
    OriginalTitle               string              `gorm:"size:255;not null"`
    LocalTitle                  *string             `gorm:"size:255"`
    OriginalLanguage            *string             `gorm:"size:255"`
    Languages                   json.RawMessage     `gorm:"type:json",sql:"type:longtext"`
    Overview                    string              `sql:"type:longtext"`
    Tagline                     string              `gorm:"size:255"`
    Genres                      json.RawMessage     `gorm:"type:json",sql:"type:longtext"`
    Homepage                    string              `gorm:"size:255"`
    Runtime                     uint64              `gorm:"size:10"`
    Status                      uint64              `gorm:"size:10"`
    Adult                       bool                `gorm:"default:false"`
    ImdbId                      string              `gorm:"size:255"`
    ReleaseDate                 string              `gorm:"size:255"`
    ProductionCompanies         json.RawMessage     `gorm:"type:json",sql:"type:longtext"`
    ProductionCountries         json.RawMessage     `gorm:"type:json",sql:"type:longtext"`
    VoteAverage                 float64             `sql:"type:double"`
    VoteCount                   uint64              `gorm:"size:10"`
    Popularity                  float64             `sql:"type:double"`
    Budget                      uint64              `gorm:"size:10"`
    Revenue                     uint64              `gorm:"size:10"`
    Backdrop                    string              `gorm:"size:255"`
    Poster                      string              `gorm:"size:255"`
    Timestamps
}

// Provide the table name for Gorm
func (movie *Movie) TableName() string {
    return "movies"
}

// Get primary key of the table
func (movie *Movie) PrimaryKey() uint64 {
    return movie.ID
}

// Action which will be executed before new model is added to the database
func (movie *Movie) BeforeCreate(transaction *gorm.DB) error {
    return nil
}

// Action which will be executed before model is updated
func (movie *Movie) BeforeUpdate(transaction *gorm.DB) error {
    return nil
}

// Action which will be executed before model entry is deleted from the database
func (movie *Movie) BeforeDelete(transaction *gorm.DB) error {
    return nil
}

const (
    KeyMovieID                      =   "id"
    KeyMovieTitle                   =   "title"
    KeyMovieOriginalTitle           =   "original_title"
    KeyMovieLocalTitle              =   "local_title"
    KeyMovieOriginalLanguage        =   "original_language"
    KeyMovieLanguages               =   "languages"
    KeyMovieOverview                =   "overview"
    KeyMovieTagline                 =   "tagline"
    KeyMovieGenres                  =   "genres"
    KeyMovieHomepage                =   "homepage"
    KeyMovieRuntime                 =   "runtime"
    KeyMovieStatus                  =   "status"
    KeyMovieAdult                   =   "adult"
    KeyMovieImdbId                  =   "imdb_id"
    KeyMovieReleaseDate             =   "release_date"
    KeyMovieProductionCompanies     =   "production_companies"
    KeyMovieProductionCountries     =   "production_countries"
    KeyMovieVoteAverage             =   "vote_average"
    KeyMovieVoteCount               =   "vote_count"
    KeyMoviePopularity              =   "popularity"
    KeyMovieBudget                  =   "budget"
    KeyMovieRevenue                 =   "revenue"
    KeyMovieBackdrop                =   "backdrop"
    KeyMoviePoster                  =   "poster"
    KeyMovieCreatedAt               =   "created_at"
    KeyMovieUpdatedAt               =   "updated_at"
)

// Get list of updated fields
func (movie *Movie) UpdatedFields(updateID bool, updatedFields ...string) map[string]interface{} {
    fieldValueMap := make(map[string]interface{})
    for _, fieldName := range updatedFields {
        switch fieldName {
            case KeyMovieID:
                if updateID {
                    fieldValueMap[fieldName] = movie.ID
                }
            case KeyMovieTitle:
                fieldValueMap[fieldName] = movie.Title
            case KeyMovieOriginalTitle:
                fieldValueMap[fieldName] = movie.OriginalTitle
            case KeyMovieLocalTitle:
                fieldValueMap[fieldName] = movie.LocalTitle
            case KeyMovieOriginalLanguage:
                fieldValueMap[fieldName] = movie.OriginalLanguage
            case KeyMovieLanguages:
                fieldValueMap[fieldName] = movie.Languages
            case KeyMovieOverview:
                fieldValueMap[fieldName] = movie.Overview
            case KeyMovieTagline:
                fieldValueMap[fieldName] = movie.Tagline
            case KeyMovieGenres:
                fieldValueMap[fieldName] = movie.Genres
            case KeyMovieHomepage:
                fieldValueMap[fieldName] = movie.Homepage
            case KeyMovieRuntime:
                fieldValueMap[fieldName] = movie.Runtime
            case KeyMovieStatus:
                fieldValueMap[fieldName] = movie.Status
            case KeyMovieAdult:
                fieldValueMap[fieldName] = movie.Adult
            case KeyMovieImdbId:
                fieldValueMap[fieldName] = movie.ImdbId
            case KeyMovieReleaseDate:
                fieldValueMap[fieldName] = movie.ReleaseDate
            case KeyMovieProductionCompanies:
                fieldValueMap[fieldName] = movie.ProductionCompanies
            case KeyMovieProductionCountries:
                fieldValueMap[fieldName] = movie.ProductionCountries
            case KeyMovieVoteAverage:
                fieldValueMap[fieldName] = movie.VoteAverage
            case KeyMovieVoteCount:
                fieldValueMap[fieldName] = movie.VoteCount
            case KeyMoviePopularity:
                fieldValueMap[fieldName] = movie.Popularity
            case KeyMovieBudget:
                fieldValueMap[fieldName] = movie.Budget
            case KeyMovieRevenue:
                fieldValueMap[fieldName] = movie.Revenue
            case KeyMovieBackdrop:
                fieldValueMap[fieldName] = movie.Backdrop
            case KeyMoviePoster:
                fieldValueMap[fieldName] = movie.Poster
            case KeyMovieCreatedAt:
                fieldValueMap[fieldName] = movie.CreatedAt
            case KeyMovieUpdatedAt:
                fieldValueMap[fieldName] = movie.UpdatedAt
        }
    }
    return fieldValueMap
}