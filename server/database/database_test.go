package database

import (
	"testing"
)

func TestCreateConnection(t *testing.T) {
	tests := []struct {
		name     string
		username string
		password string
		want     error
	}{
		{
			name:     "correct",
			username: "prom_tester",
			password: "tester",
			want:     nil,
		},
		{
			name:     "invalid password",
			username: "prom_tester",
			password: "not_tester",
			want:     ErrConnectionFailed,
		},
		{
			name:     "invalid username",
			username: "not_prom",
			password: "tester",
			want:     ErrConnectionFailed,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, err := CreateConnection("prom_test", test.username, test.password)
			if db != nil {
				defer db.Close()
			}
			if err != test.want {
				t.Fail()
			}
		})
	}
}
