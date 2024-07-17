package webFramework

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//required installation commands
// go get github.com/gofiber/fiber/v2
// go get github.com/jmoiron/sqlx
// go get github.com/lib/pq

func TestGofiber() {
	fmt.Println("----------------------------------------Start Gofiber----------------------------------------")
	initDatabase()
	app := fiber.New()

	app.Get("/users", getUsers)
	app.Get("/users/:id", getUser)
	app.Post("/users", createUser)
	app.Put("/users/:id", updateUser)
	app.Delete("/users/:id", deleteUser)

	log.Fatal(app.Listen(":3000"))
	fmt.Println("----------------------------------------End Gofiber  ----------------------------------------")
}

type UserInfo struct {
	ID    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}

var db *sqlx.DB

func initDatabase() {
	var err error
	db, err = sqlx.Connect("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	schema := `CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100),
        email VARCHAR(100)
    )`
	db.MustExec(schema)
}

func getUsers(c *fiber.Ctx) error {
	var users []UserInfo
	err := db.Select(&users, "SELECT id, name, email FROM users")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(users)
}

func getUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user UserInfo
	err := db.Get(&user, "SELECT id, name, email FROM users WHERE id=$1", id)
	if err != nil {
		return c.Status(404).SendString("User not found")
	}
	return c.JSON(user)
}

func createUser(c *fiber.Ctx) error {
	user := new(UserInfo)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	result, err := db.NamedExec("INSERT INTO users (name, email) VALUES (:name, :email)", user)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	id, _ := result.LastInsertId()
	user.ID = int(id)
	return c.JSON(user)
}

func updateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(UserInfo)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	num, _ := strconv.Atoi(id)
	user.ID = num
	_, err := db.NamedExec("UPDATE users SET name=:name, email=:email WHERE id=:id", user)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(user)
}

func deleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := db.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.SendString("User successfully deleted")
}
