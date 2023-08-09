package testdata

import (
	"aigr20/prom/models"
	"time"
)

var SampleProjects = []models.Project{
	{
		ID:      1,
		Name:    "Glada Schemat",
		Created: time.Date(2023, time.June, 24, 10, 0, 0, 0, time.Local),
		Updated: time.Date(2023, time.June, 24, 10, 15, 0, 0, time.Local),
	},
	{
		ID:      2,
		Name:    "Max Power",
		Created: time.Date(2023, time.June, 24, 9, 10, 0, 0, time.Local),
		Updated: time.Date(2023, time.June, 24, 16, 13, 12, 0, time.Local),
	},
	{
		ID:      3,
		Name:    "Mina Recept",
		Created: time.Date(2023, time.June, 25, 15, 22, 35, 0, time.Local),
		Updated: time.Date(2023, time.June, 26, 12, 0, 14, 0, time.Local),
	},
}
