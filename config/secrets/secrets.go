//=============================================================================
// developer: boxlesslabsng@gmail.com
// Manage environment variables, currently using a .env file 
// to run the application, you need a valid .env file
//
// to run with docker, load .env into docker with this command
// sudo docker run --env-file .env .......
//=============================================================================
 
/**
 * Define package secrets
 **
 * @struct Secrets
 * @api - Loads .env file into application instance
 * @GetSecrets() return secrets 
**/

package secrets

import (
	"log"
	"os"
	"path"
	"runtime"

	"github.com/joho/godotenv"
)

// Secrets ...
type Secrets struct {
	DatabaseName        string
	DatabaseURL         string
	ApplicationPort     string
	ApplicationName     string
	JWTSecrets          string
	Environment         string
	EyowoBaseUrl		string
	EyowoAppKey			string
	EyowoRefreshToken	string
	EyowoCheckoutBaseURL string
	APIBaseURL			string
}

// api ...
func init() {
	_, b, _, _ := runtime.Caller(0)
	BasePath := path.Dir(b)
	if err := godotenv.Load(BasePath + "/.env"); err != nil {
		log.Println("Error loading .env file")
	}
}

// GetSecrets ...
func GetSecrets() Secrets {
	var secrets Secrets

	secrets.DatabaseName = os.Getenv("DATABASE_NAME")
	secrets.DatabaseURL = os.Getenv("DATABASE_URL")
	secrets.ApplicationPort = os.Getenv("PORT")
	secrets.ApplicationName = os.Getenv("APPLICATION_NAME")
	secrets.JWTSecrets = os.Getenv("JWT_SECRET")
	secrets.Environment = os.Getenv("ENVIRONMENT")
	secrets.EyowoBaseUrl = os.Getenv("EYOWO_BASE_URL")
	secrets.EyowoAppKey = os.Getenv("EYOWO_APP_KEY")
	secrets.EyowoRefreshToken = os.Getenv("EYOWO_REFRESH_TOKEN")
	secrets.APIBaseURL = os.Getenv("API_BASE_URL")
	secrets.EyowoCheckoutBaseURL = os.Getenv("EYOWO_CHECKOUT_BASE_URL")

	return secrets
}
