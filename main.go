package main

import (
	"fmt"
	"github.com/Yoak3n/EdgeGPT-http/api/router"
	"github.com/Yoak3n/EdgeGPT-http/config"
	"github.com/Yoak3n/EdgeGPT-http/internal/database"
)

// Created at 2023/4/10 14:46
// Created by Yoake

func main() {
	r := router.R
	defer database.CloseConnect()
	err := r.Run(fmt.Sprintf("127.0.0.1:%d", config.Preset.Server.Port))
	if err != nil {
		panic(err)
	}
}
