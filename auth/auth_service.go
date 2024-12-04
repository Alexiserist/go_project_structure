package auth


type AuthService interface {
	GenerateToken(username string) (string,error);
	EncodingPassword(password string) (string,error);
	CompareHashAndPassword(hash,password string) (bool);
	LoginHandler(username,password string) (UserData,error);
	ValidateToken(token string)(error);
	ValidateRefreshToken(token string)(error);
	ExchangeRefreshToken(refreshToken string)(string,error);
}

type authService struct {
	authRepository AuthRepository
}

func NewAuthService(repo AuthRepository) AuthService{
	return &authService {
		authRepository: repo,
	}
}

func (s *authService) GenerateToken(username string) (string, error) {
	return s.authRepository.GenerateToken(username)
}

func (s *authService) EncodingPassword(password string) (string,error){
	return s.authRepository.EncodingPassword(password)
}

func (s *authService) CompareHashAndPassword(hash,password string) (bool){
	return s.authRepository.CompareHashAndPassword(hash, password)
}

func (s *authService) LoginHandler(username,password string)(UserData,error){
	return s.authRepository.LoginHandler(username,password)
}

func (s *authService) ValidateToken(token string)(error){
	return s.authRepository.ValidateToken(token)
}

func (s *authService) ValidateRefreshToken(token string)(error){
	return s.authRepository.ValidateRefreshToken(token)
}

func (s *authService) ExchangeRefreshToken(refreshToken string)(string,error){
	return s.authRepository.ExchangeRefreshToken(refreshToken)
}