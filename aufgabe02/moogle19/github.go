package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/tabwriter"
)

type Foo struct {
	ID   int64 `json:"id"`
	User struct {
		Login string `json:"login"`
	}
	Title string `json:"title"`
}

func main() {
	username := os.Args[1]
	repos := os.Args[2]
	var ownusername string
	var ownsecret string
	if len(os.Args) == 5 {
		ownusername = os.Args[3]
		ownsecret = os.Args[4]
	}
	arr := fetch(username, repos, ownusername, ownsecret)
	printout(arr)
}

func fetch(uname string, repo string, ownuname string, ownsecret string) []Foo {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/repos/"+uname+"/"+repo+"/pulls", nil)
	if ownuname != "" && ownsecret != "" {
		req.SetBasicAuth(ownuname, ownsecret)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	var arr []Foo
	err = json.Unmarshal(data, &arr)

	return arr
}

func printout(fooarr []Foo) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 10, 0, '\t', 0)
	for _, f := range fooarr {
		fmt.Fprintln(w, fmt.Sprintf("%d", f.ID)+"\t\t"+f.User.Login+"\t\t"+f.Title)
	}
	w.Flush()
}
