package courior

type couriorRepository interface {
}

type Service struct {
	couriorRepository couriorRepository
}

func New(
	sr couriorRepository,
) Service {
	return Service{
		couriorRepository: sr,
	}
}
