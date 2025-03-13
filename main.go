package main

import (
	"fmt"

	"github.com/azacdev/go-react/database"
)

func init() {
	database.ConnectToDB()
	database.SyncDatabase()
}

func main() {
	fmt.Println("Wubba lubba dub dub")
}
