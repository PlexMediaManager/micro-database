package service

import (
    "encoding/json"
    format "fmt"
    "github.com/plexmediamanager/micro-database/errors"
    "github.com/plexmediamanager/micro-database/proto"
    serviceHandler "github.com/plexmediamanager/service"
)

const (
    userServiceName                 =   "micro.database"
)

// Get service name
func getServiceName(serviceName string) string {
    if application, done := serviceHandler.FromContext(); done {
        return format.Sprintf("%s.%s-%s", application.Vendor(), serviceName, application.Environment())
    }
    panic("[Database::Service::getServiceName] This was not supposed to happen")
}

// Convert database response to model
func protoToStructure(output interface{}, result *proto.DatabaseResponse, err error) error {
    if err != nil {
        return err
    }
    err = json.Unmarshal(result.Result, &output)
    if err != nil {
        return errors.DatabaseUnmarshalError.ToError(err)
    }
    return nil
}
