
type User {
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
