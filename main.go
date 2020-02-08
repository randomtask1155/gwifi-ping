package main


import (
	"net/http"
	"strconv"
	"os"
	"time"
	"log"
	"crypto/tls"
	"io/ioutil"
)

var (
	pingHost string
	pingInterval int64
	client *http.Client
	logger         *log.Logger
)


func execPing() {
	logger.Println("ping start")
	r, err := client.Get(pingHost)
	if err != nil {
		logger.Println(err)
		return
	}

	_, err = ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Println(err)
		return
	}
	defer r.Body.Close()
	logger.Println("ping end")
}

func main() {
	logger = log.New(os.Stdout, "logger: ", log.Ldate|log.Ltime|log.Lshortfile)
	pingHost = os.Getenv("PING_HOST")
	pingInt, err := strconv.Atoi(os.Getenv("PING_INTERVAL"))
	if err != nil {
		panic(err)
	}
	pingInterval = int64(pingInt)


	tr := &http.Transport{
		MaxIdleConns:       -1,
		IdleConnTimeout:    30 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client = &http.Client{Transport: tr,
			Timeout: 30 * time.Second,
	}

	for {
		execPing()
		time.Sleep(time.Duration(pingInterval) * time.Second)
	}



}