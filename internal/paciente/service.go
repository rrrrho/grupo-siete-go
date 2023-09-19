package paciente

import "errors"

type Service struct {
	repository RepositoryInterface
}

func NewService(repository RepositoryInterface) *Service {
	return &Service{repository: repository}
}

type ServiceInterface interface {
	GetByID(id int) (Paciente, error)
	Save(pacienteInput Paciente) (Paciente, error)
	Update(id int, pacienteInput Paciente) (Paciente, error)
	Delete(id int) error
}

var (
	ErrPacienteNotFound = errors.New("paciente not found")
	ErrSavePaciente     = errors.New("error saving paciente")
	ErrUpdatePaciente   = errors.New("error updating paciente")
	ErrDeletePaciente   = errors.New("error deleting paciente")
)

func (s *Service) GetByID(id int) (Paciente, error) {
	paciente, err := s.repository.GetByID(id)
	if err != nil {
		return Paciente{}, ErrPacienteNotFound
	}
	return paciente, nil
}

func (s *Service) Save(pacienteInput Paciente) (Paciente, error) {
	paciente, err := s.repository.Save(pacienteInput)
	if err != nil {
		return Paciente{}, ErrSavePaciente
	}
	return paciente, nil
}

func (s *Service) Update(id int, pacienteInput Paciente) (Paciente, error) {
	pacienteToUpdate, err := s.GetByID(id)
	if err != nil {
		return Paciente{}, ErrPacienteNotFound
	}
	if pacienteInput.Nombre != "" {
		pacienteToUpdate.Nombre = pacienteInput.Nombre
	}
	if pacienteInput.Apellido != "" {
		pacienteToUpdate.Apellido = pacienteInput.Apellido
	}
	if pacienteInput.Domicilio != "" {
		pacienteToUpdate.Domicilio = pacienteInput.Domicilio
	}
	if pacienteInput.DNI != "" {
		pacienteToUpdate.DNI = pacienteInput.DNI
	}
	if pacienteInput.FechaDeAlta != "" {
		pacienteToUpdate.FechaDeAlta = pacienteInput.FechaDeAlta
	}
	paciente, err := s.repository.Modify(id, pacienteToUpdate)
	if err != nil {
		return Paciente{}, ErrUpdatePaciente
	}
	paciente.ID = id
	return paciente, nil
}
func (s *Service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return ErrDeletePaciente
	}
	return nil
}
