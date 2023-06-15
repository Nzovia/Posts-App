package initializers

import(
	"log"
	"github.com/joho/godotenv"
)

func LoadEnvVariables(){
	// loading the env files
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error occured when Loading .env")
	}

}