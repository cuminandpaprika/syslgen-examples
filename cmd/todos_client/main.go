package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptrace"

	"github.com/anz-bank/syslgen-examples/gen/todos"
	"github.com/anz-bank/syslgen-examples/pkg/restlib"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	todo   = kingpin.Command("todo", "Get todos from server.")
	todoID = todo.Arg("id", "Todo ID").Required().Int64()

	posts = kingpin.Command("posts", "Get posts from server.")

	comment      = kingpin.Command("comment", "Create comments on server.")
	commentTitle = comment.Arg("title", "Comment Title").Required().String()
	commentBody  = comment.Arg("body", "Comment Content").Required().String()
)

func withTrace(ctx context.Context) context.Context {
	trace := &httptrace.ClientTrace{
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", dnsInfo)
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", connInfo.Conn.RemoteAddr())
		},
	}
	return httptrace.WithClientTrace(ctx, trace)
}

func prettyPrint(i interface{}) string {
	s, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return ""
	}
	return string(s)
}

func printResult(result *restlib.HTTPResult, err error) {
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else {
		fmt.Printf("Status: %d\nValue: \n%+v\n", result.HTTPResponse.StatusCode, prettyPrint(result.Response))
	}
}

func main() {
	httpClient := http.Client{}
	client := todos.NewClient(&httpClient, "http://localhost:8080")
	ctx, cancel := context.WithCancel(context.Background())
	ctx = withTrace(ctx)
	defer cancel()

	switch kingpin.Parse() {
	case "todo":
		printResult(client.GetTodosID(ctx, nil, *todoID))
	case "posts":
		printResult(client.GetPosts(ctx, nil))
	case "comment":
		post := todos.Post{
			ID:     1,
			Body:   *commentBody,
			Title:  *commentTitle,
			UserID: 4,
		}
		printResult(client.PostComments(ctx, nil, &post))
	}

}
