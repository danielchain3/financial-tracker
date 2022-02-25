package main

import (
 	"context"
 	"encoding/json"
 	"fmt"
 	"io/ioutil"
 	"log"
 	"net/http"
 	"os"


 	"golang.org/x/oauth2"
 	"golang.org/x/oauth2/google"
 	"google.golang.org/api/sheets/v4"
 	"google.golang.org/api/drive/v3"
 	"google.golang.org/api/option"

 	"drive_connection"
 	"sheets_connection"

)
	

// Retrieve token, saves token, returns generated client
func getClient(config *oauth2.Config) *http.Client{
	token_file := "token,.json"
	token, error := tokenFromFile(token_file)

	if error != nil{
		token = getTokenFromWeb(config)
		saveToken(token_file, token)
	}

 	return config.Client(context.Background(), token)

}

// Request a token from the web, then returns token
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token{
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the link and type in code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}

	return tok
}

// Retrieve token from local file
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Save token to file path
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credentials to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)

	if err != nil {
		log.Fatalf("Unable to cache token: %v", err)
	}

	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

// Create new Sheets
func createSheet(title_name string) {
	ss, err := service.CreateSpreadSheet(spreadsheet.Spreadsheet{
		Properties: spreadsheet.Properties{
			Title: title_name
		},
	})
}

func main() {
	ctx := context.Background()
	b, err := ioutil.ReadFile("credentials.json")

	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying scopes, delete previously saved token.json
	config, err := google.ConfigFromJSON(b, connection)

	if err != nil {
		log.Fatalf("Unable to parse file to config: %v", err)
	}

	client := getClient(config)
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))

	if err != nil {
		log.Fatalf("Unable to retrieve data from sheets: %v", err)
	}

	// get input for sheets name
	fmt.Println("Enter Sheets Name: ")
	var sheetName string
	fmt.Scan(&sheetName)

	// check to see if sheets exists 
	// check all active sheets
	activeSheets, err := srv.Spreadsheets.Sheets
	if err != nil{
		log.Fatalf("Unable to find sheets: %v", err)
	}

	fmt.Println(activeSheets)


	// Prints the names and majors of students in a sample spreadsheet:
    // https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
    /*spreadsheetId := "1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms"
    readRange := "Class Data!A2:E"
    resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
    if err != nil {
            log.Fatalf("Unable to retrieve data from sheet: %v", err)
    }

    if len(resp.Values) == 0 {
            fmt.Println("No data found.")
    } else {
            fmt.Println("Name, Major:")
            for _, row := range resp.Values {
                // Print columns A and E, which correspond to indices 0 and 4.
                fmt.Printf("%s, %s\n", row[0], row[4])
            }
    }*/

}

