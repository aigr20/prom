package testdata

import (
	"aigr20/prom/models"
	"time"
)

var SampleSprints = []models.Sprint{
	{
		ID:       1,
		Name:     "Sprint 1",
		Start:    time.Date(2023, time.June, 24, 0, 0, 0, 0, time.Local),
		End:      time.Date(2023, time.June, 30, 0, 0, 0, 0, time.Local),
		Finished: true,
		Current:  false,
	},
	{
		ID:       2,
		Name:     "Sprint 2",
		Start:    time.Date(2023, time.July, 1, 0, 0, 0, 0, time.Local),
		End:      time.Date(2023, time.July, 8, 0, 0, 0, 0, time.Local),
		Finished: false,
		Current:  true,
	},
	{
		ID:       3,
		Name:     "Sprint 1",
		Start:    time.Date(2023, time.June, 24, 0, 0, 0, 0, time.Local),
		End:      time.Date(2023, time.June, 28, 0, 0, 0, 0, time.Local),
		Finished: false,
		Current:  true,
	},
}
