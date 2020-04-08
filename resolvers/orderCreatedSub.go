package resolvers

import (
	"context"

	"github.com/3dw1nM0535/deli/db/models"
	"github.com/gofrs/uuid"
)

var orderCreatedChannel map[string]map[string]chan *models.Order

func (r *subscriptionResolver) OrderCreated(ctx context.Context, id string) (<-chan *models.Order, error) {
	subID := uuid.Must(uuid.NewV4()).String()

	orderEvent := make(chan *models.Order, 1)

	go func() {
		<-ctx.Done()
		delete(orderCreatedChannel[id], subID)
	}()

	if orderCreatedChannel[id] == nil {
		orderCreatedChannel[id] = make(map[string]chan *models.Order, 1)
	}
	orderCreatedChannel[id][subID] = orderEvent
	return orderEvent, nil
}
