package database

import (
    format "fmt"
    "github.com/plexmediamanager/micro-database/models"
    tmdbService "github.com/plexmediamanager/micro-tmdb/service"
)

// Synchronize database with TMDB sources
func (client *DatabaseClient) SynchronizeWithTMDB() error {
    var err error
    tmdbError := tmdbSynchronizeLanguages()
    if tmdbError != nil {
        format.Println("TMDB Service was not ready when we started database, not a big problem")
    }

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
            err = CreateRecord(CreationParameters {
                Model:        model,
                CreateWithID: false,
            })
            if err != nil {
                return err
            }
        }
    }
    return nil
}