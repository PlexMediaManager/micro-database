package resolver

import (
    "context"
    "github.com/plexmediamanager/micro-database/models"
    "github.com/plexmediamanager/micro-database/proto"
    "github.com/plexmediamanager/micro-database/repository"
)

type GenreService struct {}

func (service GenreService) Create (_ context.Context, parameters *proto.Genre, response *proto.DatabaseResponse) error {
    instance := models.Genre{}
    instance.FromProto(parameters)
    result, err := repository.Genre.Create(&instance)
    return structureToBytesWithError(result, err, response)
}

func (service GenreService) Update (_ context.Context, parameters *proto.Genre, response *proto.DatabaseResponse) error {
    instance := models.Genre{}
    instance.FromProto(parameters)
    fields := instance.UpdatedFields(false, parameters.UpdatedFields...)
    result, err := repository.Genre.Update(&instance, fields)
    return structureToBytesWithError(result, err, response)
}

func (service GenreService) SoftDelete (_ context.Context, parameters *proto.Genre, response *proto.DatabaseResponse) error {
    instance := models.Genre{}
    instance.FromProto(parameters)
    result, err := repository.Genre.SoftDelete(&instance)
    return structureToBytesWithError(result, err, response)
}

func (service GenreService) ForceDelete (_ context.Context, parameters *proto.Genre, response *proto.DatabaseResponse) error {
    instance := models.Genre{}
    instance.FromProto(parameters)
    result, err := repository.Genre.ForceDelete(&instance)
    return structureToBytesWithError(result, err, response)
}

func (service GenreService) FindAll(_ context.Context, parameters *proto.DatabaseEmpty, response *proto.DatabaseResponse) error {
    result, err := repository.Genre.FindAll()
    return structureToBytesWithError(result, err, response)
}

func (service GenreService) FindOneByID (_ context.Context, parameters *proto.DatabaseID, response *proto.DatabaseResponse) error {
    result, err := repository.Genre.FindOneByID(parameters.Value)
    return structureToBytesWithError(result, err, response)
}

func (service GenreService) FindManyByID (_ context.Context, parameters *proto.DatabaseIDs, response *proto.DatabaseResponse) error {
    result, err := repository.Genre.FindManyByID(parameters.Values...)
    return structureToBytesWithError(result, err, response)
}

func (service GenreService) FindOneByName (_ context.Context, parameters *proto.DatabaseString, response *proto.DatabaseResponse) error {
    result, err := repository.Genre.FindOneByName(parameters.Value)
    return structureToBytesWithError(result, err, response)
}

func (service GenreService) FindManyByName (_ context.Context, parameters *proto.DatabaseStrings, response *proto.DatabaseResponse) error {
    result, err := repository.Genre.FindManyByName(parameters.Values...)
    return structureToBytesWithError(result, err, response)
}