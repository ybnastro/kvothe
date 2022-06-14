package repositories

import (
	"errors"
	"reflect"
	"testing"

	"github.com/SurgicalSteel/kvothe/mocks"
	"github.com/SurgicalSteel/kvothe/models"
	"github.com/golang/mock/gomock"
)

func TestRepositories_GetSongQuoteByIDPostgres(t *testing.T) {
	type testCase struct {
		name           string                          //title of the test case
		id             int64                           //parameter for the function
		doMockDatabase func(mock *mocks.MockIDatabase) //mock function
		expectedResult *models.SongQuote               //expected success result
		expectedError  error                           //expected error result
	}

	songQuoteSample := &models.SongQuote{
		// ID:         123,
		// BandName:   "Nightwish",
		// SongTitle:  "Gethsemane",
		// AlbumTitle: "Oceanborn",
		// AlbumYear:  1998,
		// QuoteText:  "Without you, the poetry within me is dead",
	}

	testCases := []testCase{
		{
			name: "error get from DB",
			id:   666,
			doMockDatabase: func(mock *mocks.MockIDatabase) {
				mock.
					EXPECT().
					Get(gomock.Any(), "SELECT id, quote_text, song_title, album_title, album_year, band_name FROM song_quotes WHERE id = $1", int64(666)).
					Return(errors.New("error cuy"))
			},
			expectedResult: nil,
			expectedError:  errors.New("error cuy"),
		},
		{
			name: "success get from DB",
			id:   123,
			doMockDatabase: func(mock *mocks.MockIDatabase) {
				mock.
					EXPECT().
					Get(songQuoteSample, "SELECT id, quote_text, song_title, album_title, album_year, band_name FROM song_quotes WHERE id = $1", int64(123)).
					Return(nil)
			},
			expectedResult: songQuoteSample,
			expectedError:  nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			//initialize mock controller
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			//initialize mockedDB (from mock interface)
			mockedDB := mocks.NewMockIDatabase(mockCtrl)

			//inject mockedDB into a new testing repository
			kr := &KvotheRepository{
				DB: mockedDB,
			}

			//run doMockDatabase (to make sure if the actual function calls the method from mockedDB)
			tc.doMockDatabase(mockedDB)

			//run the actual function and check the actual result against the expected result
			songQuote, err := kr.GetSongQuoteByIDPostgres(tc.id)
			if !reflect.DeepEqual(songQuote, tc.expectedResult) {
				t.Fatalf("Unexpected songQuote result found.\nExpected %v\nWant %v", tc.expectedResult, songQuote)
			}
			if !reflect.DeepEqual(err, tc.expectedError) {
				t.Fatalf("Unexpected error result found.\nExpected %v\nWant %v", tc.expectedError, err)
			}
		})
	}
}
