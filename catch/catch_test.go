package catch

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockErrorHandler struct {
	mock.Mock
}

func (m *MockErrorHandler) Handle(err error) {
	m.Called(err)
}

func TestCatch_Success(t *testing.T) {
	result := Catch(
		func() (int, error) {
			return 42, nil
		},
		func(err error) {
			assert.Fail(t, "Error handler should not be called on success")
		},
	)

	assert.Equal(t, 42, result, "Expected result does not match")
}

func TestCatch_Error(t *testing.T) {
	expectedErr := errors.New("test error")
	mockHandler := new(MockErrorHandler)
	mockHandler.On("Handle", expectedErr).Return()

	result := Catch(
		func() (int, error) {
			return 0, expectedErr
		},
		mockHandler.Handle,
	)

	assert.Equal(t, 0, result, "Expected result should be the zero value of the type")
	mockHandler.AssertCalled(t, "Handle", expectedErr)
}

func TestCatch_ErrorWithNilHandler(t *testing.T) {
	expectedErr := errors.New("test error")
	var wasCalled bool

	result := Catch(
		func() (int, error) {
			return 0, expectedErr
		},
		func(err error) {
			wasCalled = true
		},
	)

	assert.Equal(t, 0, result, "Expected result should be the zero value of the type")
	assert.True(t, wasCalled, "Error handler should have been called")
}
