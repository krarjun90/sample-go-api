package repository

import (
	"database/sql"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/krarjun90/sample-go-api/models"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestTrackRepositorySuite struct {
	suite.Suite
	db *sqlx.DB
}

func (s *TestTrackRepositorySuite) SetupTest() {
	s.db = sqlx.MustConnect("postgres", "user=postgres dbname=sample_go_api sslmode=disable")

	_, err := s.db.Exec("truncate table tracks;")
	if err != nil {
		panic("error dropping tracks, err: %s" + err.Error())
	}
}

func (s *TestTrackRepositorySuite) TestTrackRepository() {
	_ = pq.Einfo

	repo := NewTrackRepository(s.db)

	track1 := models.Track{
		Title:    "JaiHo",
		Album:    "Slumdog Millionaire",
		Singer:   "Sukhwinder Singh",
		Duration: "5:19",
	}

	track2 := models.Track{
		Title:    "Aaj Ki Raat",
		Album:    "Slumdog Millionaire",
		Singer:   "Sonu Nigam",
		Duration: "6:07",
	}

	_, err := repo.AddTrack(&track1)
	assert.Nil(s.T(), err)
	assert.True(s.T(), track1.ID > 0)

	_, err = repo.AddTrack(&track2)
	assert.Nil(s.T(), err)

	allTracks, err := repo.GetAllTracks()
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 2, len(allTracks))

	actualTrack, err := repo.GetTrackById(track1.ID)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), track1, actualTrack)

	err = repo.DeleteTrackById(track1.ID)
	assert.Nil(s.T(), err)

	_, err = repo.GetTrackById(track1.ID)
	assert.Equal(s.T(), sql.ErrNoRows, err)
}

func TestNewTrackRepository(t *testing.T) {
	suite.Run(t, &TestTrackRepositorySuite{})
}
