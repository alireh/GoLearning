package webFramework

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Installation commands
//go get github.com/gin-gonic/gin
//go get -u github.com/jinzhu/gorm
//go get -u github.com/jinzhu/gorm/dialects/postgres

// User represents a user in the database
type User struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
}

// Initialize the database connection
func initDB() *gorm.DB {
	dsn := "host=localhost user=postgres dbname=postgres sslmode=disable password=postgres"
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})
	return db
}

func TestGinPostgres() {
	fmt.Println("----------------------------------------Start GinPostgres----------------------------------------")
	db := initDB()
	defer db.Close()

	router := gin.Default()

	// Create a new user
	router.POST("/users", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	})

	// Get all users
	router.GET("/users", func(c *gin.Context) {
		var users []User
		if err := db.Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, users)
	})

	// Get a user by ID
	router.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user User
		if err := db.First(&user, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, user)
	})

	// Update a user by ID
	router.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user User
		if err := db.First(&user, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Save(&user)
		c.JSON(http.StatusOK, user)
	})

	// Delete a user by ID
	router.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&User{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
	})

	router.Run(":8080")
	fmt.Println("----------------------------------------End GinPostgres  ----------------------------------------")
}
