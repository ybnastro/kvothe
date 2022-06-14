//package services defines the app dependency and route mapping
package services

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/SurgicalSteel/kvothe/interfaces"
	"github.com/SurgicalSteel/kvothe/models"
	"github.com/SurgicalSteel/kvothe/resources"
)

type KvotheService struct {
	Repo    interfaces.InterfaceKvotheRepository
	HTTP    interfaces.IHTTP
	Conf    *resources.AppConfig
	Context interfaces.IContext
}

func (ks *KvotheService) GetSongQuoteByID(id int64) (*models.SongQuote, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.New("invalid ID")
	}

	songQuote, err := ks.Repo.GetSongQuoteByIDPostgres(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, http.StatusNotFound, errors.New(fmt.Sprintf("Song Quote with ID %d is not found", id))
		}
		return nil, http.StatusInternalServerError, err
	}

	return songQuote, http.StatusOK, nil
}

func (ks *KvotheService) GetAllSongData() ([]models.SongQuote, int, error) {
	songQuotes, err := ks.Repo.GetAllSongQuotesPostgres()
	if err != nil && err != sql.ErrNoRows {
		return songQuotes, http.StatusInternalServerError, err
	}

	return songQuotes, http.StatusOK, nil
}

func (ks *KvotheService) BackfillRedis() error {
	allSongQuotes, err := ks.Repo.GetAllSongQuotesPostgres()
	if err != nil {
		return err
	}

	for _, songQuote := range allSongQuotes {
		err = ks.Repo.UpsertSongQuoteRedis(songQuote)
		if err != nil {
			return err
		}
	}
	return nil
}
