package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/mitubaEX/graphQL_sample/application/graphql_util"
)

func main() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Print(err)
		}

		result := executeQuery(fmt.Sprintf("%s", body), graphql_util.Schema)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		fmt.Println(result.Errors)
	}
	return result
}
