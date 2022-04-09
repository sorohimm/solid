package errors

import "errors"

var ErrConnectionRefuse = errors.New("comments_tree: can't connect to the server, please check your internet connection")
var ErrInternal = errors.New("service unavailable")
var ErrCommentNotFound = errors.New("comment not found")
