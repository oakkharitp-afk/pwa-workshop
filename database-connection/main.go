package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {

	/* คำสั่ง SQL สำหรับสร้างตาราง users ใน PostgreSQL
		   CREATE TABLE main.users (
		       id serial PRIMARY KEY,
		       name VARCHAR(60)
		   );

		   	INSERT INTO  main.users (name)
		   	VALUES ('user 1'),
	       		('user 2');
	*/

	postgresUri := `postgresql://admin:admin123@localhost:5432/appdb?sslmode=disable`
	pgCli, err := connectPostgres(postgresUri)
	if err != nil {
		log.Fatal(err)
	}
	defer pgCli.Close()
	fmt.Println("PostgreSQL connected")

	query := `SELECT id,name FROM main.users`
	rows, err := pgCli.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("User ID: %d, Name: %s\n", id, name)
	}

	pgCli.Exec(`INSERT INTO main.users (name) VALUES ($1)`, "hello")

	// เพิ่มข้อมูลผู้ใช้ใหม่
	stmt, err := pgCli.Prepare(`
		INSERT INTO  main.users (name)
		   	VALUES ($1);
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	stmt.Exec("user 3")
	stmt.Exec("user 4")

	/* คำสั่ง MongoDB สำหรับสร้างฐานข้อมูล pwa_training และคอลเล็กชัน users
	   use pwa_training
	       db.createCollection("users")
	       db.users.insertOne({
	           name:"user 1"
	           })
	*/

	mongoUri := `mongodb://admin:admin123@localhost:27017/?authSource=admin`
	ctx := context.Background()

	mgCli, err := connectMongo(mongoUri)
	if err != nil {
		log.Fatal(err)
	}
	defer mgCli.Disconnect(ctx)
	fmt.Println("MongoDB connected")

	collection := mgCli.Database("pwa_training").Collection("users")
	collectionsCursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer collectionsCursor.Close(ctx)

	for collectionsCursor.Next(ctx) {
		var result bson.M
		if err := collectionsCursor.Decode(&result); err != nil {
			log.Fatal(err)
		}

		jsonBytes, err := json.Marshal(result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(jsonBytes))
	}

	collection.InsertOne(context.Background(), bson.M{
		"name":       "hello",
		"created_at": time.Now(),
	})

}
