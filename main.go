package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codefordenver/your-colorado-charity/server"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

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

// Retrieves a token from a local file.
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

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func main() {
	// b, err := ioutil.ReadFile("credentials2.json")
	// if err != nil {
	// 	log.Fatalf("Unable to read client secret file: %v", err)
	// }

	// // If modifying these scopes, delete your previously saved token.json.
	// config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	// if err != nil {
	// 	log.Fatalf("Unable to parse client secret file to config: %v", err)
	// }
	// client := getClient(config)

	// srv, err := sheets.New(client)
	// if err != nil {
	// 	log.Fatalf("Unable to retrieve Sheets client: %v", err)
	// }

	// // https://docs.google.com/spreadsheets/d/1IlpUip5zeDDZHfih7tF8Cugu9pqoYPKHq7mSMnbq5vs/edit?usp=sharing
	// spreadsheetID := "1IlpUip5zeDDZHfih7tF8Cugu9pqoYPKHq7mSMnbq5vs"
	// readRange := "Config!A1:B"
	// resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	// if err != nil {
	// 	log.Fatalf("Unable to retrieve data from sheet: %v", err)
	// }

	// // var fields, values string
	// var charity string
	// if len(resp.Values) == 0 {
	// 	fmt.Println("No data found.")
	// } else {
	// 	// for _, row := range resp.Values {
	// 	// 	fields = fmt.Sprintf("%s, data_%s", fields, row[0])
	// 	// 	value := strings.ReplaceAll(fmt.Sprintf("%s", row[1]), "\"", "'")
	// 	// 	value = strings.ReplaceAll(fmt.Sprintf("%s", value), "\n", "")
	// 	// 	values = fmt.Sprintf("%s, \"%s\"", values, value)
	// 	// }
	// 	// fields = strings.Replace(fields, ", ", "", 1)
	// 	// values = strings.Replace(values, ", ", "", 1)

	// 	// fmt.Printf("%d \n", len(values))
	// 	// fmt.Printf("INSERT INTO milehigh_ycc.ycc_data (%s) VALUES (%s)", fields, values)
	// 	// fmt.Printf(fields, values)

	// 	for _, row := range resp.Values {
	// 		var key, val string
	// 		key = strings.ReplaceAll(fmt.Sprintf("%s", row[0]), "\"", "'")
	// 		val = strings.ReplaceAll(fmt.Sprintf("%s", row[1]), "\"", "'")
	// 		val = strings.ReplaceAll(fmt.Sprintf("%s", val), "\n", "")
	// 		charity = charity + `"` + key + `": "` + val + `", `
	// 	}

	// 	charity = strings.TrimSuffix(charity, ", ")
	// 	charity = "{" + charity + "}"

	// 	// charityJson := json.RawMessage(charity)
	// 	// b, err := json.MarshalIndent(&charityJson, "", "\t")
	// 	// if err != nil {
	// 	// 	fmt.Println("error:", err)
	// 	// }
	// 	// fmt.Printf("%T\n", b)
	// 	// os.Stdout.Write()

	// 	// newJson := &db.CharityJSON{}
	// 	// err = json.Unmarshal([]byte(charity), newJson)
	// 	// if err != nil {
	// 	// 	panic(err.Error())
	// 	// }

	// 	// fmt.Printf("%T\n", b)

	// 	// for _, row := range resp.Values {
	// 	// 	var key, val string
	// 	// 	key = strings.ReplaceAll(fmt.Sprintf("%s", row[0]), "\"", "'")
	// 	// 	val = strings.ReplaceAll(fmt.Sprintf("%s", row[1]), "\"", "'")
	// 	// 	val = strings.ReplaceAll(fmt.Sprintf("%s", val), "\n", "")
	// 	// 	charity = append(charity, key, val)
	// 	// }
	// 	// fmt.Println(charity)
	// 	insertCharity := db.InsertData(charity)
	// 	fmt.Println(insertCharity)
	// }
	server.ServerConnect()

}
