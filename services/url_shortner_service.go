package services

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"

	"time"
)

type Url struct {
	url   string
	views int32
}

type UrlShortnerService struct {
	db *sql.DB
}

func NewUrlShortnerService(db *sql.DB) *UrlShortnerService {
	return &UrlShortnerService{
		db: db,
	}
}

// todo add error
func (svc *UrlShortnerService) NewUrl(ct context.Context, data UploadUrlDTO) (*UploadUrlResponseDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `
	INSERT INTO urls (url, shortened_url, views, created_at, updated_at)
	values ($1, $2, $3, $4, $5) returning *
`

	hash_url := sha256.Sum256([]byte(data.Payload))
	url_id := hex.EncodeToString(hash_url[:3])

	_, err := svc.db.ExecContext(
		ctx,
		query,
		data.Payload,
		url_id,
		0,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return nil, err
	}

	return &UploadUrlResponseDTO{
		ShortenedUrl: url_id,
		ViewCount:    0,
	}, nil
}

func (svc *UrlShortnerService) GetUrl(ct context.Context, data GetUrlDTO) (*GetUrlResponseDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := "SELECT url, views FROM urls WHERE shortened_url = $1"
	row := svc.db.QueryRowContext(ctx, query, data.ShortenedUrl)

	dest := Url{}

	err := row.Scan(&dest.url, &dest.views)

	if err != nil {
		return nil, err
	}

	update_views_query := `
	UPDATE urls SET views = $1, updated_at = $2
	WHERE shortened_url = $3
`

	svc.db.ExecContext(
		ctx,
		update_views_query, dest.views+1, time.Now(), data.ShortenedUrl)

	return &GetUrlResponseDTO{
		TrueUrl:   dest.url,
		ViewCount: dest.views + 1,
	}, nil
}
