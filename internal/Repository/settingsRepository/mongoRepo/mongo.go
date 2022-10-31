package mongoRepo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sls/internal/Repository/settingsRepository"
	"sls/internal/entity/adminEntity"
	"time"
)

type mongoRepository struct {
	client *mongo.Client
}

func (m *mongoRepository) FetchSettings() (*adminEntity.AdminSettings, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*20)
	defer cancelFunc()

	var settings adminEntity.AdminSettings
	filter := bson.D{{}}
	coll := m.client.Database("sls").Collection("admin_settings")
	err := coll.FindOne(ctx, filter).Decode(&settings)
	if err != nil {
		return nil, err
	}
	return &settings, nil
}

func (m *mongoRepository) CreateSettings(settings *adminEntity.AdminSettings) (*adminEntity.AdminSettings, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*20)
	defer cancelFunc()

	coll := m.client.Database("sls").Collection("admin_settings")
	_, err := coll.InsertOne(ctx, settings)
	if err != nil {
		return nil, err
	}

	return settings, nil
}

func (m *mongoRepository) EditSettings(id string, settings *adminEntity.AdminSettings) (*adminEntity.AdminSettings, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*20)
	defer cancelFunc()

	filter := bson.D{{"_id", id}}
	update := bson.D{

		{
			"$set", bson.D{
				{"name", settings.Name},
				{"number_of_bikes", settings.NumberOfBikes},
				{"assigning_order", settings.AssigningOrder},
				{"bike_availability", settings.BikeAvailability},
				{"delivery_pricing_rate", settings.DeliveryPricingRate},
				{"order_rejection_reasons", settings.OrderRejectionReasons},
				{"business_category", settings.BusinessCategory},
				{"ticket_category", settings.TicketCategory},
				{"expenses", settings.Expenses},
			},
		},
	}
	coll := m.client.Database("sls").Collection("admin_settings")
	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return settings, err
}

func NewMongoRepository(client *mongo.Client) settingsRepository.SettingsRepo {
	return &mongoRepository{client: client}
}
