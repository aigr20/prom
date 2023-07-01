package database

import "errors"

var ErrConnectionFailed = errors.New("connecting to the database failed")
var ErrProjectNotFound = errors.New("could not find requested project")
var ErrIssuesNotFound = errors.New("could not find issues for project")
var ErrProjectScan = errors.New("project could not be scanned")
