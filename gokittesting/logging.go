package main

import (
	"time"

	"github.com/go-kit/kit/log"
)

//wraps our service up with extra logging stuff
// this is an example of a service middleware
type LoggingMiddleware struct {
	logger log.Logger
	next   StringService
}

// Implements StringService
func (loggingMiddleware LoggingMiddleware) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		loggingMiddleware.logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = loggingMiddleware.next.Uppercase(s)
	return
}

// Implements StringService
func (loggingMiddleware LoggingMiddleware) Count(s string) (n int) {
	defer func(begin time.Time) {
		loggingMiddleware.logger.Log(
			"method", "count",
			"input", s,
			"n", n,
			"took", time.Since(begin),
		)
	}(time.Now())

	n = loggingMiddleware.next.Count(s)
	return

}
