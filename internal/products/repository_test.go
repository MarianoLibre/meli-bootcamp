package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)



type StubStore struct{}

func (s StubStore) Read(data interface{}) error {
	a, ok := data.(*[]Product)
	if !ok {
		return errors.New("it failed!")
	}
	*a = append(*a, Product{
			Id: 1,
			Name: "celu",
			Colour: "black",
			Price: 200.5,
			Stock: 12,
			Code: "XXX",
			Published: true,
			CreatedAt: "today",
		})
	*a = append(*a, Product{
			Id: 2,
			Name: "celu",
			Colour: "red",
			Price: 500.5,
			Stock: 12,
			Code: "ZZZ",
			Published: true,
			CreatedAt: "today",
		})
	return nil
}

func (s StubStore) Write(data interface{}) error{
	return nil
}


func TestGetAll(t *testing.T) {
	var db StubStore
	r := NewRepository(db)

	out, err := r.GetAll()
	expected := []Product{
		{
			Id: 1,
			Name: "celu",
			Colour: "black",
			Price: 200.5,
			Stock: 12,
			Code: "XXX",
			Published: true,
			CreatedAt: "today",
		},
		{
			Id: 2,
			Name: "celu",
			Colour: "red",
			Price: 500.5,
			Stock: 12,
			Code: "ZZZ",
			Published: true,
			CreatedAt: "today",
		},
	}

	assert.Nil(t, err, "")
	assert.Equal(t, expected, out, "")

}
