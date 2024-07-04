package swagger

import (
	"log"
	"memorize/pkg/files"
	"os"
	"os/exec"
)

func InitSwagger() {
	// swag init -g .\cmd\memorize\main.go -o ./docs --pd
	cmd := exec.Command("swag", "init", "-g", ".\\cmd\\memorize\\main.go", "-o", "./docs", "--pd")
	rootPath, err := files.FindProjectRoot()
	if err != nil {
		log.Panicf("swag init command failed: %s", err.Error())
	}

	os.Chdir(*rootPath)
	if err := cmd.Run(); err != nil {
		log.Panicf("swag init command failed: %s", err.Error())
	}
}
