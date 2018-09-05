package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Привет, мир!")
	w.Write([]byte("!!!"))
}

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

<<<<<<< HEAD
	http.HandleFunc("/", handler)
	fmt.Println("starting server at :" + port)
	http.ListenAndServe(":"+port, nil)
=======
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	fmt.Println(":" + port)
	router.Run(":" + port)
>>>>>>> 0edde985e4c30f57f07e90709e75e0a4ef80db89
}
