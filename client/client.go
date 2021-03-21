package main

import (
	"bufio"
	"context"
	"fmt"
	"grpcblog/blog"
	"log"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := blog.NewBlogServiceClient(conn)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Blog Title: ")
	title, _ := reader.ReadString('\n')
	title = strings.Trim(title, "\n")
	fmt.Print("Enter Blog contents: ")
	body, _ := reader.ReadString('\n')
	body = strings.Trim(body, "\n")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.CreatePost(ctx, &blog.BlogRequest{
		Blog: &blog.Blog{
			Title: title,
			Body:  body,
		},
	})
	if err != nil {
		log.Fatalf("Could not create Blog Post :%v", err)
	}

	log.Printf("Post Successfully Created")
}
