package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/krarjun90/sample-go-api/models"
)

type TrackRepository interface {
	GetAllTracks() ([]models.Track, error)
	GetTrackById(int32) (models.Track, error)
	AddTrack(*models.Track) (*models.Track, error)
	DeleteTrackById(int32) error
}

type trackRepository struct {
	db *sqlx.DB
}

func NewTrackRepository(db *sqlx.DB) TrackRepository {
	return &trackRepository{
		db: db,
	}
}

func (t *trackRepository) GetAllTracks() ([]models.Track, error) {
	var tracks []models.Track
	err := t.db.Select(&tracks, "select id, title, singer, album, duration from tracks")
	return tracks, err
}

func (t *trackRepository) GetTrackById(id int32) (models.Track, error) {
	var track models.Track
	err := t.db.Get(&track, "select id, title, singer, album, duration from tracks where id = $1", id)
	return track, err
}

func (t *trackRepository) AddTrack(track *models.Track) (*models.Track, error) {
	nstmt, err := t.db.PrepareNamed("insert into tracks(title, album, singer, duration)" +
		"values(:title, :album, :singer, :duration) returning id")
	if err != nil {
		return nil, err
	}
	err = nstmt.Get(track, track)
	return track, err
}

func (t *trackRepository) DeleteTrackById(id int32) error {
	_, err := t.db.Exec("delete from tracks where id = $1", id)
	return err
}
