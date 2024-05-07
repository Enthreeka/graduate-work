package postgres

import (
	"context"
	"fmt"
	"github.com/Enthreeka/aggregator-service/internal/entity"
	"github.com/Enthreeka/aggregator-service/internal/repo"
	"github.com/Enthreeka/aggregator-service/pkg/postgres"
	"github.com/lib/pq"
	"strings"
)

type postgresRepo struct {
	*postgres.Postgres
}

func NewRedisRepo(pg *postgres.Postgres) repo.StoragePostgres {
	return &postgresRepo{
		pg,
	}
}

func (p *postgresRepo) SearchByText(ctx context.Context, text string) ([]entity.Movie, error) {
	q := `SELECT * FROM movies WHERE title ILIKE '%' || $1 || '%'`

	movies := make([]entity.Movie, 0)

	rows, err := p.Pool.Query(ctx, q, text)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var movie entity.Movie
		if err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Overview,
			&movie.ReleaseDate,
			&movie.Runtime,
			&movie.Budget,
			&movie.OriginalLanguage,
			&movie.TrailerYT,
		); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}

func (p *postgresRepo) GetByArrID(ctx context.Context, id []int) ([]entity.Movie, error) {
	q := `select * from movie where id = unnest($1)`

	movies := make([]entity.Movie, 0)
	idArrStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(id)), ","), "[]")

	err := p.Pool.QueryRow(ctx, q, idArrStr).Scan(pq.Array(&movies))
	if err != nil {
		return nil, err
	}

	return movies, nil
}
