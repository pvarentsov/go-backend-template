package repository

type Repositories struct {
	User UserRepository
}

func NewRepositories(queries Queries) Repositories {
	return Repositories{
		User: NewUserRepository(queries),
	}
}
