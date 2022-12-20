package user

type CreateNewUserViewModel struct {
	LastName      string `validate:"required,min=3,max=30"`
	FirstName     string `validate:"required,min=3,max=30"`
	BirthDate     string `validate:"required,datetime=2006/01/02"`
	Code          string `validate:"required,min=0,max=10"`
	UserName      string
	Password      string
	AvatarName    string
	CreatorUserId string
}
