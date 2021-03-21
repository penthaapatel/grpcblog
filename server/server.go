package main

import (
	"grpcblog/blog"
	"grpcblog/storage"
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type BlogServer struct {
	blog.UnimplementedBlogServiceServer
	blogStorage storage.BlogStorage
}

func (s *BlogServer) CreatePost(ctx context.Context, req *blog.BlogRequest) (*blog.BlogResponse, error) {
	input := req.GetBlog()

	if input.Title == "" {
		return nil, errors.New("title cannot be empty")
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("cannot generate blog post ID: %w",err)
	}

	err = s.blogStorage.Save(input, id.String())
	if err != nil {
		res := &blog.BlogResponse{
			Created: false,
		}
		return res, fmt.Errorf("blog post could not be created: %w",err)
	}
	res := &blog.BlogResponse{
		Id:      id.String(),
		Created: true,
	}

	log.Printf("Blog post successfully created with ID : %s", id)
	//Uncomment this to view list of all saved posts until now
	//s.blogStorage.View()
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	bs := storage.NewInMemoryBlogStorage()

	blog.RegisterBlogServiceServer(s, &BlogServer{blogStorage: bs})

	log.Printf("Server started at port %s", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
