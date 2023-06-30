package database

import "errors"

var ErrConnectionFailed = errors.New("connecting to the database failed")
var ErrProjectNotFound = errors.New("could not find requested project")
var ErrProjectScan = errors.New("project could not be scanned")
