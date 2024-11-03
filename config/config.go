package config

import(
	"fmt"
	"log"
	"os"
	"github.com/jinzhu/gorm"
	_"github.com/lib/pq"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func InitDB(){
	err:= godotenv.Load()
	if err!=nil{
		log.Fatal("Error loading .env file")
	}


	dbURI:=fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
	  os.Getenv("DB_HOST"),
	  os.Getenv("DB_PORT"),
	  os.Getenv("DB_USER"),
	  os.Getenv("DB_NAME"),
	  os.Getenv("DB_PASSWORD"),
      )  
	  db,err:=gorm.Open("postgres",dbURI)
	  if err != nil {
		log.Fatal("Error in connecting the db",err);
	  }
	  DB=db
	  log.Println("Database connected successfully")
}