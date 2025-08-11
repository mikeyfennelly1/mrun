package init

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

// Interface
type DataStore interface {
	GetUser(id int) string
}

// Testify mock
type MockStore struct {
	mock.Mock
}

func (m *MockStore) GetUser(id int) string {
	args := m.Called(id)
	return args.String(0)
}

func TestGreetUser(t *testing.T) {
	mockStore := new(MockStore)
	mockStore.On("GetUser", 1).Return("Mocked User")

	got := GreetUser(mockStore, 1)
	want := "Hello Mocked User"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	mockStore.AssertExpectations(t)
}
