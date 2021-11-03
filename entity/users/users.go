package users

type User interface {
	ID() string
}

type user struct {
	id string
}

func NewUser(id string) User {
	return &user{
		id: id,
	}
}

func (u *user) ID() string {
	return u.id
}
