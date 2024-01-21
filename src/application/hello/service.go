package hello

type Service interface {
	Hello(name string) string
}

type service struct {
}

func NewService() Service {
	return service{}
}

func (s service) Hello(name string) string {
	return "Hello " + name
}
