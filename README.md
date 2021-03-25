# grpcblog

Implementation of a simple RPC service using Golang. The RPC service can create and save blog posts.
 - Client sends a request to the server to create a blog post with a given title and content.
 - The server saves the blog posts created to an in-memory storage.

Directory structure
```bash
grpcblog
├── Makefile
├── blog
│   ├── blog.pb.go
│   ├── blog.proto
│   └── blog_grpc.pb.go
├── client
│   └── client.go
├── go.mod
├── go.sum
├── server
│   └── server.go
├── storage
    └── storage.go
```
## Generating protocol buffer code

To clean up the generated proto files run:
```bash
make clean
```
To compile the proto file run:
```bash
make gen
```
This generates two files:
```bash
blog/blog_grpc.pb.go
blog/blog.pb.go
```
***blog_grpc.pb.go*** contains server and client stubs.

***blog.pb.go*** contains protocol buffer related code - responsible for binary serialization of data when it is transported between server and client.

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