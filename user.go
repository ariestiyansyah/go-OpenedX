package main

import (
    //"bytes"
    //"time"
    "database/sql"
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
)


func main() {
    db, err := sql.Open("mysql", "user:password@tcp(localhost)/edxapp")
    if err != nil {
        fmt.Print(err.Error())
    }
    defer db.Close()
    err = db.Ping()
    if err != nil {
        fmt.Print(err.Error())
    }

    type auth_userprofile struct {
        Id  int
        Name string
        Language string
        Location string
        Gender string
        City string
        Country string
        Year_Of_Birth string
    }

    router := gin.Default()

    router.GET("/auth_userprofile/:id", func(c *gin.Context) {
        var (
            auth_userprofile auth_userprofile
            result gin.H
        )
        id := c.Param("id")
        row := db.QueryRow("select id, name, language, location, gender, city, country, year_of_birth from auth_userprofile where id = ?;", id)
        err = row.Scan(&auth_userprofile.Id, &auth_userprofile.Name, &auth_userprofile.Language, &auth_userprofile.Location, &auth_userprofile.Gender, &auth_userprofile.City, &auth_userprofile.Country, &auth_userprofile.Year_Of_Birth)
        if err != nil {
            result = gin.H{
                "result": nil,
                "count": 0,
            }
        } else {
            result = gin.H{
                "result": auth_userprofile ,
                "count": 1,
            }
        }
        c.JSON(http.StatusOK, result)
    })
    router.Run(":3000")
}
