// To run the tests : go test ./api/books -v

package api

import (
	"testing"

	"gotest.tools/assert"
	//is "gotest.tools/assert/cmp"
)

const bookMockedJSON string = `{"title":"Cloud Native Go","author":"Omer Yesil","isbn":"123123123","description":"Desc1"}`

func TestBookToJSON(t *testing.T) {
	book := Book{
		Title:       "Cloud Native Go",
		Author:      "Omer Yesil",
		ISBN:        "123123123",
		Description: "Desc1",
	}

	json := book.ToJSON()

	assert.Equal(t,
		bookMockedJSON,
		string(json))

}

func TestBookFromJSON(t *testing.T) {

	book := fromJSON([]byte(bookMockedJSON))

	assert.Equal(t, book.Author, "Omer Yesil")
}
