package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli"
)

// Response structs
type Response struct {
	Image string `json:"image"`
	Body  string `json:"body"`
}

func main() {
	app := cli.NewApp()
	app.Name = "horesase"
	app.Usage = "Display http://jigokuno.com/ images and text at random"
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) error {
		client := &http.Client{}

		request, err := http.NewRequest("GET", "http://horesase.github.io/horesase-boys/meigens.json", nil)
		request.Header.Add("Accept", "application/jsonl")
		request.Header.Add("Accept-Encoding", "deflate")
		request.Header.Add("Accept-Encoding", "gzip")

		resp, err := client.Do(request)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		var data []Response
		if err := json.Unmarshal([]byte(string(body)), &data); err != nil {
			fmt.Println("JSON Unmarshal error:", err)
			return nil
		}

		rand.Seed(time.Now().UnixNano())
		target := data[rand.Intn(len(data))]
		fmt.Println(strings.Replace(target.Body, "\n", "", -1))
		fmt.Println(strings.Replace(target.Image, "\n", "", -1))
		return nil
	}

	app.Run(os.Args)
}
