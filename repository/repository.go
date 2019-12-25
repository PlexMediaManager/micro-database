package repository

type UpdatedFields map[string]interface{}

var (
    Movie                   =   MovieRepository{}
    Permission              =   PermissionRepository{}
    User                    =   UserRepository{}
)
