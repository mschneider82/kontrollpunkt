package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mschneider82/kontrollpunkt/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var myDB = NewInMemoryDB()

// getCheckForName should query database interface and returning the check
func getCheckForName(c *gin.Context) {
	instance := c.Param("instance")
	category := c.Param("category")
	checkname := c.Param("checkname")

	replyValue := "unknown"
	var replyHint string

	for _, cat := range myDB.Instance[instance] {
		if strings.ToLower(cat.CategoryName) == strings.ToLower(category) {
			if strings.ToLower(cat.CheckName) == strings.ToLower(checkname) {
				replyValue = cat.CheckValue.String()
				replyHint = cat.CheckHint
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": replyValue, "hint": replyHint})
}

// putCheckForName should update database interface and return OK
func putCheckForName(c *gin.Context) {
	//PUT /inst/category/checkname?status=[error,warn,ok]&updateIntervalSecs=300&expirySecs=900
	instance := c.Param("instance")
	category := c.Param("category")
	checkname := c.Param("checkname")
	status := NewCheckStatus(c.DefaultQuery("status", "ok"))
	//updateIntervalSecs := c.DefaultQuery("updateIntervalSecs", "300") // TODO use instance or global config?
	//expirySecs := c.DefaultQuery("expirySecs", "2880")                // TODO use instance or global config?
	hint := c.DefaultPostForm("hint", "")
	myCheck := Category{
		CategoryName: category,
		CheckName:    checkname,
		CheckValue:   status,
		CheckHint:    hint,
	}
	myDB.Instance[instance] = append(myDB.Instance[instance], myCheck)
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	var configuration config.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	log.Infof("database uri is %s", configuration.Database.ConnectionURI)
	log.Infof("port for this application is %d", configuration.Server.Port)

	r := gin.Default()

	for instance, accountmap := range configuration.Instances {
		log.Infof("xx: %v", instance)
		log.Infof("v: %v", accountmap) // -> gin.Accounts

		authorized := r.Group("/:instance", gin.BasicAuth(accountmap))
		authorized.GET("/:category/:checkname", getCheckForName)
		authorized.PUT("/:category/:checkname", putCheckForName)
	}
	r.Run(fmt.Sprintf(":%d", configuration.Server.Port))
}
