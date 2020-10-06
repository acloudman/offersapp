package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"offersapp/routes"
)

func main() {
	conn, err := connectDB()
	if err != nil {
		return
	}

	router := gin.Default()
	router.Use(dbMiddleware(*conn))
	userGroup := router.Group("users")
	{
		userGroup.POST("register", routes.UserRegister)
	}
	router.Run(":3000")
}

func connectDB() (c *pgx.Conn, err error) {
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:@localhost:5432/offersapp")
	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n\n", err)
		fmt.Println(err.Error())
	}
	_ = conn.Ping(context.Background())
	return conn, err
}

func dbMiddleware(conn pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", conn)
		c.Next()
	}
}
