package repo

import (
	"log"

	"backend/server/internal/config"
	"backend/server/internal/models"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(cfg *config.Config) (*Repository, error) {
	db, err := sqlx.Open("pgx", cfg.DbPath)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	log.Println("DataBase Connection established")
	return &Repository{db: db}, nil
}

func (r *Repository) GetAllBins() ([]models.BinBank, error) {
	var data []models.BinBank
	err := r.db.Select(&data, `SELECT bin, bank FROM bin`)
	log.Println("data dumped")
	return data, err
}

func (r *Repository) InsertBin(bin int, bank string) error {
	_, err := r.db.Exec(`
        INSERT INTO bin (bin, bank) 
        VALUES ($1, $2)
        ON CONFLICT (bin) 
        DO UPDATE SET bank = EXCLUDED.bank
    `, bin, bank)
	return err
}
