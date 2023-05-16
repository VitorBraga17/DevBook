package models

import "time"

type Usuario struct {
	ID       uint64    `json:"id,omitempty"` //,omitempty para uso de DTOs
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}