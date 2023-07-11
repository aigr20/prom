package database

import "testing"

var sampleStatuses = []struct {
	ID   int
	Name string
}{
	{
		ID:   1,
		Name: "TODO",
	},
	{
		ID:   2,
		Name: "In Progress",
	},
	{
		ID:   3,
		Name: "Finished",
	},
}

func getStatusRepository(t *testing.T) *StatusRepository {
	db, err := CreateConnection("prom_test", "prom_tester", "tester")
	if err != nil {
		t.Error("Failed on database connection")
	}

	return NewStatusRepository(db)
}

func TestGetIDByName(t *testing.T) {
	tests := []struct {
		name          string
		statusName    string
		expectedId    int
		expectedError error
	}{
		{
			name:          "success",
			statusName:    "In Progress",
			expectedId:    2,
			expectedError: nil,
		},
		{
			name:          "doesn't_exist",
			statusName:    "Will do",
			expectedId:    -1,
			expectedError: ErrStatusNotFound,
		},
	}

	repo := getStatusRepository(t)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			id, err := repo.GetIDByName(test.statusName)
			if err != test.expectedError {
				t.FailNow()
			}
			if id != test.expectedId {
				t.FailNow()
			}
		})
	}
}
