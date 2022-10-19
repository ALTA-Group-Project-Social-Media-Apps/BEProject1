package services

import (
	"errors"
	"testing"

	"github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDelete(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("success", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(nil).Once()
		srv := New(repo)
		input := uint(1)
		err := srv.Delete(input)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(errors.New("no data")).Once()
		srv := New(repo)
		input := uint(7)
		err := srv.Delete(input)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "no data", "error message doesn't match")
		repo.AssertExpectations(t)
	})
}
