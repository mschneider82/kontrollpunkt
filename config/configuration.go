package config

import "github.com/gin-gonic/gin"

type Configuration struct {
	Server    ServerConfiguration
	Database  DatabaseConfiguration
	Instances map[string]gin.Accounts
}
