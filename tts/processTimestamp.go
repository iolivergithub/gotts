package tts

import(
	"net/http"

	"github.com/labstack/echo/v4"
)

type timestampInput struct {
	Hash    string    `json:"hash"`
}

type timestampError struct {
	ErrorMessage     string    `json:"errormessage"`
}

type timestampResponse struct {
	TTS           string    `json:"tts"`
	Signature     string    `json:"sig"`

}


func ProcessTimestamp(c echo.Context) error {
	message := new(timestampInput)

	if err := c.Bind(message); err != nil {	
		clienterr := timestampError{ "Invalid syntax: "+err.Error() }
		return c.JSON(http.StatusBadRequest, clienterr)
	}

	res,sig,err := CallTTS( message.Hash )

	if err != nil {	
		clienterr := timestampError{ "Processing Error: "+err.Error() }
		return c.JSON(http.StatusBadRequest, clienterr)
	}

	response := timestampResponse{ res,sig }
	return c.JSON(http.StatusOK, response)

}