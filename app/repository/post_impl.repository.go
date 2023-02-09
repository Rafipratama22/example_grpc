package repository

import (
	"context"
	"errors"
	"log"
	mongomodel "proto_example/app/model/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repo struct {
	client *mongo.Client
}

func NewRepo(client *mongo.Client) Repo {
	return &repo{
		client: client,
	}
}

func (r *repo) CreatePost(in *mongomodel.CreatePostRequest) (out *mongomodel.DBPost, err error) {
	var newPost *mongomodel.DBPost
	ctx := context.Background()
	res, err := r.client.Database("store").Collection("posts").InsertOne(ctx, in)
	if err != nil {
		log.Fatal(err)
	}

	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"title": 1}, Options: opt}

	if _, err := r.client.Database("store").Collection("posts").Indexes().CreateOne(ctx, index); err != nil {
		return nil, errors.New("could not create index for title")
	}

	query := bson.M{"_id": res.InsertedID}

	if err := r.client.Database("store").Collection("posts").FindOne(ctx, query).Decode(&newPost); err != nil {
		return nil, errors.New("could not create index for title")
	}

	return newPost, nil
}

func (r *repo) GetPost(id string) (out *mongomodel.DBPost, err error) {
	var getPost *mongomodel.DBPost
	ctx := context.Background()
	obId, _ := primitive.ObjectIDFromHex(id)

	query := bson.M{"_id": obId}

	if err := r.client.Database("store").Collection("posts").FindOne(ctx, query).Decode(&getPost); err != nil {
		return nil, errors.New("could not create index for title")
	}
	return getPost, nil
}

func (r *repo) GetAllPost(page, limit int) (out []*mongomodel.DBPost, err error) {
	ctx := context.Background()
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 1
	}

	skip := (page - 1) * limit

	opt := options.FindOptions{}
	opt.SetLimit(int64(limit))
	opt.SetSkip(int64(skip))
	opt.SetSort(bson.M{"created_at": -1})

	query := bson.M{}

	cursor, err := r.client.Database("store").Collection("posts").Find(ctx, query, &opt)

	defer cursor.Close(ctx)

	var posts []*mongomodel.DBPost

	for cursor.Next(ctx) {
		post := &mongomodel.DBPost{}
		err := cursor.Decode(post)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(posts) == 0 {
		return []*mongomodel.DBPost{}, nil
	}

	return posts, nil
}

func (r *repo) UpdatePost(id string, in *mongomodel.UpdatePost) (out *mongomodel.DBPost, err error) {
	var updated *mongomodel.DBPost
	ctx := context.Background()

	bsId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: bsId}}
	update := bson.D{{Key: "$set", Value: in}}
	res := r.client.Database("store").Collection("posts").FindOneAndUpdate(ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	err = res.Decode(&updated)

	if err != nil {
		log.Fatal(err)
	}

	return updated, nil
}

func (r *repo) DeletePost(id string) (err error) {
	ctx := context.Background()
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := r.client.Database("store").Collection("posts").DeleteOne(ctx, query)
	if err != nil {
		log.Fatal(err)
	}

	if res.DeletedCount == 0 {
		return nil
	}

	return nil
}
