package service

import (
    "context"
    microClient "github.com/micro/go-micro/client"
    "github.com/plexmediamanager/micro-database/models"
    "github.com/plexmediamanager/micro-database/proto"
)

func genreToStructure(result *proto.DatabaseResponse, err error) (*models.Genre, error) {
    var record *models.Genre
    return record, protoToStructure(&record, result, err)
}

func genreArrayToStructure(result *proto.DatabaseResponse, err error) ([]*models.Genre, error) {
    var record []*models.Genre
    return record, protoToStructure(&record, result, err)
}

func GetGenreService(client microClient.Client) proto.GenreService {
    return proto.NewGenreService(getServiceName(genreServiceName), client)
}

func GenreServiceCreate(client microClient.Client, model *models.Genre) (*models.Genre, error) {
    service := GetGenreService(client)
    return genreToStructure(service.Create(context.TODO(), model.ToProto()))
}

func GenreServiceUpdate(client microClient.Client, model *models.Genre, updatedFields ...string) (*models.Genre, error) {
    service := GetGenreService(client)
    protoModel := model.ToProto()
    protoModel.UpdatedFields = updatedFields
    return genreToStructure(service.Create(context.TODO(), protoModel))
}

func GenreServiceSoftDelete(client microClient.Client, model *models.Genre) (*models.Genre, error) {
    service := GetGenreService(client)
    return genreToStructure(service.SoftDelete(context.TODO(), model.ToProto()))
}

func GenreServiceForceDelete(client microClient.Client, model *models.Genre) (*models.Genre, error) {
    service := GetGenreService(client)
    return genreToStructure(service.ForceDelete(context.TODO(), model.ToProto()))
}

func GenreServiceFindAll(client microClient.Client) ([]*models.Genre, error) {
    service := GetGenreService(client)
    parameters := &proto.DatabaseEmpty {}
    return genreArrayToStructure(service.FindAll(context.TODO(), parameters))
}

func GenreServiceFindOneByID(client microClient.Client, id uint64) (*models.Genre, error) {
    service := GetGenreService(client)
    parameters := &proto.DatabaseID { Value: id }
    return genreToStructure(service.FindOneByID(context.TODO(), parameters))
}

func GenreServiceFindManyByID(client microClient.Client, id ...uint64) ([]*models.Genre, error) {
    service := GetGenreService(client)
    parameters := &proto.DatabaseIDs { Values: id }
    return genreArrayToStructure(service.FindManyByID(context.TODO(), parameters))
}

func GenreServiceFindOneByName(client microClient.Client, title string) (*models.Genre, error) {
    service := GetGenreService(client)
    parameters := &proto.DatabaseString { Value: title }
    return genreToStructure(service.FindOneByName(context.TODO(), parameters))
}

func GenreServiceFindManyByName(client microClient.Client, title ...string) ([]*models.Genre, error) {
    service := GetGenreService(client)
    parameters := &proto.DatabaseStrings { Values: title }
    return genreArrayToStructure(service.FindManyByName(context.TODO(), parameters))
}