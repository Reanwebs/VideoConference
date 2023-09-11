package utility

import "time"

func TimeCalculator(start time.Time, end time.Time) (int32, error) {
	duration := end.Sub(start)

	minutes := int32(duration.Minutes())

	return minutes, nil
}
