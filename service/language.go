package service

import (
    "context"
    microClient "github.com/micro/go-micro/client"
    "github.com/plexmediamanager/micro-database/models"
    "github.com/plexmediamanager/micro-database/proto"
)

func languageToStructure(result *proto.DatabaseResponse, err error) (*models.Language, error) {
    var record *models.Language
    return record, protoToStructure(&record, result, err)
}

func languageArrayToStructure(result *proto.DatabaseResponse, err error) ([]*models.Language, error) {
    var record []*models.Language
    return record, protoToStructure(&record, result, err)
}

func GetLanguageService(client microClient.Client) proto.LanguageService {
    return proto.NewLanguageService(getServiceName(languageServiceName), client)
}

func LanguageServiceCreate(client microClient.Client, model *models.Language) (*models.Language, error) {
    service := GetLanguageService(client)
    return languageToStructure(service.Create(context.TODO(), model.ToProto()))
}

func LanguageServiceUpdate(client microClient.Client, model *models.Language, updatedFields ...string) (*models.Language, error) {
    service := GetLanguageService(client)
    protoModel := model.ToProto()
    protoModel.UpdatedFields = updatedFields
    return languageToStructure(service.Create(context.TODO(), protoModel))
}

func LanguageServiceSoftDelete(client microClient.Client, model *models.Language) (*models.Language, error) {
    service := GetLanguageService(client)
    return languageToStructure(service.SoftDelete(context.TODO(), model.ToProto()))
}

func LanguageServiceForceDelete(client microClient.Client, model *models.Language) (*models.Language, error) {
    service := GetLanguageService(client)
    return languageToStructure(service.ForceDelete(context.TODO(), model.ToProto()))
}

func LanguageServiceFindAll(client microClient.Client) ([]*models.Language, error) {
    service := GetLanguageService(client)
    parameters := &proto.DatabaseEmpty {}
    return languageArrayToStructure(service.FindAll(context.TODO(), parameters))
}

func LanguageServiceFindOneByISO(client microClient.Client, code string) (*models.Language, error) {
    service := GetLanguageService(client)
    parameters := &proto.DatabaseString { Value: code }
    return languageToStructure(service.FindOneByISO(context.TODO(), parameters))
}

func LanguageServiceFindManyByISO(client microClient.Client, code ...string) (*models.Language, error) {
    service := GetLanguageService(client)
    parameters := &proto.DatabaseStrings { Values: code }
    return languageToStructure(service.FindManyByISO(context.TODO(), parameters))
}