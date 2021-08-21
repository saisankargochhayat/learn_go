package main

import (
	"fmt"
)

type errorMap struct {
	errorName    string
	errorPattern string
}

// Map of error to possible strings to classify them. Put more generic error categories towards the end.
var errorList = []errorMap{
	{"DatabaseSubscriptionsContextError", "(.+)cosmosdb.documents.azure.com/dbs/customlocations/colls/Subscriptions(.+)context canceled"},
	{"DatabaseConnectResetError", "(.+)cosmosdb.documents.azure.com/dbs/customlocations/colls/Subscriptions(.+)connection reset by peer"},
	{"DatabaseContextError", "(.+)cosmosdb.documents.azure.com(.+)context canceled"},
}

func main() {
	for _, err := range errorList {
		// var test = regexp.MustCompile(errorPattern)
		// var str = `Get \"https://cl-prod-eastus2euap-cosmosdb.documents.azure.com/dbs/customlocations/colls/Subscriptions/docs/204898ee-cd13-4332-b9d4-55ca5c25496d\": read tcp 10.244.2.41:56296->52.138.92.2:443: read: connection reset by peer`
		// fmt.Println(test.MatchString(str), errorName)
		fmt.Println(err.errorName)
	}
}

// Or use package reflect TypeOf
