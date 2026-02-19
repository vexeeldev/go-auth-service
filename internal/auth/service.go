package auth

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllUsers() ([]User, error) {
	return s.repo.GetAll()
}

func (s *Service) CreateUser(name, email string) (*User, error) {
	return s.repo.Create(name, email)
}

func (s *Service) GetUserByID(id int) (*User, error) {
	return s.repo.GetByID(id)
}

func (s *Service) DeleteUser(id int) error {
	return s.repo.Delete(id)
}

