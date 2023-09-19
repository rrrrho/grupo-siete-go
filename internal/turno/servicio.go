package turno

type Repository interface {
	Save(turno Turno) (Turno, error)
	GetByID(id int) (Turno, error)
	Replace(turno Turno) (Turno, error)
	Update(id int, turno Turno) (Turno, error)
	Delete(id int) (string, error)
	GetByDNI(dni string) ([]Turno, error)
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Save(turno Turno) (Turno, error) {
	return s.repository.Save(turno)
}

func (s *Service) GetByID(id int) (Turno, error) {
	return s.repository.GetByID(id)
}

func (s *Service) Replace(turno Turno) (Turno, error) {
	return s.repository.Replace(turno)
}

func (s *Service) Update(id int, turno Turno) (Turno, error) {
	return s.repository.Update(id, turno)
}

func (s *Service) Delete(id int) (string, error) {
	return s.repository.Delete(id)
}

func (s *Service) GetByDNI(dni string) ([]Turno, error) {
	return s.repository.GetByDNI(dni)
}
