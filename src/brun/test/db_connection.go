package main

import "ensifera/db"

func main() {
	db.InitDB()
	//db.Delete()
	//db.Update()
	//db.Query()
	db.QueryRow()
}
