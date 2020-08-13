package resolvers

import (
	"context"
	models1 "github.com/3dw1nM0535/Byte/db/models"
)

func (r *seasonResolver) ID(ctx context.Context, obj *models1.Season) (string, error) {
	id := obj.ID.String()
	return id, nil
}
