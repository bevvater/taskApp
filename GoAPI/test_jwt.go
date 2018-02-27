package main

import (
	"encoding/json"
	"fmt"
//	"strings"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"crypto/rsa"
	"github.com/codegangsta/negroni"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

//	"reflect"
)
const (
	privKeyPath = "keys/app.rsa"
	pubKeyPath = "keys/app.rsa.pub"
)
var (

//	verifyKey, signKey interface{}
	verifyKey *rsa.PublicKey
	signKey *rsa.PrivateKey
	
	session *mgo.Session
	databaseName string	
)

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//struct User for parsing login credentials
type User struct {
//	ID int `json:"id"`
	Fullname string `json:"fullname"`
	Role string 	`json:"role"`
	Username string `json:"username"`
	Password string `json:"password"`

}

type Response struct {
	Data string `json:"data"`
}


// type Token struct {
// 	Token string `json:"token"`
// }

type ResponseLogin struct {
	Token string 	`json:"token"`
	User User 		`json:"user"`
}
//var session *mgo.Session

// type DataStore struct {
// 	session *mgo.Session
// }

// func NewDataStore() *DataStore {
// 	ds := &DataStore{
// 		session: session.Copy(),
// 	}
// 	return ds
// }

func initKeys() {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal("Error reading private key 1")
		return
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatal("Error reading private key 2")
		return
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal("Error reading private key 3")
		return
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatal("Error reading private key 4")
		return
	}

}

func StartServer() {
	// Non-Protected Endpoint


	log.Println("Now listening.....port :8080")

    router := mux.NewRouter()

	router.HandleFunc("/login", LoginHandler).Methods("POST")
	router.HandleFunc("/register", RegisterUser).Methods("POST")

	// Protected Endpoints
	router.Handle("/resource", negroni.New(
		negroni.HandlerFunc(ValidateTokenMiddleware),
		negroni.Wrap(http.HandlerFunc(ProtectedHandler)),
		))

    http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router))

}

func main() {
	var err error
	session, err = mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	databaseName = "appTask"
	initKeys()
	StartServer()
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{"Gained access to protected resource"}
	JsonResponse(response, w)
}


func RegisterUser(w http.ResponseWriter, r *http.Request) {
	user := User{"","member","",""}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Error!")
		fmt.Println(err)
		return
	}

	sessionCopy := session.Copy()
	defer sessionCopy.Close()
	c := sessionCopy.DB(databaseName).C("users")
	err = c.Insert(&user)
	if err != nil {
		log.Fatal(err)
	}

	// -------------------------------------------------------
	// ds := NewDataStore()
	// defer ds.session.Close()

	// c := ds.session.DB("appTask").C("users")
	// err = c.Insert(&user)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// ---------------------------------------------------------
	// session, err := mgo.Dial("localhost")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer session.Close()
	// session.SetMode(mgo.Monotonic, true)
	// c := session.DB("appTask").C("users")
	
	// err = c.Insert(&user)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	response := "hello"
	JsonResponse(response, w)
}


//reads the login credentials, checks them and creates JWT the token
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user UserCredentials
	//decode into User struct

	//r.ParseForm()

	//fmt.Println(r.FormValue("username"))
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Wrong info")
		fmt.Println(err)
		return
	}

	sessionCopy := session.Copy()
	defer sessionCopy.Close()
	c := sessionCopy.DB(databaseName).C("users")

	var resultUser User
	err = c.Find(bson.M{"username": user.Username, "password": user.Password}).One(&resultUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Login failed")
	} else {
		token := jwt.New(jwt.SigningMethodRS256)
		claims := make(jwt.MapClaims)
		claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
		claims["iat"] = time.Now().Unix()
		token.Claims = claims

		tokenString, err := token.SignedString(signKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Error while signing the token")
		}
		//responseToken := Token{tokenString}
		//JsonResponse(response, w)

		resp := ResponseLogin{tokenString, resultUser}
		JsonResponse(resp, w)
	}

	// if strings.ToLower(user.Username) != "khoavic" || user.Password != "pass" {
	// 	w.WriteHeader(http.StatusForbidden)
	// 	fmt.Println("Error logging in")
	// 	fmt.Fprint(w, "Invalid credentials")
	// 	return
	// }


}

func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {

		// fmt.Println(verifyKey)
		// fmt.Println("----------------")
		// fmt.Println(request.AuthorizationHeaderExtractor)
		// fmt.Println("----------------")

		return verifyKey, nil
	})

	if(err==nil && token.Valid) {
		next(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized access to this resource")
	}

	// token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
	// 	return verifyKey, nil
	// 	})
	// if(err==nil && token.Valid) {
	// 	next(w, r)
	// } else {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	fmt.Fprint(w, "Unauthorized access to this resource")
	// }

}

func JsonResponse(response interface{}, w http.ResponseWriter) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

