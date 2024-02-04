package postgres

import (
	"context"
	"errors"

	"github.com/bookoo-billy/jukebox/db"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	v1 "github.com/bookoo-billy/jukebox/proto/api/v1"
	"github.com/google/uuid"
)

type ReceiverDAO struct {
	db *gorm.DB
}

func NewReceiverDAO(db *gorm.DB) db.ReceiverDAO {
	return &ReceiverDAO{db: db}
}

func (a *ReceiverDAO) Create(ctx context.Context, receiver *v1.Receiver) (*v1.Receiver, error) {
	err := a.validate(receiver)
	if err != nil {
		return nil, err
	}

	receiver.Id = uuid.New().String()
	tx := a.db.WithContext(ctx).Create(receiver)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return receiver, nil
}

func (a *ReceiverDAO) Delete(ctx context.Context, query *v1.ReceiverQuery) (*v1.Receiver, error) {
	deleted := &v1.Receiver{}
	tx := a.db.Model(deleted).Clauses(clause.Returning{}).WithContext(ctx).Delete(query)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return deleted, nil
}

func (a *ReceiverDAO) Get(ctx context.Context, query *v1.ReceiverQuery) (*v1.Receiver, error) {
	receiver := &v1.Receiver{}
	tx := a.db.Model(receiver).WithContext(ctx).First(query)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return receiver, nil
}

func (a *ReceiverDAO) List(ctx context.Context, query *v1.ReceiverQuery) (*v1.ReceiverList, error) {
	var items []*v1.Receiver

	tx := a.db.Model(items).WithContext(ctx).Find(query)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &v1.ReceiverList{Items: items}, nil
}

func (a *ReceiverDAO) Update(ctx context.Context, receiver *v1.Receiver) (*v1.Receiver, error) {
	if receiver.Id == "" {
		return nil, errors.New("receiver.Id is a required field")
	}

	updated := &v1.Receiver{}
	tx := a.db.Model(updated).WithContext(ctx).Clauses(clause.Returning{}).Updates(receiver)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return updated, nil
}

func (a *ReceiverDAO) validate(receiver *v1.Receiver) error {
	if receiver.Name == "" {
		return errors.New("receiver.Name is a required field")
	}

	return nil
}
