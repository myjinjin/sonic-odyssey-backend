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

type UpdatePasswordInput struct {
	UserID       uint
	CurrPassword string
	NewPassword  string
}

type SearchTrackOutput struct {
	Tracks []Track
	Total  int
}

type Track struct {
	ID      string
	Name    string
	Artists []Artist
}

type Artist struct {
	ID   string
	Name string
}
