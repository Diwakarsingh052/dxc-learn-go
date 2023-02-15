package middlewares

import "log"

type Mid struct {
	log *log.Logger
}

func NewMid(log *log.Logger) Mid {
	return Mid{log: log}
}
