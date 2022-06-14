package services

import (
	"database/sql"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/SurgicalSteel/kvothe/mocks"
	"github.com/SurgicalSteel/kvothe/models"
	"github.com/golang/mock/gomock"
)

func TestService_GetSongQuoteByID_ErrorSQLNotFound_ExpectErrorNotFound(t *testing.T) {
	t.Run("ErrorSQLNotFound_ExpectErrorNotFound", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		testCase := struct {
			name               string
			id                 int64
			doMockRepository   func(mock *mocks.MockIRepository)
			expectedResult     *models.SongQuote
			expectedHTTPStatus int
			expectedError      error
		}{
			name: "error from Repo Get By ID : sql error no rows",
			id:   6661,
			doMockRepository: func(mock *mocks.MockIRepository) {
				mock.EXPECT().GetSongQuoteByIDPostgres(int64(6661)).Return(nil, sql.ErrNoRows)
			},
			expectedResult:     nil,
			expectedHTTPStatus: http.StatusNotFound,
			expectedError:      errors.New("Song Quote with ID 6661 is not found"),
		}

		mockRepo := mocks.NewMockIRepository(mockCtrl)

		ks := &KvotheService{
			Repo: mockRepo,
		}
		testCase.doMockRepository(mockRepo)
		actualData, actualHTTPStatus, actualErr := ks.GetSongQuoteByID(testCase.id)
		if !reflect.DeepEqual(actualData, testCase.expectedResult) {
			t.Fatalf("Unexpected songQuote result found.\nExpected %v\nWant %v", testCase.expectedResult, actualData)
		}
		if !reflect.DeepEqual(actualErr, testCase.expectedError) {
			t.Fatalf("Unexpected error result found.\nExpected %v\nWant %v", testCase.expectedError, actualErr)
		}
		if actualHTTPStatus != testCase.expectedHTTPStatus {
			t.Fatalf("Unexpected HTTP status found.\nExpected %v\nWant %v", testCase.expectedHTTPStatus, actualHTTPStatus)
		}
	})

}

func TestService_GetSongQuoteByID(t *testing.T) {
	songQuoteSample := &models.SongQuote{
		ID:         123,
		BandName:   "Nightwish",
		SongTitle:  "Gethsemane",
		AlbumTitle: "Oceanborn",
		AlbumYear:  1998,
		QuoteText:  "Without you, the poetry within me is dead",
	}

	testCases := []struct {
		name               string
		id                 int64
		doMockRepository   func(mock *mocks.MockIRepository)
		expectedResult     *models.SongQuote
		expectedHTTPStatus int
		expectedError      error
	}{
		{
			name:               "error negative ID expect error invalid ID",
			id:                 -666,
			doMockRepository:   func(mock *mocks.MockIRepository) {},
			expectedResult:     nil,
			expectedHTTPStatus: http.StatusBadRequest,
			expectedError:      errors.New("invalid ID"),
		},
		{
			name:               "error zero ID",
			id:                 0,
			doMockRepository:   func(mock *mocks.MockIRepository) {},
			expectedResult:     nil,
			expectedHTTPStatus: http.StatusBadRequest,
			expectedError:      errors.New("invalid ID"),
		},
		{
			name: "error from Repo Get By ID : general error",
			id:   666,
			doMockRepository: func(mock *mocks.MockIRepository) {
				mock.EXPECT().GetSongQuoteByIDPostgres(int64(666)).Return(nil, errors.New("error cuy"))
			},
			expectedResult:     nil,
			expectedHTTPStatus: http.StatusInternalServerError,
			expectedError:      errors.New("error cuy"),
		},
		{
			name: "error from Repo Get By ID : sql error no rows",
			id:   6661,
			doMockRepository: func(mock *mocks.MockIRepository) {
				mock.EXPECT().GetSongQuoteByIDPostgres(int64(6661)).Return(nil, sql.ErrNoRows)
			},
			expectedResult:     nil,
			expectedHTTPStatus: http.StatusNotFound,
			expectedError:      errors.New("Song Quote with ID 6661 is not found"),
		},
		{
			name: "success flow (no error)",
			id:   7,
			doMockRepository: func(mock *mocks.MockIRepository) {
				mock.EXPECT().GetSongQuoteByIDPostgres(int64(7)).Return(songQuoteSample, nil)
			},
			expectedResult:     songQuoteSample,
			expectedHTTPStatus: http.StatusOK,
			expectedError:      nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			//initialize mock controller
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			//initialize mockedRepo (from mock interface)
			mockedRepo := mocks.NewMockIRepository(mockCtrl)

			//inject mockedRepo into a new testing service
			ks := &KvotheService{
				Repo: mockedRepo,
			}

			//run doMockDatabase (to make sure if the actual function calls the method from mockedDB)
			tc.doMockRepository(mockedRepo)
			actualResult, actualHTTPStatus, actualError := ks.GetSongQuoteByID(tc.id)
			if !reflect.DeepEqual(actualResult, tc.expectedResult) {
				t.Fatalf("Unexpected songQuote result found.\nExpected %v\nWant %v", tc.expectedResult, actualResult)
			}
			if !reflect.DeepEqual(actualError, tc.expectedError) {
				t.Fatalf("Unexpected error result found.\nExpected %v\nWant %v", tc.expectedError, actualError)
			}
			if actualHTTPStatus != tc.expectedHTTPStatus {
				t.Fatalf("Unexpected HTTP status found.\nExpected %v\nWant %v", tc.expectedHTTPStatus, actualHTTPStatus)
			}
		})
	}
}
