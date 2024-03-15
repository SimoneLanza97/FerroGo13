// config/config.go

package config

import (
    "io/ioutil"
    "log"

    "gopkg.in/yaml.v2"
)

// Config rappresenta la configurazione dell'applicazione
type Config struct {
    DBHost     string `yaml:"db_host"`
    DBPort     string `yaml:"db_port"`
    DBUser     string `yaml:"db_user"`
    DBPassword string `yaml:"db_password"`
    DBName     string `yaml:"db_name"`
    // Aggiungi altri campi di configurazione secondo necessità
}

// LoadConfig carica la configurazione dal file YAML
func LoadConfig(filePath string) *Config {
    yamlFile, err := ioutil.ReadFile(filePath)
    if err != nil {
        log.Fatalf("Errore durante la lettura del file di configurazione YAML: %v", err)
    }

    var config Config
    err = yaml.Unmarshal(yamlFile, &config)
    if err != nil {
        log.Fatalf("Errore durante il parsing del file YAML: %v", err)
    }

    return &config
}
