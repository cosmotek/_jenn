
type User {
	ID uuid.UUID
	FirstName string
	LastName string
	JoinedAt time.Time
	PhoneNumber string
}

func ArchiveUser(id string) error {
	return nil
}

func GetUser(id string) (User, error) {
	return User{}, nil
}

func UpdateUser(id string) (User, error) {
	return User{}, nil
}


type Cocktail {
	ID uuid.UUID
	Name string
}

func ArchiveCocktail(id string) error {
	return nil
}

func GetCocktail(id string) (Cocktail, error) {
	return Cocktail{}, nil
}

func UpdateCocktail(id string) (Cocktail, error) {
	return Cocktail{}, nil
}


