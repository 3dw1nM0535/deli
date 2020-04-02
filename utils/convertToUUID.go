package utils

import (
	"github.com/gofrs/uuid"
)

// ParseUUID : parse string to UUID
func ParseUUID(s string) uuid.UUID {
	specialUUID := "00000000-0000-0000-0000-000000000000"
	_, err := uuid.FromString(s)
	if err != nil {
		return uuid.Must(uuid.FromString(specialUUID))
	}
	id := uuid.Must(uuid.FromString(s))
	return uuid.Must(uuid.FromString(id.String()))
}
