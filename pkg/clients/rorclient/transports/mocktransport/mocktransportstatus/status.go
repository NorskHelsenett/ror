package mocktransportstatus

type MockTransportStatus struct{}

func NewMockTransportStatus() *MockTransportStatus {
	return &MockTransportStatus{}
}

func (m *MockTransportStatus) IsEstablished() bool {
	return true
}

func (m *MockTransportStatus) GetApiVersion() string {
	return "v1"
}

func (m *MockTransportStatus) GetLibVersion() string {
	return "1.0.0-mock"
}
