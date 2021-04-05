package UsersDB

import (
	"fmt"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
	"github.com/IBM/go-sdk-core/core"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	//Client *sql.DB

	Client cloudantv1.CloudantV1

	envMap map[string]string

	errEnv error
)

func init() {

	/* envMap, errEnv = godotenv.Read(".env")

	fmt.Println("errEnv: ", errEnv)

	if errEnv != nil {
		panic(errEnv)
	}

	username := envMap["MYSQL_USERS_USERNAME"]
	password := envMap["MYSQL_USERS_PASSWORD"]
	host := envMap["MYSQL_USERS_HOST"]
	schema := envMap["MYSQL_USERS_SCHEMA"]

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, schema)

	var err error
	Client, err = sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	} */

	fmt.Println("Starting Cloud Service and DataBase Configuration")

	envMap, errEnv = godotenv.Read(".env")

	if errEnv != nil {
		panic(errEnv)
	}

	ApiKeyCloud := envMap["CLOUDANT_APIKEY"]

	Auth := &core.IamAuthenticator{
		ApiKey: ApiKeyCloud,
	}

	CloudURL := envMap["CLOUDANT_URL"]

	// 1. Create a Cloudant Instance Client with IAM Authenticator ============

	client, err := cloudantv1.NewCloudantV1(
		&cloudantv1.CloudantV1Options{
			URL:           CloudURL,
			Authenticator: Auth,
		},
	)
	if err != nil {
		panic(err)
	}

	// 2. Get server information ===========================================
	serverInformationResult, _, err := client.GetServerInformation(
		client.NewGetServerInformationOptions(),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Server Version: ", *serverInformationResult.Version)
	fmt.Println("CouchDB: ", *serverInformationResult.Couchdb)

	// 3. Get database information for "users" ==========================
	dbName := envMap["CLOUDANT_DBNAME"]

	DbInfo, _, err := client.GetDatabaseInformation(
		client.NewGetDatabaseInformationOptions(dbName),
	)

	if err != nil {
		panic(err)
	}

	Client = *client

	fmt.Printf("DataBase [%s] successfully configured \n", *DbInfo.DbName)

}
