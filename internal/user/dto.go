package user

type UserDto struct {
	Id        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func (dto UserDto) MapFromModel(user UserModel) UserDto {
	dto.Id = user.Id
	dto.FirstName = user.FirstName
	dto.LastName = user.LastName
	dto.Email = user.Email

	return dto
}

type AddUserDto struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (dto AddUserDto) MapToModel() (UserModel, error) {
	return NewUser(
		dto.FirstName,
		dto.LastName,
		dto.Email,
		dto.Password,
	)
}

type UpdateUserDto struct {
	Id        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type ChangeUserPasswordDto struct {
	Id       int64  `json:"id"`
	Password string `json:"password"`
}
