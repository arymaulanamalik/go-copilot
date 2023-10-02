package service

// create a new service

func NewService() *Service {
	return &Service{}
}

// Service is a service
type Service struct {
}

// Create a new library
// request Library
// response error
func (s *Service) CreateLibrary(request Library) error {
	return nil
}

// View List Library
// response []Library
func (s *Service) ViewLibrary() []Library {
	//create dummy data
	library := []Library{
		{
			ID:   1,
			Name: "Book 1",
		},
		{
			ID:   2,
			Name: "Book 2",
		},
		{
			ID:   3,
			Name: "Book 3",
		},
	}

	return library
}
