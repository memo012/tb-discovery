package errors

import "errors"

var (
	NothingFound       = errors.New("-404") // 啥都木有
	NotModified        = errors.New("-304") // 木有改动
)

