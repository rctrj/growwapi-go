package growwapi

import (
	"encoding/json"
	"fmt"
	"time"
)

// NullableTime represents a nullable version of Time
// See Time for more details
type NullableTime struct {
	*time.Time
}

// Time represents a time coming from groww apis.
// There are multiple representations for the same time present in APIs. This aims to parse all variations properly
type Time struct {
	time.Time
}

func (t *Time) UnmarshalCSV(in string) error {
	// csv only received time.DateOnly
	parsed, err := time.Parse(time.DateOnly, in)
	if err != nil {
		return fmt.Errorf("time.Parse(%q): %w", in, err)
	}

	t.Time = parsed
	return nil
}

func (t *Time) UnmarshalJSON(bytes []byte) error {
	// check if it's an integer. In that case, parse it as epoch
	var epoch int64
	if err := json.Unmarshal(bytes, &epoch); err == nil {
		t.Time = time.Unix(epoch, 0)
		return nil
	}

	var asString string
	if err := json.Unmarshal(bytes, &asString); err != nil {
		return fmt.Errorf("json.Unmarshal(%q): cannot parse as time", string(bytes))
	}

	layouts := []string{
		time.RFC3339,
		time.DateTime,
		time.DateOnly,
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04:05",
	}

	for _, layout := range layouts {
		parsed, err := time.Parse(layout, asString)
		if err == nil {
			t.Time = parsed
			return nil
		}
	}

	return fmt.Errorf("cannot parse as time: %q", asString)
}

func (t *NullableTime) UnmarshalCSV(in string) error {
	if in == "" {
		t.Time = nil
		return nil
	}

	var parsed Time
	if err := parsed.UnmarshalCSV(in); err != nil {
		return err
	}

	t.Time = &parsed.Time
	return nil
}
