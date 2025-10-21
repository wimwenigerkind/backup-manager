package main

import (
	"fmt"

	"github.com/wimwenigerkind/backup-manager/server/internal/config"
)

func main() {
	cfg := config.LoadConfig()
	fmt.Println(*cfg)
}
