package repositories

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/SurgicalSteel/kvothe/interfaces"
	"github.com/SurgicalSteel/kvothe/models"
	"github.com/SurgicalSteel/kvothe/resources"
)

type KvotheRepository struct {
	DB      interfaces.IDatabase
	Redis   interfaces.IRedis
	HTTP    interfaces.IHTTP
	Conf    *resources.AppConfig
	Context interfaces.IContext
}

func (kr *KvotheRepository) GetSongQuoteByIDPostgres(id int64) (*models.SongQuote, error) {
	var result models.SongQuote
	err := kr.DB.Get(&result, "SELECT id, quote_text, song_title, album_title, album_year, band_name FROM song_quotes WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (kr *KvotheRepository) GetAllSongQuotesPostgres() ([]models.SongQuote, error) {
	var result []models.SongQuote
	err := kr.DB.Select(&result, "SELECT id, quote_text, song_title, album_title, album_year, band_name FROM song_quotes ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (kr *KvotheRepository) UpsertSongQuoteRedis(sq models.SongQuote) error {
	raw, err := json.Marshal(sq)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("song_quotes:%d", sq.ID)

	err = kr.Redis.Set(key, string(raw))
	return err
}

//GetSongQuoteByIDRedis
func (kr *KvotheRepository) GetSongQuoteByIDRedis(id int64) (*models.SongQuote, error) {
	key := fmt.Sprintf("song_quotes:%d", id)
	rawValue, err := kr.Redis.Get(key)
	if err != nil {
		log.Printf("Err Redis Get %s\n", err.Error())
		return nil, err
	}

	var songQuote models.SongQuote
	err = json.Unmarshal([]byte(rawValue), songQuote)
	if err != nil {
		log.Printf("Err Unmarshal %s\n", err.Error())
		return nil, err
	}

	return &songQuote, nil
}
