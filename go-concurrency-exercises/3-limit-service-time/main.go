//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
//

package main

import (
	"fmt"
	"time"
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
}

var freemium_threshold int64 = 10

func (u *User) UpdateTime(time_elapsed int64) {
	u.TimeUsed += time_elapsed
	fmt.Println("User", u.ID, "has accumulated", u.TimeUsed, "seconds of total processing time.")
}

// Runs the processes requested by users. Returns false if process had to be killed
// In this case:
// processes 2 & 5 should always work (since u2 is premium)
// processes 1 & 3 should always work (since < 10s)
// process 4 should fail (non-premium + >10s)
func HandleRequest(process func(), u *User) bool {
	if u.IsPremium {
		process()
		return true
	} else {
		start := time.Now()
		process()
		elapsed := time.Since(start).Seconds()
		if int64(elapsed) > freemium_threshold {
			fmt.Println("Request took", elapsed, "s >", freemium_threshold, "s.")
			return false
		} else {
			return true
		}
	}
}

// In this case:
// processes 2 & 5 should always work (since u2 is premium)
// process 1 => u1 accumulates ~6s
// process 3 => u1 accumulates ~6s ==> 12s total
// process 4 should fail (non-premium + >10s total)
func HandleRequestPerUser(process func(), u *User) bool {
	if u.IsPremium {
		process()
		return true
	} else {
		if u.TimeUsed < freemium_threshold {
			start := time.Now()
			process()
			elapsed := time.Since(start).Seconds()
			u.UpdateTime(int64(elapsed))
			if u.TimeUsed > freemium_threshold {
				return false
			} else {
				return true
			}
		}
		return false
	}
}

func main() {
	RunMockServer()
}
