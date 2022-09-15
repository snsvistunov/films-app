package repository

import (
	"errors"
	"time"
)

var errUserNotFound = errors.New("user not found")

func (r *AuthDB) SaveToken(userID []uint8, token string, ttl time.Duration) error {
	_, err := r.storage.Set(token, string(userID), ttl).Result()
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthDB) GetUserID(token string) (string, error) {
	get := r.storage.Get(token)

	userID := get.Val()
	err := get.Err()
	if err != nil {
		return "", err
	}

	return userID, nil
}

func (r *AuthDB) DeleteToken(token string) error {
	userID, err := r.GetUserID(token)
	if err != nil {
		return err
	}

	_, err = r.storage.Del(token).Result()
	if err != nil {
		return err
	}

	if userID == "" {
		return errUserNotFound
	}

	return nil
}
