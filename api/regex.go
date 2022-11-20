package api

import "regexp"

var emailRegex, _ = regexp.Compile(`.+@.+\..+`)
var openApiUrlRegex, _ = regexp.Compile(`https?://[a-zA-Z.]+(:\d+)?(/.*)?`)
var pathRegex, _ = regexp.Compile(`.+`)
var methodRegex, _ = regexp.Compile(`(GET)|(HEAD)|(POST)|(PUT)|(DELETE)|(CONNECT)|(OPTIONS)|(TRACE)|(PATCH)`)
