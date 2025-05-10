package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"taskEffectiveMobile/internal/app/models"
)

type Repository interface {
	CreatePerson(ctx context.Context, person *models.Person) error
	GetPersonByID(ctx context.Context, id uuid.UUID) (*models.Person, error)
	UpdatePerson(ctx context.Context, person *models.Person) error
	DeletePerson(ctx context.Context, id uuid.UUID) error
	GetPeople(ctx context.Context, filter models.PeopleFilter) ([]models.Person, error)
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

func (r *repository) CreatePerson(ctx context.Context, person *models.Person) error {
	query := `INSERT INTO people (id, full_name, age, gender, nationality, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := r.db.Exec(ctx, query,
		person.ID,
		person.FullName,
		person.Age,
		person.Gender,
		person.Nationality,
		person.CreatedAt,
		person.UpdatedAt)

	return err
}

func (r *repository) GetPersonByID(ctx context.Context, id uuid.UUID) (*models.Person, error) {
	var person models.Person

	row := r.db.QueryRow(ctx, `SELECT id, full_name, age, gender, nationality, created_at, updated_at FROM people WHERE id = $1`, id)

	err := row.Scan(&person.ID, &person.FullName, &person.Age, &person.Gender, &person.Nationality, &person.CreatedAt, &person.UpdatedAt)

	return &person, err
}

func (r *repository) UpdatePerson(ctx context.Context, person *models.Person) error {
	_, err := r.db.Exec(ctx, `UPDATE people SET full_name = $1, age = $2, gender = $3, nationality=$4 WHERE id = $5`, person.FullName, person.Age, person.Gender, person.Nationality, person.ID)

	return err
}

func (r *repository) DeletePerson(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.Exec(ctx, `DELETE FROM people WHERE id = $1`, id)
	return err
}

func (r *repository) GetPeople(ctx context.Context, filter models.PeopleFilter) ([]models.Person, error) {
	query := `SELECT id, full_name, age, gender, nationality, created_at, updated_at FROM people WHERE 1=1`
	args := []interface{}{}
	argIndex := 1

	if filter.Gender != nil {
		query += fmt.Sprintf(" AND LOWER(gender) = LOWER($%d)", argIndex)
		args = append(args, *filter.Gender)
		argIndex++
	}

	if filter.Nationality != nil {
		query += fmt.Sprintf(" AND LOWER(nationality) = LOWER($%d)", argIndex)
		args = append(args, *filter.Nationality)
		argIndex++
	}

	if filter.Age != nil {
		query += fmt.Sprintf(" AND age = $%d", argIndex)
		args = append(args, *filter.Age)
		argIndex++
	}

	limit := filter.Limit
	if limit == 0 {
		limit = 10
	}
	offset := (filter.Page - 1) * limit
	if offset < 0 {
		offset = 0
	}

	query += fmt.Sprintf(" ORDER BY created_at DESC LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var people []models.Person
	for rows.Next() {
		var p models.Person

		err := rows.Scan(
			&p.ID,
			&p.FullName,
			&p.Age,
			&p.Gender,
			&p.Nationality,
			&p.CreatedAt,
			&p.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		people = append(people, p)
	}

	return people, nil
}
