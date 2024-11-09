package repository

type Promocode interface {
}

type Reward interface {
}

type Repository struct {
	Promocode
	Reward
}

func NewRepository() *Repository {
	return &Repository{}
}
