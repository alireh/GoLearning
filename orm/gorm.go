package orm

import "fmt"

import "log"
import "gorm.io/gorm"
import "gorm.io/driver/postgres"

//to installing gorm
//go get gorm.io/gorm
//to configure for postgres
//go get gorm.io/driver/postgres

// Product model
type Product struct {
    ID        uint   `gorm:"primaryKey"`
    Name      string `gorm:"not null"`
    Price     float64 `gorm:"not null"`
    CreatedAt int64  `gorm:"autoCreateTime"`
    UpdatedAt int64  `gorm:"autoUpdateTime"`
}

func TestGorm() {
	fmt.Println("----------------------------------------Start Gorm----------------------------------------")
	 // Set up the PostgreSQL connection details
	 dsn := "host=localhost user=postgres password=postgres dbname=test port=5432"

	 
    // Connect to the PostgreSQL database
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    // Migrate the Product model
    err = db.AutoMigrate(&Product{})
    if err != nil {
        log.Fatalf("Failed to migrate the database: %v", err)
    }

    fmt.Println("Connected to the PostgreSQL database!")

    // Create a new product
    product := Product{
        Name:  "Example Product",
        Price: 19.99,
    }

    // Save the product to the database
    result := db.Create(&product)
    if result.Error != nil {
        log.Fatalf("Failed to create product: %v", result.Error)
    }
    fmt.Printf("New product created with ID: %d\n", product.ID)

    // Retrieve a product by ID
    var retrievedProduct Product
    result = db.First(&retrievedProduct, 1)
    if result.Error != nil {
        log.Fatalf("Failed to retrieve product: %v", result.Error)
    }
    fmt.Printf("Retrieved product: %+v\n", retrievedProduct)
	fmt.Println("----------------------------------------End Gorm  ----------------------------------------")
}
