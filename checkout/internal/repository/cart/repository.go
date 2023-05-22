package cart

type Repository interface{}

type repository struct{}

func NewRepo() *repository {
	return &repository{}
}
