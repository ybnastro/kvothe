package interfaces

//go:generate mockgen --destination=../mocks/mock_irepository.go --package=mocks --source=irepository.go

import "github.com/SurgicalSteel/kvothe/models"

type InterfaceKvotheRepository interface {
	GetSongQuoteByIDPostgres(id int64) (*models.SongQuote, error)

	GetAllSongQuotesPostgres() ([]models.SongQuote, error)
	UpsertSongQuoteRedis(sq models.SongQuote) error
}
