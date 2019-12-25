package resolver

import (
    "context"
    "github.com/plexmediamanager/micro-database/proto"
    "github.com/plexmediamanager/micro-database/repository"
)

type MovieService struct {}

func (service MovieService) FindDownloaded (_ context.Context, parameters *proto.DatabaseEmpty, response *proto.DatabaseResponse) error {
    result, err := repository.Movie.FindDownloaded()
    return structureToBytesWithError(result, err, response)
}