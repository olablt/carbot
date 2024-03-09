package main

import (
	"context"
	"fmt"
)

type NewNote struct {
	Note string `json:"note"`
}

type Note struct {
	ID   int    `json:"id"`
	Note string `json:"note"`
}

// Service is responsible for managing notes.
// All behaviour is hardcoded for demo purposes.
type Service struct{}

func (s *Service) CreateNote(ctx context.Context, n NewNote) (Note, error) {
	fmt.Printf("CreateNote called: %+v\n", n)
	return Note{
		ID:   1203,
		Note: n.Note,
	}, nil
}

func (s *Service) UpdateNote(ctx context.Context, n Note) (Note, error) {
	fmt.Printf("UpdateNote called: %+v\n", n)
	return n, nil
}
