package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/krarjun90/sample-go-api/mocks"
	"github.com/krarjun90/sample-go-api/models"
	"github.com/stretchr/testify/assert"
)

func TestTrackHandlerAddTrack(t *testing.T) {
	req, err := http.NewRequest("GET", "/tracks", nil)
	if err != nil {
		t.Fatal(err)
	}

	expectedTracks := []models.Track{
		{
			Title: "title1",
		},
		{
			Title: "title2",
		},
	}

	trackRepo := &mocks.TrackRepository{}
	trackRepo.On("GetAllTracks").Return(expectedTracks, nil)

	h := NewTrackHandler(trackRepo)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.GetAllTracks)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	var actualTracks []models.Track
	err = json.NewDecoder(rr.Body).Decode(&actualTracks)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(actualTracks))
}
