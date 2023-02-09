package repository

import mongomodel "proto_example/app/model/mongo"

type Repo interface {
	CreatePost(in *mongomodel.CreatePostRequest) (out *mongomodel.DBPost, err error)
	GetPost(id string) (out *mongomodel.DBPost, err error)
	GetAllPost(page, limit int) (out []*mongomodel.DBPost, err error)
	UpdatePost(id string, in *mongomodel.UpdatePost) (out *mongomodel.DBPost, err error)
	DeletePost(id string) (err error)
}
