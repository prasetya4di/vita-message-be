package local_time

import "time"

func CurrentTime() time.Time {
	return time.Now().Local()
}
