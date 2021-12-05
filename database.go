package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func database() (*os.File, error) {
	db, err := os.OpenFile("database.json", os.O_RDWR, os.ModePerm)

	return db, err
}

func WriteToDB(data interface{}) error {
	db, err := database()
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if _, err := db.Write([]byte(jsonData)); err != nil {
		return err
	}
	if err = db.Close(); err != nil {
		return err
	}
	//err = ioutil.WriteFile("database.json", jsonData, 0644)
	return nil
}

func ReadFromDB(data interface{}) error {
	db, err := database()
	fileData, err := ioutil.ReadAll(db)
	if err != nil {
		return err
	}
	err = json.Unmarshal(fileData, data)
	if err != nil {
		return err
	}
	return nil
}
