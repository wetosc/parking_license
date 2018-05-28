package main

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	"./tcpclient"
	"github.com/openalpr/openalpr/src/bindings/go/openalpr"
	"github.com/patrickmn/go-cache"
)

var alpr *openalpr.Alpr

var c *cache.Cache

func main() {
	alpr = openalpr.NewAlpr("eu", "/etc/openalpr/openalpr.conf", "/srv/openalpr/runtime_data")
	defer alpr.Unload()

	if !alpr.IsLoaded() {
		fmt.Println("OpenAlpr failed to load!")
		return
	}
	alpr.SetTopN(5)

	c = cache.New(30*time.Second, 5*time.Minute)

	tcpClient.StartServer(":8123", onConnect)
}

func onConnect(conn net.Conn) {
	client := tcpClient.NewClient(conn)
	client.ReadFileAsync(onImage)
}

func onImage(c *tcpClient.Client, data []byte) {
	if len(data) == 0 {
		return
	}

	if wasCarDetected() {
		fmt.Println("A car was already detected.")
	}

	resultFromBlob, err := alpr.RecognizeByBlob(data)
	checkError(err, "Bad bytes")

	if len(resultFromBlob.Plates) > 0 {
		// fmt.Println("START GUESSING")
		guesses := resultFromBlob.Plates[0].TopNPlates

		for _, guess := range guesses {
			addConfidence(guess.Characters, guess.OverallConfidence)
		}
	}
}

func sendResult(plateNr string) {
	// u, err := url.Parse("http://172.31.166.224:8080/entry/decrement")
	u, err := url.Parse("http://104.196.102.76/entry/decrement")
	checkError(err, "Bad base URL")

	q := url.Values{}
	q.Set("cameraNumber", "test_camera_utm")
	q.Set("carNumber", plateNr)
	u.RawQuery = q.Encode()
	url := u.String()

	fmt.Printf("\n%v\n", url)

	req, err := http.NewRequest("POST", url, nil)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	fmt.Printf("\n%v\n", resp)
	checkError(err, "Bad request")

	defer resp.Body.Close()
}

func checkError(err error, info string) {
	if err != nil {
		fmt.Printf("\n%v : %v", info, err)
	}
}

func addConfidence(plate string, confidence float32) {
	if wasCarDetected() {
		return
	}

	confidence = confidence / 100

	value, found := c.Get(plate)

	if found {
		floatVal := value.(float32)

		floatVal += confidence
		// fmt.Printf("CONFIDENCE: %v", floatVal)
		if floatVal > 3 {

			c.Set("detected", true, 30*time.Second)

			sendResult(plate)

			return
		}

		c.Set(plate, floatVal, cache.DefaultExpiration)

	} else {

		c.Set(plate, confidence, cache.DefaultExpiration)
	}
}

func wasCarDetected() bool {
	value, found := c.Get("detected")
	return found && value.(bool)
}
