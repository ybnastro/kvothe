package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/SurgicalSteel/kvothe/mocks"
	"github.com/SurgicalSteel/kvothe/models"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func TestControllers_GetSongQuoteByIDHandler(t *testing.T) {
	songQuoteSample := &models.SongQuote{
		ID:         1,
		BandName:   "Incubus",
		SongTitle:  "Warning",
		AlbumTitle: "Morning View",
		AlbumYear:  2002,
		QuoteText:  "Don't ever let life pass you by",
	}

	testCases := []struct {
		name               string
		initRouter         func() (*gin.Context, *httptest.ResponseRecorder)
		doMockService      func(mock *mocks.MockIService)
		expectedHTTPStatus int
	}{
		{
			name: "error : id in URL is not defined",
			initRouter: func() (*gin.Context, *httptest.ResponseRecorder) {
				h := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(h)
				// c.Params = make(gin.Params, 1)
				// c.Params[0].Key = "id"
				// c.Params[0].Value = "-19"
				c.Request, _ = http.NewRequest(http.MethodGet, "localhost:8082/api/quote", nil)
				return c, h
			},
			doMockService:      func(mock *mocks.MockIService) {},
			expectedHTTPStatus: http.StatusBadRequest,
		},
		{
			name: "error : SongQuote is not found",
			initRouter: func() (*gin.Context, *httptest.ResponseRecorder) {
				h := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(h)
				c.Params = make(gin.Params, 1)
				c.Params[0].Key = "id"
				c.Params[0].Value = "6661"
				c.Request, _ = http.NewRequest(http.MethodGet, "localhost:8082/api/quote/6661", nil)
				return c, h
			},
			doMockService: func(mock *mocks.MockIService) {
				mock.EXPECT().GetSongQuoteByID(int64(6661)).Return(nil, http.StatusNotFound, errors.New("Song Quote with ID 6661 is not found"))
			},
			expectedHTTPStatus: http.StatusNotFound,
		},
		{
			name: "success : SongQuote is found",
			initRouter: func() (*gin.Context, *httptest.ResponseRecorder) {
				h := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(h)
				c.Params = make(gin.Params, 1)
				c.Params[0].Key = "id"
				c.Params[0].Value = "1"
				c.Request, _ = http.NewRequest(http.MethodGet, "localhost:8082/api/quote/1", nil)
				return c, h
			},
			doMockService: func(mock *mocks.MockIService) {
				mock.EXPECT().GetSongQuoteByID(int64(1)).Return(songQuoteSample, http.StatusOK, nil)
			},
			expectedHTTPStatus: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockIService(mockCtrl)

			kc := &KvotheController{
				Services: mockService,
			}

			tc.doMockService(mockService)

			ginContext, _ := tc.initRouter()
			kc.GetSongQuoteByIDHandler(ginContext)

			if !reflect.DeepEqual(ginContext.Writer.Status(), tc.expectedHTTPStatus) {
				t.Errorf("KvotheController.GetSongQuoteByIDHandler got = %v, want %v", ginContext.Writer.Status(), tc.expectedHTTPStatus)
			}
		})
	}
}
