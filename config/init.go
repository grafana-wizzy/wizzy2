package config

import (
	"github.com/grafana-wizzy/wizzy2/log"
	"os"
	"path"
)

const directory = "conf"
const file = "wizzy.json"

func InitCmd() error {
	// Create directory
	err := os.Mkdir(directory, 0777)
	if err != nil {
		return err
	}
	log.Info("configuration directory created")

	// Save file
	config := Config()
	err = config.WriteConfigAs(path.Join(directory, file))
	if err != nil {
		return err
	}
	log.Info("configuration file created")

	log.Info("wizzy successfully initialized")
	return nil
}
