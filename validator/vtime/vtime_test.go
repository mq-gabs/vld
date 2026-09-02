package vtime

import (
	"testing"
	"time"
)

func Test_Time_Grouped(t *testing.T) {
	now := time.Now()
	validator := Time("timestamp",
		NotZero(),
		Past(),
	)

	tests := []struct {
		name  string
		value time.Time
		valid bool
	}{
		{
			name:  "valid non-zero past time",
			value: now.Add(-1 * time.Hour),
			valid: true,
		},
		{
			name:  "zero time",
			value: time.Time{},
			valid: false,
		},
		{
			name:  "future time",
			value: now.Add(1 * time.Hour),
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected invalid but got no error")
			}
		})
	}
}

func Test_NotZero(t *testing.T) {
	validator := NotZero()

	tests := []struct {
		name  string
		value time.Time
		valid bool
	}{
		{
			name:  "non-zero time",
			value: time.Now(),
			valid: true,
		},
		{
			name:  "specific time",
			value: time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC),
			valid: true,
		},
		{
			name:  "zero time",
			value: time.Time{},
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected invalid but got no error")
			}
		})
	}
}

func Test_Midnight(t *testing.T) {
	validator := Midnight()

	tests := []struct {
		name  string
		value time.Time
		valid bool
	}{
		{
			name:  "midnight UTC",
			value: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			valid: true,
		},
		{
			name:  "midnight in other timezone",
			value: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
			valid: true,
		},
		{
			name:  "noon",
			value: time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC),
			valid: false,
		},
		{
			name:  "one second after midnight",
			value: time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
			valid: false,
		},
		{
			name:  "one nanosecond after midnight",
			value: time.Date(2020, 1, 1, 0, 0, 0, 1, time.UTC),
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected invalid but got no error")
			}
		})
	}
}

func Test_After(t *testing.T) {
	baseTime := time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC)
	validator := After(baseTime)

	tests := []struct {
		name  string
		value time.Time
		valid bool
	}{
		{
			name:  "after base time",
			value: baseTime.Add(1 * time.Second),
			valid: true,
		},
		{
			name:  "after base time (hours later)",
			value: baseTime.Add(24 * time.Hour),
			valid: true,
		},
		{
			name:  "equal to base time",
			value: baseTime,
			valid: false,
		},
		{
			name:  "before base time",
			value: baseTime.Add(-1 * time.Second),
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected invalid but got no error")
			}
		})
	}
}

func Test_NotAfter(t *testing.T) {
	baseTime := time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC)
	validator := NotAfter(baseTime)

	tests := []struct {
		name  string
		value time.Time
		valid bool
	}{
		{
			name:  "before base time",
			value: baseTime.Add(-1 * time.Second),
			valid: true,
		},
		{
			name:  "equal to base time",
			value: baseTime,
			valid: true,
		},
		{
			name:  "after base time",
			value: baseTime.Add(1 * time.Second),
			valid: false,
		},
		{
			name:  "after base time (hours later)",
			value: baseTime.Add(24 * time.Hour),
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected invalid but got no error")
			}
		})
	}
}

func Test_Before(t *testing.T) {
	baseTime := time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC)
	validator := Before(baseTime)

	tests := []struct {
		name  string
		value time.Time
		valid bool
	}{
		{
			name:  "before base time",
			value: baseTime.Add(-1 * time.Second),
			valid: true,
		},
		{
			name:  "before base time (days earlier)",
			value: baseTime.Add(-24 * time.Hour),
			valid: true,
		},
		{
			name:  "equal to base time",
			value: baseTime,
			valid: false,
		},
		{
			name:  "after base time",
			value: baseTime.Add(1 * time.Second),
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected invalid but got no error")
			}
		})
	}
}

func Test_NotBefore(t *testing.T) {
	baseTime := time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC)
	validator := NotBefore(baseTime)

	tests := []struct {
		name  string
		value time.Time
		valid bool
	}{
		{
			name:  "after base time",
			value: baseTime.Add(1 * time.Second),
			valid: true,
		},
		{
			name:  "equal to base time",
			value: baseTime,
			valid: true,
		},
		{
			name:  "before base time",
			value: baseTime.Add(-1 * time.Second),
			valid: false,
		},
		{
			name:  "before base time (days earlier)",
			value: baseTime.Add(-24 * time.Hour),
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected invalid but got no error")
			}
		})
	}
}

func Test_Between(t *testing.T) {
	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 1, 31, 23, 59, 59, 0, time.UTC)
	validator := Between(start, end)

	tests := []struct {
		name  string
		value time.Time
		valid bool
	}{
		{
			name:  "in the middle",
			value: time.Date(2020, 1, 15, 12, 0, 0, 0, time.UTC),
			valid: true,
		},
		{
			name:  "after start and before end",
			value: start.Add(1 * time.Second),
			valid: true,
		},
		{
			name:  "exactly at start (exclusive)",
			value: start,
			valid: false,
		},
		{
			name:  "exactly at end (exclusive)",
			value: end,
			valid: false,
		},
		{
			name:  "before start",
			value: start.Add(-1 * time.Second),
			valid: false,
		},
		{
			name:  "after end",
			value: end.Add(1 * time.Second),
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected invalid but got no error")
			}
		})
	}
}

func Test_Past(t *testing.T) {
	validator := Past()
	now := time.Now()

	tests := []struct {
		name  string
		value time.Time
		valid bool
	}{
		{
			name:  "past time (1 hour ago)",
			value: now.Add(-1 * time.Hour),
			valid: true,
		},
		{
			name:  "past time (1 day ago)",
			value: now.Add(-24 * time.Hour),
			valid: true,
		},
		{
			name:  "current time (approximately)",
			value: now.Add(100 * time.Millisecond),
			valid: false,
		},
		{
			name:  "future time (1 second from now)",
			value: now.Add(1 * time.Second),
			valid: false,
		},
		{
			name:  "future time (1 hour from now)",
			value: now.Add(1 * time.Hour),
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected invalid but got no error")
			}
		})
	}
}

func Test_Future(t *testing.T) {
	validator := Future()
	now := time.Now()

	tests := []struct {
		name  string
		value time.Time
		valid bool
	}{
		{
			name:  "future time (1 hour from now)",
			value: now.Add(1 * time.Hour),
			valid: true,
		},
		{
			name:  "future time (1 day from now)",
			value: now.Add(24 * time.Hour),
			valid: true,
		},
		{
			name:  "current time (approximately)",
			value: now,
			valid: false,
		},
		{
			name:  "past time (1 second ago)",
			value: now.Add(-1 * time.Second),
			valid: false,
		},
		{
			name:  "past time (1 hour ago)",
			value: now.Add(-1 * time.Hour),
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected invalid but got no error")
			}
		})
	}
}

func Test_Today(t *testing.T) {
	validator := Today()
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	tomorrow := today.Add(24 * time.Hour)
	yesterday := today.Add(-24 * time.Hour)

	tests := []struct {
		name  string
		value time.Time
		valid bool
	}{
		{
			name:  "current time",
			value: now,
			valid: true,
		},
		{
			name:  "midnight today",
			value: today,
			valid: true,
		},
		{
			name:  "noon today",
			value: time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, time.Local),
			valid: true,
		},
		{
			name:  "tomorrow",
			value: tomorrow,
			valid: false,
		},
		{
			name:  "yesterday",
			value: yesterday,
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected invalid but got no error")
			}
		})
	}
}

func Test_SameDate(t *testing.T) {
	baseDate := time.Date(2020, 1, 15, 12, 30, 45, 0, time.UTC)
	validator := SameDate(baseDate)

	tests := []struct {
		name  string
		value time.Time
		valid bool
	}{
		{
			name:  "same date different time",
			value: time.Date(2020, 1, 15, 9, 0, 0, 0, time.UTC),
			valid: true,
		},
		{
			name:  "same date midnight",
			value: time.Date(2020, 1, 15, 0, 0, 0, 0, time.UTC),
			valid: true,
		},
		{
			name:  "same date end of day",
			value: time.Date(2020, 1, 15, 23, 59, 59, 0, time.UTC),
			valid: true,
		},
		{
			name:  "previous day",
			value: time.Date(2020, 1, 14, 12, 30, 45, 0, time.UTC),
			valid: false,
		},
		{
			name:  "next day",
			value: time.Date(2020, 1, 16, 12, 30, 45, 0, time.UTC),
			valid: false,
		},
		{
			name:  "different month",
			value: time.Date(2020, 2, 15, 12, 30, 45, 0, time.UTC),
			valid: false,
		},
		{
			name:  "different year",
			value: time.Date(2021, 1, 15, 12, 30, 45, 0, time.UTC),
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected invalid but got no error")
			}
		})
	}
}

func Test_Weekday(t *testing.T) {
	validator := Weekday()

	tests := []struct {
		name  string
		value time.Time
		valid bool
	}{
		{
			name:  "Monday",
			value: time.Date(2020, 1, 6, 0, 0, 0, 0, time.UTC), // Monday
			valid: true,
		},
		{
			name:  "Wednesday",
			value: time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), // Wednesday
			valid: true,
		},
		{
			name:  "Friday",
			value: time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC), // Friday
			valid: true,
		},
		{
			name:  "Saturday",
			value: time.Date(2020, 1, 11, 0, 0, 0, 0, time.UTC), // Saturday
			valid: false,
		},
		{
			name:  "Sunday",
			value: time.Date(2020, 1, 12, 0, 0, 0, 0, time.UTC), // Sunday
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected invalid but got no error")
			}
		})
	}
}

func Test_Weekend(t *testing.T) {
	validator := Weekend()

	tests := []struct {
		name  string
		value time.Time
		valid bool
	}{
		{
			name:  "Saturday",
			value: time.Date(2020, 1, 11, 0, 0, 0, 0, time.UTC), // Saturday
			valid: true,
		},
		{
			name:  "Sunday",
			value: time.Date(2020, 1, 12, 0, 0, 0, 0, time.UTC), // Sunday
			valid: true,
		},
		{
			name:  "Monday",
			value: time.Date(2020, 1, 6, 0, 0, 0, 0, time.UTC), // Monday
			valid: false,
		},
		{
			name:  "Wednesday",
			value: time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), // Wednesday
			valid: false,
		},
		{
			name:  "Friday",
			value: time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC), // Friday
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected invalid but got no error")
			}
		})
	}
}

func Test_MatchWeekday(t *testing.T) {
	validator := MatchWeekday(time.Monday)

	tests := []struct {
		name  string
		value time.Time
		valid bool
	}{
		{
			name:  "Monday",
			value: time.Date(2020, 1, 6, 0, 0, 0, 0, time.UTC), // Monday
			valid: true,
		},
		{
			name:  "Monday different date",
			value: time.Date(2020, 1, 13, 0, 0, 0, 0, time.UTC), // Monday
			valid: true,
		},
		{
			name:  "Tuesday",
			value: time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC), // Tuesday
			valid: false,
		},
		{
			name:  "Sunday",
			value: time.Date(2020, 1, 12, 0, 0, 0, 0, time.UTC), // Sunday
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected invalid but got no error")
			}
		})
	}
}

func Test_ComplexTime_Validation(t *testing.T) {
	// Complex validation: must be in past, on a weekday, and after a specific date
	startDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	validator := Time("appointment",
		NotZero(),
		Past(),
		After(startDate),
		Weekday(),
	)

	// Use a known past weekday (Monday 2020-01-20 is in the past and after 2020-01-01)
	validPastWeekday := time.Date(2020, 1, 20, 10, 0, 0, 0, time.UTC) // Monday

	tests := []struct {
		name  string
		value time.Time
		valid bool
	}{
		{
			name:  "valid past weekday after start date",
			value: validPastWeekday,
			valid: true,
		},
		{
			name:  "future date",
			value: time.Now().Add(24 * time.Hour),
			valid: false,
		},
		{
			name:  "zero time",
			value: time.Time{},
			valid: false,
		},
		{
			name:  "past but on weekend",
			value: time.Date(2020, 1, 19, 0, 0, 0, 0, time.UTC), // Sunday
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)

			if tt.valid && err != nil {
				t.Fatalf("expected valid but got error: %v", err)
			}

			if !tt.valid && err == nil {
				t.Fatalf("expected invalid but got no error")
			}
		})
	}
}

func Test_MatchWeekday_AllDays(t *testing.T) {
	// Test each day of the week
	testDays := []struct {
		weekday time.Weekday
		date    time.Time
	}{
		{time.Sunday, time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC)},
		{time.Monday, time.Date(2020, 1, 6, 0, 0, 0, 0, time.UTC)},
		{time.Tuesday, time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC)},
		{time.Wednesday, time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC)},
		{time.Thursday, time.Date(2020, 1, 9, 0, 0, 0, 0, time.UTC)},
		{time.Friday, time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC)},
		{time.Saturday, time.Date(2020, 1, 11, 0, 0, 0, 0, time.UTC)},
	}

	for _, td := range testDays {
		t.Run(td.weekday.String(), func(t *testing.T) {
			validator := MatchWeekday(td.weekday)
			err := validator(td.date)

			if err != nil {
				t.Fatalf("expected valid for %s but got error: %v", td.weekday.String(), err)
			}
		})
	}
}
