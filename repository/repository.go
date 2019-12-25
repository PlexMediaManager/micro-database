package repository

type UpdatedFields map[string]interface{}

var (
    Genre                   =   GenreRepository{}
    Language                =   LanguageRepository{}
    Movie                   =   MovieRepository{}
    Permission              =   PermissionRepository{}
    User                    =   UserRepository{}
)
