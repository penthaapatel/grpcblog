package storage

import (
	"grpcblog/blog"
	"fmt"
	"log"
	"sync"
)

type BlogStorage interface {
	Save(blog *blog.Blog, id string) error
	View()
}

type InMemoryBlogStorage struct {
	mutex sync.RWMutex
	blogs map[string]*blog.Blog
}

func NewInMemoryBlogStorage() *InMemoryBlogStorage {
	return &InMemoryBlogStorage{
		blogs: make(map[string]*blog.Blog),
	}
}

func (storage *InMemoryBlogStorage) Save(b *blog.Blog, id string) error {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()

	storage.blogs[id] = b

	if storage.blogs[id] == nil {
		return fmt.Errorf("could not save new blog post in memory")
	}
	
	return nil
}

func (storage *InMemoryBlogStorage) View() {
	for id, blog := range storage.blogs {
		log.Printf("ID : %s ", id)
		log.Printf("Blog Post: %s \n", blog)
	}
}
