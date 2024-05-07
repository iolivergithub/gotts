package tts

import(
	"github.com/beevik/ntp"
)

func GetTime() (string,error) {
    time, err := ntp.Time("fi.pool.ntp.org")

    return time.GoString(),err
}