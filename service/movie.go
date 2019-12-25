package service

import (
    "context"
    microClient "github.com/micro/go-micro/client"
    "github.com/plexmediamanager/micro-database/models"
    "github.com/plexmediamanager/micro-database/proto"
)

func movieToStructure(result *proto.DatabaseResponse, err error) (*models.Movie, error) {
    var record *models.Movie
    return record, protoToStructure(&record, result, err)
}

func movieArrayToStructure(result *proto.DatabaseResponse, err error) ([]*models.Movie, error) {
    var record []*models.Movie
    return record, protoToStructure(&record, result, err)
}

func GetMovieService(client microClient.Client) proto.MovieService {
    return proto.NewMovieService(getServiceName(movieServiceName), client)
}

func MovieServiceCreate(client microClient.Client, model *models.Movie) (*models.Movie, error) {
    service := GetMovieService(client)
    return movieToStructure(service.Create(context.TODO(), model.ToProto()))
}

func MovieServiceUpdate(client microClient.Client, model *models.Movie, updatedFields ...string) (*models.Movie, error) {
    service := GetMovieService(client)
    protoModel := model.ToProto()
    protoModel.UpdatedFields = updatedFields
    return movieToStructure(service.Create(context.TODO(), protoModel))
}

func MovieServiceSoftDelete(client microClient.Client, model *models.Movie) (*models.Movie, error) {
    service := GetMovieService(client)
    return movieToStructure(service.SoftDelete(context.TODO(), model.ToProto()))
}

func MovieServiceForceDelete(client microClient.Client, model *models.Movie) (*models.Movie, error) {
    service := GetMovieService(client)
    return movieToStructure(service.ForceDelete(context.TODO(), model.ToProto()))
}

func MovieServiceFindAll(client microClient.Client) ([]*models.Movie, error) {
    service := GetMovieService(client)
    parameters := &proto.DatabaseEmpty {}
    return movieArrayToStructure(service.FindAll(context.TODO(), parameters))
}

func MovieServiceFindDownloaded (client microClient.Client) ([]*models.Movie, error) {
    service := GetMovieService(client)
    parameters := &proto.DatabaseEmpty {}
    return movieArrayToStructure(service.FindDownloaded(context.TODO(), parameters))
}

func MovieServiceFindOneByID(client microClient.Client, id uint64) (*models.Movie, error) {
    service := GetMovieService(client)
    parameters := &proto.DatabaseID { Value: id }
    return movieToStructure(service.FindOneByID(context.TODO(), parameters))
}

func MovieServiceFindManyByID(client microClient.Client, id ...uint64) ([]*models.Movie, error) {
    service := GetMovieService(client)
    parameters := &proto.DatabaseIDs { Values: id }
    return movieArrayToStructure(service.FindManyByID(context.TODO(), parameters))
}

func MovieServiceFindOneByTitle(client microClient.Client, title string) (*models.Movie, error) {
    service := GetMovieService(client)
    parameters := &proto.DatabaseString { Value: title }
    return movieToStructure(service.FindOneByTitle(context.TODO(), parameters))
}

func MovieServiceFindManyByTitle(client microClient.Client, title ...string) ([]*models.Movie, error) {
    service := GetMovieService(client)
    parameters := &proto.DatabaseStrings { Values: title }
    return movieArrayToStructure(service.FindManyByTitle(context.TODO(), parameters))
}

func MovieServiceFindOneByOriginalTitle(client microClient.Client, originalTitle string) (*models.Movie, error) {
    service := GetMovieService(client)
    parameters := &proto.DatabaseString { Value: originalTitle }
    return movieToStructure(service.FindOneByOriginalTitle(context.TODO(), parameters))
}

func MovieServiceFindManyByOriginalTitle(client microClient.Client, originalTitle ...string) ([]*models.Movie, error) {
    service := GetMovieService(client)
    parameters := &proto.DatabaseStrings { Values: originalTitle }
    return movieArrayToStructure(service.FindManyByOriginalTitle(context.TODO(), parameters))
}

func MovieServiceFindOneByLocalTitle(client microClient.Client, localTitle string) (*models.Movie, error) {
    service := GetMovieService(client)
    parameters := &proto.DatabaseString { Value: localTitle }
    return movieToStructure(service.FindOneByLocalTitle(context.TODO(), parameters))
}

func MovieServiceFindManyByLocalTitle(client microClient.Client, localTitle ...string) ([]*models.Movie, error) {
    service := GetMovieService(client)
    parameters := &proto.DatabaseStrings { Values: localTitle }
    return movieArrayToStructure(service.FindManyByLocalTitle(context.TODO(), parameters))
}
