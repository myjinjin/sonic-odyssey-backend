package usecase

type SignUpInput struct {
	Email    string
	Password string
	Name     string
	Nickname string
}

type SignUpOutput struct {
	UserID uint
}

type GetUserByIDOutput struct {
	ID              uint
	Email           string
	Name            string
	Nickname        string
	ProfileImageURL string
	Bio             string
	Website         string
}

type PatchUserInput struct {
	Name     *string
	Nickname *string
	Bio      *string
	Website  *string
}
