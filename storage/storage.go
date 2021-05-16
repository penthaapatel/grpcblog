package storage

import (
	"fmt"
	"grpcblog/blog"
	"grpcblog/serializer"
	"log"
	"sync"
)

//BlogStorage is an interface to store blog data
type BlogStorage interface {
	Save(blog *blog.Blog, id string) error
	View()
}

//InMemoryBlogStorage stores blog data in memory using a map where key is blog post UUID and the value is a blog struct
type InMemoryBlogStorage struct {
	mutex sync.RWMutex
	blogs map[string]*blog.Blog
}

//NewInMemoryBlogStorage returns a new InMemoryBlogStorage
func NewInMemoryBlogStorage() *InMemoryBlogStorage {
	return &InMemoryBlogStorage{
		blogs: make(map[string]*blog.Blog),
	}
}

//Save saves the blog posts to the memory
func (storage *InMemoryBlogStorage) Save(b *blog.Blog, id string) error {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()

	storage.blogs[id] = b

	if storage.blogs[id] == nil {
		return fmt.Errorf("could not save new blog post in memory")
	}
	err := serializer.WriteProtobufToJSONFile(b, "blogs.json")
	if err != nil {
		return fmt.Errorf("error writing to json file: %w", err)
	}

	return nil
}

func (storage *InMemoryBlogStorage) View() {
	for id, blog := range storage.blogs {
		log.Printf("ID : %s ", id)
		log.Printf("Blog Post: %s \n", blog)
	}
}
