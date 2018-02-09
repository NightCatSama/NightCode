package config

import (
	"fmt"
	"os"  
	"path/filepath" 
	"strings"
	"log"
	
	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	Server server
	Database database
}

type server struct {
	Port string
}

type database struct {
	Host string
	Port string
	Db   string
}

var Conf tomlConfig

const filePath = "../config/config.toml"

func getCurrentDirectory() string {  
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))  
	if err != nil {  
			log.Fatal(err)  
	}
	return strings.Replace(dir, "\\", "/", -1)  
}

func init() {
	absPath := filepath.Join(getCurrentDirectory(), filePath)
	if _, err := toml.DecodeFile(absPath, &Conf); err != nil {
		fmt.Println(err)
		return
	}
}