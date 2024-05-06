package tts

import(
	"crypto/sha256"

	"github.com/beevik/ntp"
)

func getTime() (string,error) {
    time, err := ntp.Time("fi.pool.ntp.org")

    return time.GoString(),err
}

func CallTTS(i string) (string,string,error) {

	// Get the time

	t,err := getTime()
	if  err!=nil {
		return "","",err
	}

	//so, concatenate the given hash i with the timestamp t

	c := i+t

	//and hash it

    h := sha256.New()
    h.Write([]byte(c))
    hs := h.Sum(nil)

    // sign it

    sig := "signature"

    // and return

	return  string(hs),sig,nil
}