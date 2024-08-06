package tools

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	client *mongo.Client
}

func NewServer (c *mongo.Client) *Server {
	return &Server {
		client: c
	}
}

func (d *mockDB) setUpDatabase () error {
	return nil
}
