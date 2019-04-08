package example

import (
	"context"
        "fmt"
	"github.com/srohatgi/graph"
)

type Example struct {
	internalProperty string

	graph.Depends `yaml:",inline"` // required
	Image         string           `yaml:"image,omitempty"`
	Port          int32            `yaml:"port,omitempty"`

	// Output
	ExampleReturn string `yaml:"exampleReturn,omitempty"`
}

func (e *Example) WithSomething(s string) *Example {
	e.internalProperty = s
	return e
}

func (e *Example) Update(ctx context.Context) (string, error) {

        fmt.Printf("inside update. Name: %s, image: %s, port: %d, internalProperty: %s", e.Name, e.Image, e.Port, e.internalProperty)
	e.ExampleReturn = "SampleReturn" // optional - only in case of dependency

	return "", nil
}

func (e *Example) Delete(ctx context.Context) error {
	
        fmt.Println("inside delete")
        return nil
}
