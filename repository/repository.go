package repository

type UpdatedFields map[string]interface{}

var (
    Genre                   =   GenreRepository{}
    Movie                   =   MovieRepository{}
    Permission              =   PermissionRepository{}
    User                    =   UserRepository{}
)
