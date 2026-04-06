package repository

import (
	"database/sql"
	"biblioteca-api/fcm/domain/repository"
)

type fcmTokenRepositoryImpl struct {
	db *sql.DB
}

func NewFCMTokenRepositoryPostgres(db *sql.DB) repository.FCMTokenRepository {
	return &fcmTokenRepositoryImpl{db: db}
}

func (r *fcmTokenRepositoryImpl) Save(usuarioID int, token string) error {
	query := `
		INSERT INTO fcm_tokens (usuario_id, token, created_at, updated_at) 
		VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		ON CONFLICT (usuario_id, token) 
		DO UPDATE SET updated_at = CURRENT_TIMESTAMP
	`
	_, err := r.db.Exec(query, usuarioID, token)
	return err
}

func (r *fcmTokenRepositoryImpl) Delete(usuarioID int, token string) error {
	query := `DELETE FROM fcm_tokens WHERE usuario_id = $1 AND token = $2`
	_, err := r.db.Exec(query, usuarioID, token)
	return err
}

func (r *fcmTokenRepositoryImpl) DeleteByUsuarioID(usuarioID int) error {
	query := `DELETE FROM fcm_tokens WHERE usuario_id = $1`
	_, err := r.db.Exec(query, usuarioID)
	return err
}

func (r *fcmTokenRepositoryImpl) GetTokensByUsuarioID(usuarioID int) ([]string, error) {
	query := `SELECT token FROM fcm_tokens WHERE usuario_id = $1`
	rows, err := r.db.Query(query, usuarioID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tokens []string
	for rows.Next() {
		var token string
		if err := rows.Scan(&token); err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}
