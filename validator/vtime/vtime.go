package vtime

import (
	"errors"
	"time"

	"github.com/mq-gabs/vld/internal/utils"
)

var (
	ErrTimeMustNotBeZero    = errors.New("time must not be zero")
	ErrTimeMustBeMidnight   = errors.New("time must be midnight")
	ErrTimeMustBeAfter      = errors.New("time must be after")
	ErrTimeMustNotBeAfter   = errors.New("time must not be after")
	ErrTimeMustBeBefore     = errors.New("time must be before")
	ErrTimeMustNotBeBefore  = errors.New("time must not be before")
	ErrTimeMustBeBetween    = errors.New("time must be between")
	ErrTimeMustBeInPast     = errors.New("time must be in past")
	ErrTimeMustBeInFuture   = errors.New("time must be in future")
	ErrTimeMustBeToday      = errors.New("time must be today")
	ErrTimeMustBeSameDate   = errors.New("time must be same date")
	ErrTimeMustBeWeekday    = errors.New("time must be weekday")
	ErrTimeMustBeWeekend    = errors.New("time must be weekend")
	ErrTimeMustMatchWeekday = errors.New("time must match weekday")
)

// TimeValidator is a function type for validating time.Time values
type TimeValidator func(time.Time) error

// Time groups multiple time validators into a single validator, applying them sequentially
// The name parameter identifies the value being validated for error reporting
func Time(name string, validators ...TimeValidator) TimeValidator {
	return func(t time.Time) error {
		err := utils.NewErrorGroup(utils.WithName(name), utils.WithSeparator(","))

		for _, validate := range validators {
			err.Join(validate(t))
		}

		if !err.IsNil() {
			return err
		}

		return nil
	}
}

// NotZero validates that the time is not the zero value
func NotZero() TimeValidator {
	return func(t time.Time) error {
		if t.IsZero() {
			return ErrTimeMustNotBeZero
		}

		return nil
	}
}

// Midnight validates that the time is exactly midnight (00:00:00) on any date
func Midnight() TimeValidator {
	return func(t time.Time) error {
		if !utils.IsMidnight(t) {
			return ErrTimeMustBeMidnight
		}

		return nil
	}
}

// After validates that the time is strictly after the given time
func After(o time.Time) TimeValidator {
	return func(t time.Time) error {
		if !t.After(o) {
			return ErrTimeMustBeAfter
		}

		return nil
	}
}

// NotAfter validates that the time is not strictly after the given time
// This means the time must be equal to or before the given time
func NotAfter(o time.Time) TimeValidator {
	return func(t time.Time) error {
		if t.After(o) {
			return ErrTimeMustNotBeAfter
		}

		return nil
	}
}

// Before validates that the time is strictly before the given time
func Before(o time.Time) TimeValidator {
	return func(t time.Time) error {
		if !t.Before(o) {
			return ErrTimeMustBeBefore
		}

		return nil
	}
}

// NotBefore validates that the time is not strictly before the given time
// This means the time must be equal to or after the given time
func NotBefore(o time.Time) TimeValidator {
	return func(t time.Time) error {
		if t.Before(o) {
			return ErrTimeMustNotBeBefore
		}

		return nil
	}
}

// Between validates that the time is strictly between two times (exclusive on both ends)
// The time must be after 'b' and before 'a'
func Between(b, a time.Time) TimeValidator {
	return func(t time.Time) error {
		if !(t.After(b) && t.Before(a)) {
			return ErrTimeMustBeBetween
		}

		return nil
	}
}

// Past validates that the time is in the past (before the current time)
func Past() TimeValidator {
	return func(t time.Time) error {
		now := time.Now()

		if !t.Before(now) {
			return ErrTimeMustBeInPast
		}

		return nil
	}
}

// Future validates that the time is in the future (after the current time)
func Future() TimeValidator {
	return func(t time.Time) error {
		now := time.Now()

		if !t.After(now) {
			return ErrTimeMustBeInFuture
		}

		return nil
	}
}

// Today validates that the time is on today's date
// The time can be at any time during today, not just midnight
func Today() TimeValidator {
	return func(t time.Time) error {
		now := time.Now()

		if !utils.IsSameDate(t, now) {
			return ErrTimeMustBeToday
		}

		return nil
	}
}

// SameDate validates that the time is on the same date as the given time
// The time of day does not matter, only the date portion is compared
func SameDate(o time.Time) TimeValidator {
	return func(t time.Time) error {
		if !utils.IsSameDate(t, o) {
			return ErrTimeMustBeSameDate
		}

		return nil
	}
}

// Weekday validates that the time falls on a weekday (Monday-Friday)
func Weekday() TimeValidator {
	return func(t time.Time) error {
		if !utils.IsWeekday(t) {
			return ErrTimeMustBeWeekday
		}

		return nil
	}
}

// Weekend validates that the time falls on a weekend day (Saturday or Sunday)
func Weekend() TimeValidator {
	return func(t time.Time) error {
		if !utils.IsWeekend(t) {
			return ErrTimeMustBeWeekend
		}

		return nil
	}
}

// MatchWeekday validates that the time falls on the specified day of the week
func MatchWeekday(d time.Weekday) TimeValidator {
	return func(t time.Time) error {
		if t.Weekday() != d {
			return ErrTimeMustMatchWeekday
		}

		return nil
	}
}
