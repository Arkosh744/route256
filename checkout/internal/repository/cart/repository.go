package cart

type repository struct{}

func NewRepo() *repository {
	return &repository{}
}
