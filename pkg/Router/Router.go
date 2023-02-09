package Router

import (
	database "FirstApp/pkg/Database"
	"FirstApp/pkg/Users"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func StartServer() *chi.Mux {

	database.SetupDBConnection()
	router := chi.NewRouter()
	router.Mount("/api/users", Users.UsersRoutes())
	log.Fatal(http.ListenAndServe(":8080", router))
	fmt.Println("server is listening on port 8080")
	return router
}
