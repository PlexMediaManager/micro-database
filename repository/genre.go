package repository

import (
    "github.com/plexmediamanager/micro-database/database"
    "github.com/plexmediamanager/micro-database/models"
)

type GenreRepository struct {}

// Create new genre
func (*GenreRepository) Create(model *models.Genre) (*models.Genre, error) {
    return model, database.CreateRecord(model)
}

// Update existing genre
func (*GenreRepository) Update(model *models.Genre, fields UpdatedFields) (*models.Genre, error) {
    return model, database.UpdateRecord(model, fields)
}

// Soft delete existing genre
func (*GenreRepository) SoftDelete(model *models.Genre) (*models.Genre, error) {
    return model, database.SoftDeleteRecord(model)
}

// Force delete existing genre
func (*GenreRepository) ForceDelete(model *models.Genre) (*models.Genre, error) {
    return model, database.ForceDeleteRecord(model)
}

// Get all genres from the database
func (*GenreRepository) FindAll() ([]*models.Genre, error) {
    var genres []*models.Genre
    return genres, database.GetAllRecords(&genres)
}

// Find one genre by ID
func (repository *GenreRepository) FindOneByID(id uint64) (*models.Genre, error) {
    return repository.findOne(models.Genre{ ID: id })
}

// Find many genres by ID
func (repository *GenreRepository) FindManyByID(id ...uint64) ([]*models.Genre, error) {
    return repository.findMany(models.KeyGenreID + " IN (?)", id)
}

// Find one genre by Title
func (repository *GenreRepository) FindOneByName(name string) (*models.Genre, error) {
    return repository.findOne(models.Genre{ Name: name })
}

// Find many genres by Title
func (repository *GenreRepository) FindManyByName(name ...string) ([]*models.Genre, error) {
    return repository.findMany(models.KeyGenreName + " IN (?)", name)
}

// Find single genre
func (*GenreRepository) findOne(query interface{}) (*models.Genre, error) {
    genre := &models.Genre{}
    err := database.GetSingleRecord(genre, query)
    return genre, err
}

// Find multiple genres
func (*GenreRepository) findMany(query interface{}, parameters interface{}) ([]*models.Genre, error) {
    var genres []*models.Genre
    err := database.GetMultipleRecords(&genres, query, parameters)
    return genres, err
}