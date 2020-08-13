package proxy

import (
  "net/http"
  "net/http/httptest"
  "testing"
  "encoding/json"

  "github.com/stretchr/testify/mock"
  "github.com/stretchr/testify/assert"
)

type MockProviderService struct {
  mock.Mock
}

type MockClientService struct {
  mock.Mock
}

func (m *MockProviderService) GetAll() ([]*XRoadMember, error) {
  args := m.Called()
  return args.Get(0).([]*XRoadMember), args.Error(1)
}

func (m *MockClientService) GetAll() ([]*XRoadMember, error) {
  args := m.Called()
  return args.Get(0).([]*XRoadMember), args.Error(1)
}

func TestGetServiceProviders(t *testing.T) {
  var actualProviders []*XRoadMember
  expectedProviders := []*XRoadMember{
    &XRoadMember{ID: "CS:ORG:1111:Company1Provider", Name: "Company 1"},
    &XRoadMember{ID: "CS:ORG:1112:Company2Provider", Name: "Company 2"},
  }

  mockProviderService := new(MockProviderService)
  mockProviderService.On("GetAll").Return(expectedProviders, nil)

  ph := NewProviderHandler(mockProviderService)

  req := httptest.NewRequest("GET", "/service-providers", nil)
  w := httptest.NewRecorder()

  http.HandlerFunc(ph.GetServiceProviders).ServeHTTP(w, req)

  json.NewDecoder(w.Body).Decode(&actualProviders)

  assert.Equal(t, http.StatusOK, w.Code)
  assert.Equal(t, expectedProviders, actualProviders)
}

func TestGetServiceClients(t *testing.T) {
  var actualClients []*XRoadMember
  expectedClients := []*XRoadMember{
    &XRoadMember{ID: "CS:ORG:1111:Company1Provider", Name: "Company 1"},
    &XRoadMember{ID: "CS:ORG:1112:Company2Provider", Name: "Company 2"},
  }

  mockClientService := new(MockClientService)
  mockClientService.On("GetAll").Return(expectedClients, nil)

  ph := NewClientHandler(mockClientService)

  req := httptest.NewRequest("GET", "/service-clients", nil)
  w := httptest.NewRecorder()

  http.HandlerFunc(ph.GetServiceClients).ServeHTTP(w, req)

  json.NewDecoder(w.Body).Decode(&actualClients)

  assert.Equal(t, http.StatusOK, w.Code)
  assert.Equal(t, expectedClients, actualClients)
}