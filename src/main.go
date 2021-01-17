package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/lindamanf/gcnote/app"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Credentials struct {
	Cid string `json:"cid"`
	Csecret string `json:"csecret"`
}

type User struct {
	Sub string `json:"sub"`
	Name string `json:"name"`
	GivenName string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Profile string `json:"profile"`
	Picture string `json:"picture"`
	Email string `json:"email"`
	EmailVerified string `json:"email_verified"`
	Gender string `json:"gender"`
}

var (
	cred Credentials
	conf *oauth2.Config
	state string
	store = sessions.NewCookieStore([]byte("secret"))
)

func init() {
	file, err := ioutil.ReadFile("./config/creds.json")
	if err != nil {
		log.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	json.Unmarshal(file, &cred)

	conf = &oauth2.Config{
		ClientID:     cred.Cid,
		ClientSecret: cred.Csecret,
		RedirectURL:  "http://127.0.0.1:9090/auth",
		Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email", // You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
		},
		Endpoint: google.Endpoint,
	}
}

func indexHundler() {
	
}

func main() {
	app.StartApp()
}
