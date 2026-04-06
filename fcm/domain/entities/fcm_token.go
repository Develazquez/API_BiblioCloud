package entities

import "time"

type FCMToken struct {
	ID        int       `json:"id"`
	UsuarioID int       `json:"usuario_id"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
