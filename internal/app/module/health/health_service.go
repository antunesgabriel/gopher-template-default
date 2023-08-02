package health

import "context"

type HealthService struct {
	repository HealthRepository
}

func NewHealthService(r HealthRepository) *HealthService {
	s := HealthService{
		repository: r,
	}

	return &s
}

func (s *HealthService) Check() (*health, error) {
	ctx := context.Background()

	if err := s.repository.Ping(ctx); err != nil {
		return nil, err
	}

	return NewHealth("up"), nil
}
