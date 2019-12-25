package repository

import (
    "github.com/plexmediamanager/micro-database/database"
    "github.com/plexmediamanager/micro-database/models"
)

type MovieRepository struct {}

// Create new movie
func (*MovieRepository) Create(model *models.Movie) (*models.Movie, error) {
    return model, database.CreateRecord(model)
}

// Update existing movie
func (*MovieRepository) Update(model *models.Movie, fields UpdatedFields) (*models.Movie, error) {
    return model, database.UpdateRecord(model, fields)
}

// Soft delete existing movie
func (*MovieRepository) SoftDelete(model *models.Movie) (*models.Movie, error) {
    return model, database.SoftDeleteRecord(model)
}

// Force delete existing movie
func (*MovieRepository) ForceDelete(model *models.Movie) (*models.Movie, error) {
    return model, database.ForceDeleteRecord(model)
}

// Get list of all downloaded movies
func (repository *MovieRepository) FindDownloaded() ([]*models.Movie, error) {
    return repository.findMany(models.KeyMovieLocalTitle + " IS NOT NULL", nil)
}

// Find single movie
func (*MovieRepository) findOne(query interface{}) (*models.Movie, error) {
    movie := &models.Movie{}
    err := database.GetSingleRecord(movie, query)
    return movie, err
}

// Find multiple movies
func (*MovieRepository) findMany(query interface{}, parameters interface{}) ([]*models.Movie, error) {
    var movies []*models.Movie
    err := database.GetMultipleRecords(&movies, query, parameters)
    return movies, err
}