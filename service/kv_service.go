package service

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/kundu-ramit/zocket/infra/database"
	"github.com/kundu-ramit/zocket/model"
)

type KeyValueService interface {
	CreateAuth(name string) (*model.Auth, error)
	CreateKeyValue(key, value string, ownedBy uuid.UUID) (*model.Keypair, error)
	GetKeyValue(key string, ownedBy uuid.UUID) (*model.Keypair, error)
	UpdateKeyValue(key, value string, ownedBy uuid.UUID) (*model.Keypair, error)
	DeleteKeyValue(key string, ownedBy uuid.UUID) error
}

type keyValueService struct {
	db *sql.DB
}

func NewKeyValueService() KeyValueService {
	return &keyValueService{
		db: database.Initialize(),
	}
}

func (s *keyValueService) CreateAuth(name string) (*model.Auth, error) {
	auth := &model.Auth{ID: uuid.New(), Name: name}
	query := "INSERT INTO auth (id, name) VALUES (?, ?)"
	_, err := s.db.Exec(query, auth.ID, auth.Name)
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func (s *keyValueService) CreateKeyValue(key, value string, ownedBy uuid.UUID) (*model.Keypair, error) {
	kv := &model.Keypair{ID: uuid.New(), Key: key, Value: value, OwnedBy: ownedBy}
	query := "INSERT INTO keypairs (id, key, value, owned_by) VALUES (?, ?, ?, ?)"
	_, err := s.db.Exec(query, kv.ID, kv.Key, kv.Value, kv.OwnedBy)
	if err != nil {
		return nil, err
	}
	return kv, nil
}

func (s *keyValueService) GetKeyValue(key string, ownedBy uuid.UUID) (*model.Keypair, error) {
	var kv model.Keypair
	query := "SELECT id, key, value, owned_by FROM keypairs WHERE key = ? AND owned_by = ?"
	row := s.db.QueryRow(query, key, ownedBy)
	if err := row.Scan(&kv.ID, &kv.Key, &kv.Value, &kv.OwnedBy); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("keypair not found")
		}
		return nil, err
	}
	return &kv, nil
}

func (s *keyValueService) UpdateKeyValue(key, value string, ownedBy uuid.UUID) (*model.Keypair, error) {
	kv, err := s.GetKeyValue(key, ownedBy)
	if err != nil {
		return nil, err
	}
	kv.Value = value
	query := "UPDATE keypairs SET value = ? WHERE key = ? AND owned_by = ?"
	_, err = s.db.Exec(query, kv.Value, kv.Key, kv.OwnedBy)
	if err != nil {
		return nil, err
	}
	return kv, nil
}

func (s *keyValueService) DeleteKeyValue(key string, ownedBy uuid.UUID) error {
	query := "DELETE FROM keypairs WHERE key = ? AND owned_by = ?"
	_, err := s.db.Exec(query, key, ownedBy)
	if err != nil {
		return err
	}
	return nil
}
