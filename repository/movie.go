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

// Get all movies from the database
func (*MovieRepository) FindAll() ([]*models.Movie, error) {
    var movies []*models.Movie
    return movies, database.GetAllRecords(&movies)
}

// Get list of all downloaded movies
func (repository *MovieRepository) FindDownloaded() ([]*models.Movie, error) {
    return repository.findMany(models.KeyMovieLocalTitle + " IS NOT NULL", nil)
}

// Find one movie by ID
func (repository *MovieRepository) FindOneByID(id uint64) (*models.Movie, error) {
    return repository.findOne(models.Movie{ ID: id })
}

// Find many movies by ID
func (repository *MovieRepository) FindManyByID(id ...uint64) ([]*models.Movie, error) {
    return repository.findMany(models.KeyMovieID + " IN (?)", id)
}

// Find one movie by Title
func (repository *MovieRepository) FindOneByTitle(title string) (*models.Movie, error) {
    return repository.findOne(models.Movie{ Title: title })
}

// Find many movies by Title
func (repository *MovieRepository) FindManyByTitle(title ...string) ([]*models.Movie, error) {
    return repository.findMany(models.KeyMovieTitle + " IN (?)", title)
}

// Find one movie by Original Title
func (repository *MovieRepository) FindOneByOriginalTitle(title string) (*models.Movie, error) {
    return repository.findOne(models.Movie{ OriginalTitle: title })
}

// Find many movies by Original Title
func (repository *MovieRepository) FindManyByOriginalTitle(title ...string) ([]*models.Movie, error) {
    return repository.findMany(models.KeyMovieOriginalTitle + " IN (?)", title)
}

// Find one movie by Local Title
func (repository *MovieRepository) FindOneByLocalTitle(title string) (*models.Movie, error) {
    return repository.findOne(models.Movie{ LocalTitle: title })
}

// Find many movies by Local Title
func (repository *MovieRepository) FindManyByLocalTitle(title ...string) ([]*models.Movie, error) {
    return repository.findMany(models.KeyMovieLocalTitle + " IN (?)", title)
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