package product

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProduct
}

func (s *Service) Get(idx int) (*Product, error) {
	return &allProduct[idx], nil
}
