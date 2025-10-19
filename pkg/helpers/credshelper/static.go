package credshelper

type StaticCredsHelper struct {
	Username string
	Password string
}

func NewStaticCredsHelper(username, password string) *StaticCredsHelper {
	return &StaticCredsHelper{
		Username: username,
		Password: password,
	}
}

func (s *StaticCredsHelper) GetCredentials() (string, string) {
	return s.Username, s.Password
}

func (s *StaticCredsHelper) GetUsername() string {
	return s.Username
}

func (s *StaticCredsHelper) GetPassword() string {
	return s.Password
}

func (s *StaticCredsHelper) CheckAndRenew() bool {
	return true
}
