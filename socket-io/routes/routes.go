/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application.
| Enjoy building your API!
|
*/

package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/letsgo-framework/examples/socket-io/controllers"
	"time"
)

var SocketioServer *socketio.Server


// PaveRoutes sets up all api routes
func PaveRoutes() *gin.Engine {
	r := gin.Default()

	// CORS
	corsConfig := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true,
	}
	r.Use(cors.New(corsConfig))

	// Grouped api
	v1 := r.Group("/api/v1")
	{
		v1.GET("/", controllers.Greet)

		auth := AuthRoutes(v1)
		auth.GET("/", controllers.Verify)
	}

	socketRoutes := r.Group("socket.io")
	{
		socketRoutes.GET ( "/", controllers.SocketHandler)
		socketRoutes.POST ( "/", controllers.SocketHandler)
	}

	return r
}


