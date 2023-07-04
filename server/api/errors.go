package api

import "errors"

var ErrGetProjects = errors.New("failed retrieving projects")
var ErrBadRequest = errors.New("bad request")
