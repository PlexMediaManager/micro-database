package resolver

import (
    "encoding/json"
    "github.com/plexmediamanager/micro-database/errors"
    "github.com/plexmediamanager/micro-database/proto"
)

// Convert structure to bytes and return error if needed
func structureToBytesWithError(structure interface{}, err error, response *proto.DatabaseResponse) error {
    if err != nil {
        return err
    }
    bytes, err := json.Marshal(structure)
    if err != nil {
        return errors.DatabaseMarshalError.ToError(err)
    }
    response.Result = bytes
    return nil
}