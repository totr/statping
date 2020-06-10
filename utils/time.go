package utils

import (
	"fmt"
	"time"
)

// Now returns the UTC timestamp
func Now() time.Time {
	return time.Now().UTC()
}

// FormatDuration converts a time.Duration into a string
func FormatDuration(d time.Duration) string {
	var out string
	if d.Hours() >= 24 {
		out = fmt.Sprintf("%0.0f day", d.Hours()/24)
		if (d.Hours() / 24) >= 2 {
			out += "s"
		}
		return out
	} else if d.Hours() >= 1 {
		out = fmt.Sprintf("%0.0f hour", d.Hours())
		if d.Hours() >= 2 {
			out += "s"
		}
		return out
	} else if d.Minutes() >= 1 {
		out = fmt.Sprintf("%0.0f minute", d.Minutes())
		if d.Minutes() >= 2 {
			out += "s"
		}
		return out
	} else if d.Seconds() >= 1 {
		out = fmt.Sprintf("%0.0f second", d.Seconds())
		if d.Seconds() >= 2 {
			out += "s"
		}
		return out
	} else if rev(d.Hours()) >= 24 {
		out = fmt.Sprintf("%0.0f day", rev(d.Hours()/24))
		if rev(d.Hours()/24) >= 2 {
			out += "s"
		}
		return out
	} else if rev(d.Hours()) >= 1 {
		out = fmt.Sprintf("%0.0f hour", rev(d.Hours()))
		if rev(d.Hours()) >= 2 {
			out += "s"
		}
		return out
	} else if rev(d.Minutes()) >= 1 {
		out = fmt.Sprintf("%0.0f minute", rev(d.Minutes()))
		if rev(d.Minutes()) >= 2 {
			out += "s"
		}
		return out
	} else {
		out = fmt.Sprintf("%0.0f second", rev(d.Seconds()))
		if rev(d.Seconds()) >= 2 {
			out += "s"
		}
	}
	return out
}

func rev(f float64) float64 {
	return f * -1
}
