package main 

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/IbrahimAbunaib/go-mux-api/Documents/Clean_Architect_Api/infrastructure/datastore"
	"github.com/IbrahimAbunaib/go-mux-api/Documents/Clean_Architect_Api/infrastructure/router"
	"github.com/IbrahimAbunaib/go-mux-api/Documents/Clean_Architect_Api/registry"
)

func main () {
	config.ReadCongig ()

	db := datastore.NewDB()
	db.LogMode(true)
	defer db.Close()

	r := registry.NewRegistry(db)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("server listen to http://localhost" + ":" + config.C.server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}


}