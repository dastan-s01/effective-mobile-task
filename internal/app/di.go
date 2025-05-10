package app

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"taskEffectiveMobile/internal/app/services/enricher"
	"taskEffectiveMobile/internal/app/services/person/repository"
	usecase "taskEffectiveMobile/internal/app/usecases"
)

type DI struct {
	PersonUseCase usecase.PersonUseCase
}

func NewDI(db *pgxpool.Pool) *DI {
	personRepo := repository.NewRepository(db)
	enricher := enricher.NewEnricher()
	personUseCase := usecase.NewPersonUsecase(personRepo, enricher)
	return &DI{
		PersonUseCase: personUseCase,
	}
}
