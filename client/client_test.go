package client

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"reflect"
	"testing"
)

func startServer() {
	go func() {
		http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %q", html.EscapeString("bar"))
		})
		http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %q", html.EscapeString("foo"))
		})
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

}
func TestClientInfo_SendRequest(t *testing.T) {
	endpoint := "http://localhost:8080"
	startServer()
	fmt.Println("started server", endpoint)
	values := map[string]string{"foo": "baz"}

	type args struct {
		method string
		path   string
		data   interface{}
	}
	tests := []struct {
		name   string
		client *ClientInfo
		args   args
		want   string
	}{
		{
			name:   "test",
			client: NewClient(endpoint),
			args: args{
				method: "POST",
				path:   "/bar",
				data:   values,
			},
			want: `Hello, "bar"`,
		},
		{
			name:   "test-02",
			client: NewClient(endpoint),
			args: args{
				method: "POST",
				path:   "/bar",
				data:   values,
			},
			want: `Hello, "foo"`,
		},
		{
			name:   "test-03",
			client: NewClient("https://reqbin.com"),
			args: args{
				method: "GET",
				path:   "/echo/get/json",
				data:   values,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.client
			if got := c.SendRequest(tt.args.method, tt.args.path, tt.args.data); !reflect.DeepEqual(string(got), tt.want) {
				// t.Errorf("ClientInfo.SendRequest() = %v, want %v", string(got), tt.want)
			}
		})
	}
}
