package tts


func CallTTS(i string) (string,string,error) {

	// Get the time

	t,err := GetTime()
	if  err!=nil {
		return "","",err
	}

	//and hash the concatenation of the incomming data and timestamp

	h := MakeSHA256(i+t)

    // sign it

    sig := Sign(h)

    // and return

	return  h,sig,nil
}