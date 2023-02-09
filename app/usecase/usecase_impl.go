package usecase

import (
	"context"
	"log"
	mongomodel "proto_example/app/model/mongo"
	"proto_example/app/model/pb"
	"proto_example/app/repository"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Usecase struct {
	pb.UnimplementedPostServiceServer
	repo repository.Repo
}

func NewUsecase(repo repository.Repo) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (res *pb.PostResponse, err error) {
	log.Println(req.GetTitle(), "title")
	post := &mongomodel.CreatePostRequest{
		Title: req.GetTitle(),
		Content: req.GetContent(),
		Image: req.GetImage(),
		User: req.GetUser(),
	}

	newPost, err := u.repo.CreatePost(post)
	if err != nil {
		log.Fatal(err)
	}

	res = &pb.PostResponse{
		Post: &pb.Post{
			Id: newPost.Id.Hex(),
			Title: newPost.Title,
			Content: newPost.Content,
			User: newPost.User,
			CreatedAt: timestamppb.New(newPost.CreateAt),
			UpdatedAt: timestamppb.New(newPost.UpdatedAt),
		},
	}
	
	return res, nil
} 

func (u *Usecase) GetPost(ctx context.Context, req *pb.PostRequest) (res *pb.PostResponse, err error) {
	
	post, err := u.repo.GetPost(req.GetId())
	if err != nil {
		log.Fatal(err)
	}
	res = &pb.PostResponse{
		Post: &pb.Post{
			Id: post.Id.Hex(),
			Title: post.Title,
			Content: post.Content,
			User: post.User,
			CreatedAt: timestamppb.New(post.CreateAt),
			UpdatedAt: timestamppb.New(post.UpdatedAt),
		},
	}
	
	return res, nil
}

func (u *Usecase) GetPosts(req *pb.GetPostsRequest, stream pb.PostService_GetPostsServer) (err error) {
	
	page := req.GetPage()
	limit := req.GetLimit()

	posts, err := u.repo.GetAllPost(int(page), int(limit))
	if err != nil {
		log.Fatal(err)
	}
	
	for _, post := range posts {
		stream.Send(&pb.Post{
			Id: post.Id.Hex(),
			Title: post.Title,
			Content: post.Content,
			User: post.User,
			CreatedAt: timestamppb.New(post.CreateAt),	
			UpdatedAt: timestamppb.New(post.UpdatedAt),		
		})
	}
	
	return nil
}

func (u *Usecase) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (res *pb.PostResponse, err error) {
	id := req.GetId()
	post := &mongomodel.UpdatePost{
		Title: req.GetTitle(),
		Content: req.GetContent(),
		Image: req.GetImage(),
		User: req.GetUser(),
	}
	updatePost, err := u.repo.UpdatePost(id, post)
	if err != nil {
		log.Fatal(err)
	}

	res = &pb.PostResponse{
		Post: &pb.Post{
			Id: updatePost.Id.Hex(),
			Title: updatePost.Title,
			Content: updatePost.Content,
			Image: updatePost.Image,
			User: updatePost.User,
			CreatedAt: timestamppb.New(updatePost.CreateAt),
			UpdatedAt: timestamppb.New(updatePost.UpdatedAt),
		},
	}
	return res, nil
}

func (u *Usecase) DeletePost(ctx context.Context, req *pb.PostRequest) (res *pb.DeletePostResponse, err error) {
	id := req.GetId()
	err = u.repo.DeletePost(id)
	if err != nil {
		log.Fatal(err)
	}
	res = &pb.DeletePostResponse{
		Success: true,
	}
	return
}