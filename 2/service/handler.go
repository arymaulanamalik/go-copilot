package service

// create a new handler
// request service

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

// Handler is a handler of service
type Handler struct {
	// service *Service

	// or
	// service Service

	// or
	// service ServiceInterface

	service Service
}

// Create a new library
// request Library
// response error
func (h *Handler) CreateLibrary(request Library) error {
	//call the service
	// and then handle the error
	err := h.service.CreateLibrary(request)
	if err != nil {
		return err
	}

	return nil
}

// View List Library
// response []Library
func (h *Handler) ViewLibrary() []Library {
	// call the service
	library := h.service.ViewLibrary()

	return library
}
