package service

type service struct {
	repo Repository
}

func New(repo Repository) *service {
	return &service{
		repo: repo,
	}
}

type Repository interface{}
