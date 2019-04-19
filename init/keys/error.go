package keys

import "github.com/sungora/app/lg"

const (
	Err1000 = 1000
	Err1001 = 1001
	Err1002 = 1002
	Err1003 = 1003
	Err1004 = 1004
	Err1005 = 1005
)

func init() {
	lg.SetMessages(map[int]string{
		1000: "Message format Fmt from 1000: %s",
		1001: "Message format Fmt from 1001: %s",
		1002: "Message format Fmt from 1002: %s",
		1003: "Message format Fmt from 1003: %s",
		1004: "Message format Fmt from 1004: %s",
		1005: "Message format Fmt from 1005: %s",
	})
}
