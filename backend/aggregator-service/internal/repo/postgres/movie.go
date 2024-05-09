package postgres

import (
	"context"
	"encoding/json"
	"github.com/Enthreeka/aggregator-service/internal/repo"
	"github.com/Enthreeka/aggregator-service/pkg/postgres"
	pb "github.com/Entreeka/proto-proxy/go"
)

type postgresRepo struct {
	*postgres.Postgres
}

func NewPostgresRepo(pg *postgres.Postgres) repo.StoragePostgres {
	return &postgresRepo{
		pg,
	}
}

func (p *postgresRepo) SearchByText(ctx context.Context, text string) ([]*pb.Movie, error) {
	q := `SELECT movie FROM movies WHERE title ILIKE '%' || $1 || '%'`

	movies := make([]*pb.Movie, 0)

	rows, err := p.Pool.Query(ctx, q, text)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var movieJSON []byte
		if err := rows.Scan(&movieJSON); err != nil {
			return nil, err
		}

		var movie pb.Movie
		if err := json.Unmarshal(movieJSON, &movie); err != nil {
			return nil, err
		}
		movies = append(movies, &movie)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}

func (p *postgresRepo) GetByArrID(ctx context.Context, id []int64) ([]*pb.Movie, error) {
	sql := "SELECT movie FROM movies WHERE id = ANY($1)"

	args := []interface{}{id}

	rows, err := p.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	movieArr := make([]*pb.Movie, 0)
	for rows.Next() {
		var movieJSON []byte
		err := rows.Scan(&movieJSON)
		if err != nil {
			return nil, err
		}

		var movie pb.Movie
		err = json.Unmarshal(movieJSON, &movie)
		if err != nil {
			return nil, err
		}

		movieArr = append(movieArr, &movie)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return movieArr, nil
}

func (p *postgresRepo) CreateMovie(ctx context.Context, id int64, title string, movie []byte) error {
	q := `INSERT INTO movies (id,title,movie) VALUES ($1,$2,$3)`

	_, err := p.Pool.Exec(ctx, q, id, title, movie)
	return err
}

func (p *postgresRepo) IfExistsMovie(ctx context.Context, id int64) (bool, error) {
	q := `SELECT EXISTS(SELECT id FROM movies WHERE id = $1)`
	var exists bool

	err := p.Pool.QueryRow(ctx, q, id).Scan(&exists)
	return exists, err
}

//func (p *postgresRepo) SearchByText(ctx context.Context, text string) ([]entity.Movie, error) {
//	q := `SELECT * FROM movies WHERE title ILIKE '%' || $1 || '%'`
//
//	movies := make([]entity.Movie, 0)
//
//	rows, err := p.Pool.Query(ctx, q, text)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	for rows.Next() {
//		var movie entity.Movie
//		if err := rows.Scan(
//			&movie.ID,
//			&movie.Title,
//			&movie.Overview,
//			&movie.ReleaseDate,
//			&movie.Runtime,
//			&movie.Budget,
//			&movie.OriginalLanguage,
//			&movie.TrailerYT,
//		); err != nil {
//			return nil, err
//		}
//		movies = append(movies, movie)
//	}
//
//	if err := rows.Err(); err != nil {
//		return nil, err
//	}
//
//	return movies, nil
//}
