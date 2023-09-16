package turno

type Repository interface {
	AddTurno(turno Turno) (Turno, error)
	GetTurnoByID(id int) (Turno, error)
	ReplaceTurno(turno Turno) (Turno, error)
	UpdateTurno(turno Turno) (Turno, error)
	DeleteTurno(id int) (string, error)
	GetTurnosByDNI(dni string) ([]Turno, error)
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) AddTurno(turno Turno) (Turno, error) {
	return s.repository.AddTurno(turno)
}

func (s *Service) GetTurnoByID(id int) (Turno, error) {
	return s.repository.GetTurnoByID(id)
}

func (s *Service) ReplaceTurno(turno Turno) (Turno, error) {
	return s.repository.ReplaceTurno(turno)
}

func (s *Service) UpdateTurno(turno Turno) (Turno, error) {
	return s.repository.UpdateTurno(turno)
}

func (s *Service) DeleteTurno(id int) (string, error) {
	return s.repository.DeleteTurno(id)
}

func (s *Service) GetTurnosByDNI(dni string) ([]Turno, error) {
	return s.repository.GetTurnosByDNI(dni)
}
