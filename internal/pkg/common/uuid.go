package common

import "github.com/google/uuid"

func UID() uuid.UUID {
	return uuid.Must(uuid.NewRandom())
}
