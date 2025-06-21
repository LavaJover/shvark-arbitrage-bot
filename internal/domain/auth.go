package domain

type AuthRepository interface {
	SaveMapping(telegramID int64, traderID string) error
	GetTraderID(telegramID int64) (string, error)
	GetTelegramIDsByTraderID(traderID string) ([]int64, error)
	GetTelegramIDs() ([]int64, error)
}

type AuthUsecase interface {
	Authorize(telegramID int64, token string) (string, error)
	GetTraderIDByTelegramID(telegramID int64) (string, error)
	GetTelegramIDsByTraderID(traderID string) ([]int64, error)
	GetTelegramIDs() ([]int64, error)
}
