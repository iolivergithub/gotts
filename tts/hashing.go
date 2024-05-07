package tts

import(
	"crypto/sha256"
)

func MakeSHA256(p string) string {

    h := sha256.New()
    h.Write([]byte(p))
    hs := h.Sum(nil)

    return string(hs)
}