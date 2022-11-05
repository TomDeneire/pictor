package lib

import "os"

// This function takes a JSON string and
// creates a file db.js with a function that
// returns said string as a Javascript object
func CreateJSDatabaseFile(data string) error {
	start := "function getDB() {"
	end := "}"
	content := start + data + end
	f, err := os.Create("db.js")
	if err != nil {
		return err
	}
	f.Write([]byte(content))
	return nil
}
