package usecase

import (
	"fmt"

	"github.com/LavaJover/shvark-arbitrage-bot/internal/domain"
	"github.com/LavaJover/shvark-arbitrage-bot/internal/grpcapi"
)

type AuthUsecase struct {
	authRepo domain.AuthRepository
	ssoClient *grpcapi.SSOClient
	authzClient *grpcapi.AuthzClient
}

func NewAuthUsecase(
	repo domain.AuthRepository, 
	ssoClient *grpcapi.SSOClient,
	authzClient *grpcapi.AuthzClient,
	) *AuthUsecase {
	return &AuthUsecase{
		authRepo: repo, 
		ssoClient: ssoClient,
		authzClient: authzClient,
	}
}

func (uc *AuthUsecase) Authorize(telegramID int64, token string) (string, error) {
	// check if given token is valid
	valid, traderID, err := uc.ssoClient.ValidateToken(token)
	if err != nil || !valid{
		return "", fmt.Errorf("invalid token")
	}
	// // check user permission
	// allowed, err := uc.authzClient.CheckPermission(traderID, "*", "*")
	// if err != nil {
	// 	return "", err
	// }
	// if !allowed {
	// 	return "", status.Error(codes.PermissionDenied, "not enough rights")
	// }
	// if user got permission => we can map telegramID and traderID
	err = uc.authRepo.SaveMapping(telegramID, traderID)
	if err != nil {
		return "", err
	}
	return traderID, nil
}

func (uc *AuthUsecase) GetTraderIDByTelegramID(telegramID int64) (string, error) {
	return uc.authRepo.GetTraderID(telegramID)
}

func (uc *AuthUsecase) GetTelegramIDsByTraderID(traderID string) ([]int64, error) {
	return uc.authRepo.GetTelegramIDsByTraderID(traderID)
}

func (uc *AuthUsecase) GetTelegramIDs() ([]int64, error) {
	return uc.authRepo.GetTelegramIDs()
}