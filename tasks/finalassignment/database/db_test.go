package database

import (
	"bytes"
	"encoding/json"
	"finalassignment/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertBookData(t *testing.T) {
	r := SetUpRouter()
	books := model.Library{
		Title:  "Unit Testing Book",
		Author: "Sachin",
	}

	jsonValue, _ := json.Marshal(books)

	req, _ := http.NewRequest("POST", "/addbooks", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

}
