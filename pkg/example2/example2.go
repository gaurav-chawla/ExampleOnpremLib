package example2

import (
	"context"
	"github.com/srohatgi/graph"
	"log"
)

type Example2 struct {
	anotherInternalProperty string

	graph.Depends     `yaml:",inline"` // required
	ExampleDependency string           `yaml:"exampleDependency,omitempty"`
}

func (e *Example2) WithSomethingElse(s string) *Example2 {
	e.anotherInternalProperty = s
	return e
}

func (e *Example2) Update(ctx context.Context) (string, error) {

	log.Printf("Inside Example2 Update. Name: %s, ExampleDependency: %s, internalProperty: %s", e.Name, e.ExampleDependency, e.anotherInternalProperty)
	return "", nil
}

func (e *Example2) Delete(ctx context.Context) error {
	log.Println("Inside Example2 delete")
	return nil
}
