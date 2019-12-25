package database

import (
    "github.com/plexmediamanager/micro-database/models"
    tmdbService "github.com/plexmediamanager/micro-tmdb/service"
)

// Synchronize database with TMDB sources
func (client *DatabaseClient) SynchronizeWithTMDB() error {
    err := tmdbSynchronizeLanguages()

    return err
}

func tmdbSynchronizeLanguages() error {
    languagesList, err := tmdbService.TMDBServiceGetLanguages(application.Service().Client())
    if err != nil {
        return err
    }
    for _, language := range languagesList {
        model := &models.Language{
            Iso:            language.ISO6391,
            EnglishName:    language.EnglishName,
            Name:           language.Name,
        }
        exists, err := RecordExists(model)
        if err != nil {
            return err
        }
        if !exists {
            err = CreateRecord(model)
            if err != nil {
                return err
            }
        }
    }
    return nil
}