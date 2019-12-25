package models

import (
    "encoding/json"
    "github.com/jinzhu/gorm"
    "github.com/plexmediamanager/micro-database/proto"
    "time"
)

type Movie struct {
    ID                          uint64              `json:"id" gorm:"primary_key"`
    Title                       string              `json:"title" gorm:"size:255;not null"`
    OriginalTitle               string              `json:"original_title" gorm:"size:255;not null"`
    LocalTitle                  string              `json:"local_title" gorm:"size:255"`
    OriginalLanguage            string              `json:"original_language" gorm:"size:255"`
    Languages                   json.RawMessage     `json:"languages" gorm:"type:json" sql:"type:longtext"`
    Overview                    string              `json:"overview" sql:"type:longtext"`
    Tagline                     string              `json:"tagline" gorm:"size:255"`
    Genres                      json.RawMessage     `json:"genres" gorm:"type:json" sql:"type:longtext"`
    Homepage                    string              `json:"homepage" gorm:"size:255"`
    Runtime                     uint64              `json:"runtime" gorm:"size:10"`
    Status                      uint64              `json:"status" gorm:"size:10"`
    Adult                       bool                `json:"adult" gorm:"default:false"`
    ImdbId                      string              `json:"imdb_id" gorm:"size:255"`
    ReleaseDate                 string              `json:"release_date" gorm:"size:255"`
    ProductionCompanies         json.RawMessage     `json:"production_companies" gorm:"type:json" sql:"type:longtext"`
    ProductionCountries         json.RawMessage     `json:"production_countries" gorm:"type:json" sql:"type:longtext"`
    VoteAverage                 float64             `json:"vote_average" sql:"type:double"`
    VoteCount                   uint64              `json:"vote_count" gorm:"size:10"`
    Popularity                  float64             `json:"popularity" sql:"type:double"`
    Budget                      uint64              `json:"budget" gorm:"size:10"`
    Revenue                     uint64              `json:"revenue" gorm:"size:10"`
    Backdrop                    string              `json:"backdrop" gorm:"size:255"`
    Poster                      string              `json:"poster" gorm:"size:255"`
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

// Convert model to proto
func (movie *Movie) ToProto() *proto.Movie {
    return &proto.Movie {
        Id:                     movie.ID,
        Title:                  movie.Title,
        OriginalTitle:          movie.OriginalTitle,
        LocalTitle:             movie.LocalTitle,
        OriginalLanguage:       movie.OriginalLanguage,
        Languages:              movie.Languages,
        Overview:               movie.Overview,
        Tagline:                movie.Tagline,
        Genres:                 movie.Genres,
        Homepage:               movie.Homepage,
        Runtime:                movie.Runtime,
        Status:                 movie.Status,
        Adult:                  movie.Adult,
        ImdbId:                 movie.ImdbId,
        ReleaseDate:            movie.ReleaseDate,
        ProductionCompanies:    movie.ProductionCompanies,
        ProductionCountries:    movie.ProductionCountries,
        VoteAverage:            movie.VoteAverage,
        VoteCount:              movie.VoteCount,
        Popularity:             movie.Popularity,
        Budget:                 movie.Budget,
        Revenue:                movie.Revenue,
        Backdrop:               movie.Backdrop,
        Poster:                 movie.Poster,
        CreatedAt:              movie.CreatedAt.Unix(),
        UpdatedAt:              movie.UpdatedAt.Unix(),
    }
}

// Extract model from proto
func (movie *Movie) FromProto(protoMovie *proto.Movie) {
    movie.ID                        =   protoMovie.Id
    movie.Title                     =   protoMovie.Title
    movie.OriginalTitle             =   protoMovie.OriginalTitle
    movie.LocalTitle                =   protoMovie.LocalTitle
    movie.OriginalLanguage          =   protoMovie.OriginalLanguage
    movie.Languages                 =   protoMovie.Languages
    movie.Overview                  =   protoMovie.Overview
    movie.Tagline                   =   protoMovie.Tagline
    movie.Genres                    =   protoMovie.Genres
    movie.Homepage                  =   protoMovie.Homepage
    movie.Runtime                   =   protoMovie.Runtime
    movie.Status                    =   protoMovie.Status
    movie.Adult                     =   protoMovie.Adult
    movie.ImdbId                    =   protoMovie.ImdbId
    movie.ReleaseDate               =   protoMovie.ReleaseDate
    movie.ProductionCompanies       =   protoMovie.ProductionCompanies
    movie.ProductionCountries       =   protoMovie.ProductionCountries
    movie.VoteAverage               =   protoMovie.VoteAverage
    movie.VoteCount                 =   protoMovie.VoteCount
    movie.Popularity                =   protoMovie.Popularity
    movie.Budget                    =   protoMovie.Budget
    movie.Revenue                   =   protoMovie.Revenue
    movie.Backdrop                  =   protoMovie.Backdrop
    movie.Poster                    =   protoMovie.Poster
    movie.CreatedAt                 =   time.Unix(protoMovie.CreatedAt, 0).UTC()
    movie.UpdatedAt                 =   time.Unix(protoMovie.UpdatedAt, 0).UTC()
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