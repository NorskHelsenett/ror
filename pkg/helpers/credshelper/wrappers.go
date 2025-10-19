package credshelper

// SimpleCredHelper is a simplified interface for credential helpers that only need to provide credentials without renewal functionality.
type SimpleCredHelper interface {
	GetCredentials() (string, string) // Returns username and password
}

type SimpleWrapper struct {
	wrapped SimpleCredHelper
}

// WrapSimpleCredsHelper creates a new SimpleWrapper that wraps a SimpleCredHelper to conform to the CredHelper interface.
func WrapSimpleCredsHelper(wrapper SimpleCredHelper) *SimpleWrapper {
	return &SimpleWrapper{
		wrapped: wrapper,
	}
}

// GetUsername returns the username from the wrapped SimpleCredHelper.
func (s *SimpleWrapper) GetUsername() string {
	username, _ := s.GetCredentials()
	return username
}

// GetPassword returns the password from the wrapped SimpleCredHelper.
func (s *SimpleWrapper) GetPassword() string {
	_, password := s.GetCredentials()
	return password
}

// GetCredentials returns the username and password from the wrapped SimpleCredHelper.
func (s *SimpleWrapper) GetCredentials() (string, string) {
	return s.wrapped.GetCredentials()
}

type SimpleCredHelperWithRenew interface {
	GetCredentials() (string, string) // Returns username and password
	CheckAndRenew() bool              // Returns true if credentials were renewed
}

type SimpleWrapperWithRenew struct {
	wrapped SimpleCredHelperWithRenew
}

// WrapSimpleCredsHelper creates a new SimpleWrapper that wraps a SimpleCredHelper to conform to the CredHelper interface.
func WrapSimpleCredsHelperWithRenew(wrapper SimpleCredHelperWithRenew) *SimpleWrapperWithRenew {
	return &SimpleWrapperWithRenew{
		wrapped: wrapper,
	}
}

// GetUsername returns the username from the wrapped SimpleCredHelper.
func (s *SimpleWrapperWithRenew) GetUsername() string {
	username, _ := s.GetCredentials()
	return username
}

// GetPassword returns the password from the wrapped SimpleCredHelper.
func (s *SimpleWrapperWithRenew) GetPassword() string {
	_, password := s.GetCredentials()
	return password
}

// GetCredentials returns the username and password from the wrapped SimpleCredHelper.
func (s *SimpleWrapperWithRenew) GetCredentials() (string, string) {
	return s.wrapped.GetCredentials()
}

// CheckAndRenew calls the CheckAndRenew method on the wrapped SimpleCredHelperWithRenew.
func (s *SimpleWrapperWithRenew) CheckAndRenew() bool {
	return s.wrapped.CheckAndRenew()
}
