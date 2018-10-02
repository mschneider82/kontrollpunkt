### Design Draft

This is a golang learning project: simple http monitoring daemon.

kontrollpunkt listen on https://0.0.0.0:1122 (may use of go lego library to get a good certificate)

Http Auth is used for simplicity, a few kinds of users:

admin (can see the webpage and may delete checks)
monitor (can only the webpage)
categoryname (each category can have its own users)

PUT /inst/category/checkname?status=[error,warn,ok]&updateIntervalSecs=300&expirySecs=900
FORM hint=some more details

&updateIntervalSecs=300&expirySecs=900 are optional and can be set globally (for each category?)
When time.Now+updateIntervalSecs is reached check will change the state

on expirySecs key is deleted (sets redis expiry ts..)

Own Goals:
* interface for data storage https://github.com/dgraph-io/badger and redis 
* viper for config loadings etc.