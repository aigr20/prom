package database

import "errors"

var ErrConnectionFailed = errors.New("connecting to the database failed")
var ErrProjectNotFound = errors.New("could not find requested project")
var ErrIssuesNotFound = errors.New("could not find issues for project")
var ErrSprintNotFound = errors.New("could not find requested sprint")
var ErrProjectScan = errors.New("project could not be scanned")
var ErrProjectCreate = errors.New("project could not be created")
var ErrIssueCreate = errors.New("issue could not be created")
var ErrSprintCreate = errors.New("sprint could not be created")
var ErrIssueNotFound = errors.New("could not find issue")
var ErrStatusNotFound = errors.New("could not find status")
var ErrUpdateFieldCount = errors.New("amount of fields must match amount of values passed")
var ErrNoFields = errors.New("at least one field and value must be passed")
var ErrUpdateFailed = errors.New("update failed")
var ErrIllegalFieldName = errors.New("illegal field name passed")
var ErrTagCreate = errors.New("tag could not be created")
var ErrTagDuplicate = errors.New("duplicate tag for project")
var ErrTagNotFound = errors.New("no tags found")
var ErrTagCountFail = errors.New("failed to count tags")
