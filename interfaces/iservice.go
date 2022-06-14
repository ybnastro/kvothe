package interfaces

import "github.com/SurgicalSteel/kvothe/models"

type InterfaceKvotheService interface {
	GetSongQuoteByID(id int64) (*models.SongQuote, int, error)
	GetAllSongData() ([]models.SongQuote, int, error)
	BackfillRedis() error
}
