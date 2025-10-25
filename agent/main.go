package main

import (
	"fmt"

	"github.com/wimwenigerkind/backup-manager/agent/internal/config"
)

func main() {
	fmt.Println("Starting Agent...")
	_ = config.LoadConfig()
}
