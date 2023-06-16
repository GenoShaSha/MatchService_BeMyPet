package dbaccess

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// func ConnectToDb() (context.Context, *mongo.Database, *mongo.Collection) {
// 	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	err = client.Connect(ctx)

// 	database := client.Database("UserService_BeMyPet")
// 	usersCollection := database.Collection("user")
// 	return ctx, database, usersCollection
// }

func ConnectToDb() *sql.DB {
	db, err := sql.Open("mysql", "4d2adpc6lotoch53wrs2:pscale_pw_dlxvbdlZQuGvcyR1K8hCu0PVSLUbfLWW1hdBf11Zzrx@tcp(aws.connect.psdb.cloud)/matching_service_bemypet?tls=true")
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping: %v", err)
	}
	log.Println("Successfully connected to PlanetScale!")
	return db
}
