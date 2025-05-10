package usecase

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"taskEffectiveMobile/internal/app/models"
	"taskEffectiveMobile/internal/app/services/enricher"
	"taskEffectiveMobile/internal/app/services/person/repository"
	"time"
)

type PersonUseCase interface {
	Create(ctx context.Context, FullName string) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Person, error)
	Update(ctx context.Context, person *models.Person) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetPeople(ctx context.Context, filter models.PeopleFilter) ([]models.Person, error)
}

type personUsecase struct {
	repo     repository.Repository
	enricher enricher.Enricher
}

func NewPersonUsecase(repo repository.Repository, enricher enricher.Enricher) PersonUseCase {
	return &personUsecase{
		repo:     repo,
		enricher: enricher,
	}
}

func (u *personUsecase) Create(ctx context.Context, FullName string) error {

	id := uuid.New()
	createdAt := time.Now()
	updatedAt := time.Now()
	age, gender, nationality, err := u.enricher.Enrich(ctx, FullName)
	if err != nil {
		return err
	}
	fmt.Println(age, gender, nationality)
	person := &models.Person{
		ID:          id,
		FullName:    FullName,
		Age:         age,
		Gender:      gender,
		Nationality: nationality,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	return u.repo.CreatePerson(ctx, person)
}

func (u *personUsecase) GetByID(ctx context.Context, id uuid.UUID) (*models.Person, error) {
	return u.repo.GetPersonByID(ctx, id)
}

func (u *personUsecase) Update(ctx context.Context, p *models.Person) error {
	p.UpdatedAt = time.Now()
	return u.repo.UpdatePerson(ctx, p)
}

func (u *personUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.repo.DeletePerson(ctx, id)
}
func (u *personUsecase) GetPeople(ctx context.Context, filter models.PeopleFilter) ([]models.Person, error) {
	return u.repo.GetPeople(ctx, filter)

}
