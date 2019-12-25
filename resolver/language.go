package resolver

import (
    "context"
    "github.com/plexmediamanager/micro-database/models"
    "github.com/plexmediamanager/micro-database/proto"
    "github.com/plexmediamanager/micro-database/repository"
)

type LanguageService struct {}

func (service LanguageService) Create (_ context.Context, parameters *proto.Language, response *proto.DatabaseResponse) error {
    instance := models.Language{}
    instance.FromProto(parameters)
    result, err := repository.Language.Create(&instance)
    return structureToBytesWithError(result, err, response)
}

func (service LanguageService) Update (_ context.Context, parameters *proto.Language, response *proto.DatabaseResponse) error {
    instance := models.Language{}
    instance.FromProto(parameters)
    fields := instance.UpdatedFields(false, parameters.UpdatedFields...)
    result, err := repository.Language.Update(&instance, fields)
    return structureToBytesWithError(result, err, response)
}

func (service LanguageService) SoftDelete (_ context.Context, parameters *proto.Language, response *proto.DatabaseResponse) error {
    instance := models.Language{}
    instance.FromProto(parameters)
    result, err := repository.Language.SoftDelete(&instance)
    return structureToBytesWithError(result, err, response)
}

func (service LanguageService) ForceDelete (_ context.Context, parameters *proto.Language, response *proto.DatabaseResponse) error {
    instance := models.Language{}
    instance.FromProto(parameters)
    result, err := repository.Language.ForceDelete(&instance)
    return structureToBytesWithError(result, err, response)
}

func (service LanguageService) FindAll(_ context.Context, parameters *proto.DatabaseEmpty, response *proto.DatabaseResponse) error {
    result, err := repository.Language.FindAll()
    return structureToBytesWithError(result, err, response)
}

func (service LanguageService) FindOneByISO (_ context.Context, parameters *proto.DatabaseString, response *proto.DatabaseResponse) error {
    result, err := repository.Language.FindOneByISO(parameters.Value)
    return structureToBytesWithError(result, err, response)
}
func (service LanguageService) FindManyByISO (_ context.Context, parameters *proto.DatabaseStrings, response *proto.DatabaseResponse) error {
    result, err := repository.Language.FindManyByISO(parameters.Values...)
    return structureToBytesWithError(result, err, response)
}