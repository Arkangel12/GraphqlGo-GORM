package config

import (
	"encoding/json"
	"log"
	"os"
)

//ConfigFilePath Ruta del archivo del archivo config.json.
var ConfigFilePath string

// Configuration Contains basic settings for connection to db
type Configuration struct {
	Name        string
	Server      string
	Host        string
	Port        string
	Dbname      string
	User        string
	Password    string
	Email 		string
	EPassword 	string
}

func init() {
	ConfigFilePath = "./config.json"
}

//AppConfig Almacena la configuración del archivo config.json.
var AppConfig Configuration

//LoadAppConfig Lee el archivo config.json y lo decodifica en AppConfig
func LoadAppConfig() {

	file, err := os.Open(ConfigFilePath)

	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(file)
	AppConfig = Configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}

}
