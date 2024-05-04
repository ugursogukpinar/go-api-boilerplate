package services

import "fmt"

type MyService struct{}

func NewMyService() *MyService {
	return &MyService{}
}

func (s *MyService) Greeting(name string) string {
	return fmt.Sprintf("Hello %s\n", name)
}
