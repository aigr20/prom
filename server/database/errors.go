package database

import "errors"

var ErrConnectionFailed = errors.New("connecting to the database failed")
var ErrProjectNotFound = errors.New("could not find requested project")
var ErrIssuesNotFound = errors.New("could not find issues for project")
var ErrProjectScan = errors.New("project could not be scanned")
var ErrProjectCreate = errors.New("project could not be created")
var ErrIssueCreate = errors.New("issue could not be created")
var ErrIssueNotFound = errors.New("could not find issue")
var ErrStatusNotFound = errors.New("could not find status")
var ErrUpdateFieldCount = errors.New("amount of fields must match amount of values passed")
var ErrNoFields = errors.New("at least one field and value must be passed")
var ErrUpdateFailed = errors.New("update failed")
var ErrIllegalFieldName = errors.New("illegal field name passed")
