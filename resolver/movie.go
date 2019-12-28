package resolver

import (
    "context"
    "github.com/plexmediamanager/micro-database/models"
    "github.com/plexmediamanager/micro-database/proto"
    "github.com/plexmediamanager/micro-database/repository"
)

type MovieService struct {}

func (service MovieService) Create (_ context.Context, parameters *proto.Movie, response *proto.DatabaseResponse) error {
    instance := models.Movie{}
    instance.FromProto(parameters)
    result, err := repository.Movie.Create(&instance)
    return structureToBytesWithError(result, err, response)
}

func (service MovieService) Update (_ context.Context, parameters *proto.Movie, response *proto.DatabaseResponse) error {
    instance := models.Movie{}
    instance.FromProto(parameters)
    fields := instance.UpdatedFields(false, parameters.UpdatedFields...)
    result, err := repository.Movie.Update(&instance, fields)
    return structureToBytesWithError(result, err, response)
}

func (service MovieService) SoftDelete (_ context.Context, parameters *proto.Movie, response *proto.DatabaseResponse) error {
    instance := models.Movie{}
    instance.FromProto(parameters)
    result, err := repository.Movie.SoftDelete(&instance)
    return structureToBytesWithError(result, err, response)
}

func (service MovieService) ForceDelete (_ context.Context, parameters *proto.Movie, response *proto.DatabaseResponse) error {
    instance := models.Movie{}
    instance.FromProto(parameters)
    result, err := repository.Movie.ForceDelete(&instance)
    return structureToBytesWithError(result, err, response)
}

func (service MovieService) FindAll(_ context.Context, parameters *proto.DatabaseEmpty, response *proto.DatabaseResponse) error {
    result, err := repository.Movie.FindAll()
    return structureToBytesWithError(result, err, response)
}

func (service MovieService) FindOnlyIDs(_ context.Context, parameters *proto.DatabaseEmpty, response *proto.DatabaseResponse) error {
    result, err := repository.Movie.FindOnlyIDs()
    return structureToBytesWithError(result, err, response)
}

func (service MovieService) FindDownloaded (_ context.Context, parameters *proto.DatabaseEmpty, response *proto.DatabaseResponse) error {
    result, err := repository.Movie.FindDownloaded()
    return structureToBytesWithError(result, err, response)
}

func (service MovieService) FindOneByID (_ context.Context, parameters *proto.DatabaseID, response *proto.DatabaseResponse) error {
    result, err := repository.Movie.FindOneByID(parameters.Value)
    return structureToBytesWithError(result, err, response)
}

func (service MovieService) FindManyByID (_ context.Context, parameters *proto.DatabaseIDs, response *proto.DatabaseResponse) error {
    result, err := repository.Movie.FindManyByID(parameters.Values...)
    return structureToBytesWithError(result, err, response)
}

func (service MovieService) FindOneByTitle (_ context.Context, parameters *proto.DatabaseString, response *proto.DatabaseResponse) error {
    result, err := repository.Movie.FindOneByTitle(parameters.Value)
    return structureToBytesWithError(result, err, response)
}

func (service MovieService) FindManyByTitle (_ context.Context, parameters *proto.DatabaseStrings, response *proto.DatabaseResponse) error {
    result, err := repository.Movie.FindManyByTitle(parameters.Values...)
    return structureToBytesWithError(result, err, response)
}

func (service MovieService) FindOneByOriginalTitle (_ context.Context, parameters *proto.DatabaseString, response *proto.DatabaseResponse) error {
    result, err := repository.Movie.FindOneByOriginalTitle(parameters.Value)
    return structureToBytesWithError(result, err, response)
}

func (service MovieService) FindManyByOriginalTitle (_ context.Context, parameters *proto.DatabaseStrings, response *proto.DatabaseResponse) error {
    result, err := repository.Movie.FindManyByOriginalTitle(parameters.Values...)
    return structureToBytesWithError(result, err, response)
}

func (service MovieService) FindOneByLocalTitle (_ context.Context, parameters *proto.DatabaseString, response *proto.DatabaseResponse) error {
    result, err := repository.Movie.FindOneByLocalTitle(parameters.Value)
    return structureToBytesWithError(result, err, response)
}

func (service MovieService) FindManyByLocalTitle (_ context.Context, parameters *proto.DatabaseStrings, response *proto.DatabaseResponse) error {
    result, err := repository.Movie.FindManyByLocalTitle(parameters.Values...)
    return structureToBytesWithError(result, err, response)
}
