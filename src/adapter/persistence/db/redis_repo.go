package db

import (
	"co/note-server/src/config"
	"co/note-server/src/domain/model"
	"errors"
	"sync"

	"github.com/go-redis/redis"
)

type RedisRepository struct {
	rdb  *redis.Client
	lock *sync.Mutex
}

func MakeRedisRepository() RedisRepository {
	properties := config.MakeConfigProvider()
	host := properties.GetProperty("redis.host")
	port := properties.GetProperty("redis.port")

	rdb := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: "",
		DB:       0,
	})
	return RedisRepository{rdb: rdb, lock: &sync.Mutex{}}
}

func (r RedisRepository) GetAll() ([]model.Note, error) {
	if keys, err := r.rdb.Keys("*").Result(); err != nil {
		return nil, err
	} else {
		notes := []model.Note{}
		for _, key := range keys {
			if note, err := r.GetById(key); err != nil {
				return nil, err
			} else {
				notes = append(notes, note)
			}
		}
		return notes, nil
	}
}

func (r RedisRepository) GetById(id string) (model.Note, error) {
	if note, err := r.rdb.Get(id).Result(); err == redis.Nil {
		return model.MakeInvalidNote(), errors.New("Note with id '" + id + "' does not exist")
	} else if err != nil {
		return model.MakeInvalidNote(), err
	} else {
		return model.FromJson(note)
	}
}

func (r RedisRepository) Add(note model.Note) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	if jsn, err := note.ToJson(); err != nil {
		return err
	} else {
		if r.noteExists(note.ID) {
			return errors.New("Note with id '" + note.ID + "' already exists.")
		}
		return r.rdb.Set(note.ID, jsn, 0).Err()
	}
}

func (r RedisRepository) DeleteById(id string) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	if r.noteExists(id) {
		return r.rdb.Del(id).Err()
	} else {
		return errors.New("Note with id '" + id + "' does not exist.")
	}
}

func (r RedisRepository) noteExists(id string) bool {
	n, _ := r.GetById(id)
	return n.ID != model.MakeInvalidNote().ID
}
