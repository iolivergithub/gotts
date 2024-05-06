package tts

import(
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"	
)

const PREFIX=""


func StartRESTInterface() {
    router := echo.New()

    router.HideBanner = true

    //not necessary, but I will keep this here because this is now my example of how to use middlewares
    //in echo, plus the import declaration above
    router.Use(middleware.GzipWithConfig(middleware.GzipConfig{ Level: 5,}))
    //router.Use(middleware.Logger())

	setupEndpoints(router)
	

	//start the server
	//router.Logger.Fatal(router.Start(":8081")) 
	router.Start(":8081")

}




func setupEndpoints(router *echo.Echo) {
	router.POST(PREFIX+"/timestamp", ProcessTimestamp)
}

