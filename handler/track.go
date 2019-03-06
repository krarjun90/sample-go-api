package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/krarjun90/sample-go-api/models"
	"github.com/krarjun90/sample-go-api/repository"
	"net/http"
	"strconv"
)

type TrackHandler struct {
	trackRepo repository.TrackRepository
}

func NewTrackHandler(trackRepo repository.TrackRepository) *TrackHandler {
	return &TrackHandler{
		trackRepo: trackRepo,
	}
}

func(t *TrackHandler) GetAllTracks(w http.ResponseWriter, r *http.Request) {
	allTracks, err := t.trackRepo.GetAllTracks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error: "  + err.Error()))
		return
	}

	b, _ := json.Marshal(allTracks)
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)
}

func(t *TrackHandler) GetTrackById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	trackId, err := strconv.ParseInt(params["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error: "  + err.Error()))
		return
	}

	track, err := t.trackRepo.GetTrackById(int32(trackId))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error: "  + err.Error()))
		return
	}

	b, _ := json.Marshal(track)
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)
}

func(t *TrackHandler) DeleteTrackById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	trackId, err := strconv.ParseInt(params["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error: "  + err.Error()))
		return
	}

	err = t.trackRepo.DeleteTrackById(int32(trackId))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error: "  + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func(t *TrackHandler) AddTrack(w http.ResponseWriter, r *http.Request) {
	newTrack := &models.Track{}
	err := json.NewDecoder(r.Body).Decode(newTrack)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error parsing request body: "  + err.Error()))
		return
	}

	_, err = t.trackRepo.AddTrack(newTrack)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error adding track, error: "  + err.Error()))
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
