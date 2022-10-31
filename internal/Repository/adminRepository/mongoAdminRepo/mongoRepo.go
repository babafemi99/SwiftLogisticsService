package mongoAdminRepo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"sls/internal/Repository/adminRepository"
	"sls/internal/entity/adminEntity"
	"time"
)

type mongoCl struct {
	client *mongo.Client
}

func (m *mongoCl) Fetch() ([]*adminEntity.Admin, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*20)
	defer cancelFunc()

	coll := m.client.Database("sls").Collection("admin_users")
	docs, err := coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var results []*adminEntity.Admin

	for docs.Next(ctx) {
		var result adminEntity.Admin

		err := docs.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, &result)
	}

	err = docs.Err()
	if err != nil {
		return nil, err
	}

	docs.Close(ctx)

	return results, nil
}

func (m *mongoCl) Persist(admin *adminEntity.Admin) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*20)
	defer cancelFunc()

	coll := m.client.Database("sls").Collection("admin_users")
	_, err := coll.InsertOne(ctx, admin)
	if err != nil {
		return err
	}

	return nil
}

func (m *mongoCl) FetchByEmail(email string) (*adminEntity.Admin, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*20)
	defer cancelFunc()

	var results adminEntity.Admin

	coll := m.client.Database("sls").Collection("admin_users")
	err := coll.FindOne(ctx, bson.D{{"email", email}}).Decode(&results)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no Results with that email")
		}
		return nil, err
	}

	return &results, nil
}

func (m *mongoCl) FetchById(id string) (*adminEntity.Admin, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*20)
	defer cancelFunc()

	var results adminEntity.Admin
	newId, _ := primitive.ObjectIDFromHex(id)
	log.Println(newId)

	coll := m.client.Database("sls").Collection("admin_users")
	err := coll.FindOne(ctx, bson.D{{"_id", id}}).Decode(&results)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no Results with that email")
		}
		return nil, err
	}

	return &results, nil
}

func (m *mongoCl) EditData(id string, admin *adminEntity.AdminAccess) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*20)
	defer cancelFunc()
	filter := bson.D{{"_id", id}}
	update := bson.D{

		{
			"$set", bson.D{
				{"first_name", admin.FirstName},
				{"last_name", admin.LastName},
				{"email", admin.Email},
				{"phone_number", admin.PhoneNumber},
				{"position", admin.Position},
				{"profile_picture", admin.ProfilePicture},
				{"updated_at", admin.UpdatedAt},
				{"role", admin.Role},
			},
		},
	}

	coll := m.client.Database("sls").Collection("admin_users")
	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (m *mongoCl) Delete(email string) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*20)
	defer cancelFunc()

	coll := m.client.Database("sls").Collection("admin_users")
	_, err := coll.DeleteOne(ctx, bson.D{{"email", email}})
	if err != nil {
		return err
	}

	return nil
}

func (m *mongoCl) EditPassword(email, password string) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*20)
	defer cancelFunc()

	filter := bson.D{{"email", email}}
	update := bson.D{
		{"$set", bson.D{{"password", password}}},
	}
	coll := m.client.Database("sls").Collection("admin_users")
	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

// PersistTodo Saves the data to the database
func (m *mongoCl) PersistTodo(todo *adminEntity.Todo) (*adminEntity.Todo, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*20)
	defer cancelFunc()

	coll := m.client.Database("sls").Collection("admin_todo")
	_, err := coll.InsertOne(ctx, todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (m *mongoCl) DeleteTodo(id string) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*20)
	defer cancelFunc()

	filter := bson.D{{"_id", id}}

	coll := m.client.Database("sls").Collection("admin_todo")

	_, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (m *mongoCl) FavoriteTodo(id string) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*20)
	defer cancelFunc()

	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"starred", true}}}}

	coll := m.client.Database("sls").Collection("admin_todo")
	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (m *mongoCl) DoTodo(id string) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*20)
	defer cancelFunc()

	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"status", "DONE"}}}}

	coll := m.client.Database("sls").Collection("admin_todo")
	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (m *mongoCl) FetchAll() ([]*adminEntity.Todo, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*20)
	defer cancelFunc()

	var results []*adminEntity.Todo
	coll := m.client.Database("sls").Collection("admin_todo")
	docs, err := coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer docs.Close(ctx)

	for docs.Next(ctx) {
		var result adminEntity.Todo
		err = docs.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, &result)
	}

	err = docs.Err()
	if err != nil {
		return nil, err
	}

	return results, nil

}

func (m *mongoCl) FindByTitle(title string) (*adminEntity.Todo, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*20)
	defer cancelFunc()

	var result adminEntity.Todo
	filter := bson.D{{"title", title}}

	coll := m.client.Database("sls").Collection("admin_todo")
	err := coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func NewMongoAdminRepository(client *mongo.Client) adminRepository.AdminInterface {
	return &mongoCl{client: client}
}
