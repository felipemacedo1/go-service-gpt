package middleware

import (
	"net/http"
	"sync"
	"github.com/sirupsen/logrus"
	"time"

	"github.com/labstack/echo/v4"
)

type RateLimiter struct {
	logger    *logrus.Logger
	limit     int
	duration  time.Duration

	ipRecords map[string][]time.Time
	mu        sync.Mutex
}

func NewRateLimiter(limit int, duration time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:     limit,
		logger:    logrus.New(),
		duration:  duration,
		ipRecords: make(map[string][]time.Time),
	}
}

func (rl *RateLimiter) Limit(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ip := c.RealIP()
		now := time.Now()

		rl.mu.Lock()
		defer rl.mu.Unlock()

		records, ok := rl.ipRecords[ip]
		if !ok {
			rl.ipRecords[ip] = []time.Time{now}
			return next(c)
		}

		// Clean up old records
		var recentRecords []time.Time
		for _, record := range records {
			if now.Sub(record) <= rl.duration {
				recentRecords = append(recentRecords, record)
			}
		}

		if len(recentRecords) >= rl.limit {
			rl.logger.WithFields(logrus.Fields{
				"ip":      ip,
				"limit":   rl.limit,
				"current": len(recentRecords),
				"at":      now,
			}).Warn("Rate limit exceeded")

			return c.String(http.StatusTooManyRequests, "Too Many Requests")
		}

		recentRecords = append(recentRecords, now)
		rl.ipRecords[ip] = recentRecords
		return next(c)
	}
}