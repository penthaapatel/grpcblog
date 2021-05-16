# grpcblog

[Here's is the link to my Medium article which explains the implementation of the code](https://penthaa.medium.com/how-to-implement-a-simple-grpc-service-using-golang-b9c58cab0929)

Implementation of a simple RPC service using Golang. The RPC service can create and save blog posts.
 - Client sends a request to the server to create a blog post with a given title and content.
 - The server saves the blog posts created to an in-memory storage.

Directory structure
```bash
grpcblog
├── Makefile
├── README.md
├── blog
│   ├── blog.pb.go
│   ├── blog.proto
│   └── blog_grpc.pb.go
├── blogs.json
├── client
│   └── client.go
├── client.png
├── go.mod
├── go.sum
├── serializer
│   └── json.go
├── server
│   └── server.go
├── server.png
└── storage
    └── storage.go
```
## Generating protocol buffer code

`make clean` : To clean up the generated proto files run:

`make gen` : To compile the proto file run:

This generates two files:
```bash
blog/blog_grpc.pb.go
blog/blog.pb.go
```
`blog_grpc.pb.go` : contains server and client stubs.

`blog.pb.go` : contains protocol buffer related code - responsible for binary serialization of data when it is transported between server and client.

## Running the server and client

Run the server and client code in two separate terminal windows:

```bash
go run server/server.go
```

![Server](server.png)

```bash
go run client/client.go
```
![Client](client.png)

The blog posts created are saved to a JSON file : `blogs.json`