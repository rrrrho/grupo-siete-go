package odontologo

type Repository interface {
	GetByID(id int) (Odontologo, error)
	Modify(id int, odontologo Odontologo) (Odontologo, error)
	Save(odontologo Odontologo) (Odontologo, error)
	Delete(id int) (Odontologo, error)
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetByID(id int) (Odontologo, error) {
	return s.repository.GetByID(id)
}

func (s *Service) Modify(id int, odontologo Odontologo) (Odontologo, error) {
	return s.repository.Modify(id, odontologo)
}

func (s *Service) Save(odontologo Odontologo) (Odontologo, error) {
	return s.repository.Save(odontologo)
}

func (s *Service) Delete(id int) (Odontologo, error) {
	return s.repository.Delete(id)
}

