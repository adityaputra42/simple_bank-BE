package service

import (
	"fmt"
	"simple_bank_solid/api/repository"
	"simple_bank_solid/config"
	"simple_bank_solid/model/web/request"
	"simple_bank_solid/model/web/response"
	"simple_bank_solid/token"
	"time"
)

type SessionService interface {
	RenewAccessToken(req request.RenewAccessTokenRequest) (response.RenewAccessTokenResponse, error)
}

type SessionServiceImpl struct {
	sessionRepo repository.SessionRepository
	tokenMaker  token.Maker
	config      config.Configuration
}

// RenewAccessToken implements SessionService.
func (s *SessionServiceImpl) RenewAccessToken(req request.RenewAccessTokenRequest) (response.RenewAccessTokenResponse, error) {

	refreshPayload, err := s.tokenMaker.VerifyToken(req.RefreshToken)

	if err != nil {
		return response.RenewAccessTokenResponse{}, err
	}

	session, err := s.sessionRepo.FindById(refreshPayload.ID.String())

	if err != nil {
		return response.RenewAccessTokenResponse{}, err
	}

	if session.IsBlocked {
		err := fmt.Errorf("Session is blocked")
		return response.RenewAccessTokenResponse{}, err
	}

	if session.User.ID != refreshPayload.UserId {
		err := fmt.Errorf("Incorrect Session user")
		return response.RenewAccessTokenResponse{}, err
	}

	if session.RefreshToken != req.RefreshToken {
		err := fmt.Errorf("miss match Session token")
		return response.RenewAccessTokenResponse{}, err
	}

	if time.Now().After(session.ExpiredAt) {
		err := fmt.Errorf("Expired Session")
		return response.RenewAccessTokenResponse{}, err
	}

	accessToken, accessPayload, err := s.tokenMaker.CreateToken(session.User.Username, session.UserId, s.config.AccessTokenDuration)

	if err != nil {
		return response.RenewAccessTokenResponse{}, err
	}

	return response.RenewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiredAt: accessPayload.ExpiredAt,
	}, nil
}

func NewSessionService(SessionRepo repository.SessionRepository) SessionService {

	maker := token.GetTokenMaker()
	conf := config.GetCofig()
	return &SessionServiceImpl{
		sessionRepo: SessionRepo,
		config:      conf,
		tokenMaker:  maker,
	}
}
