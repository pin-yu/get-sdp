package main

import (
	"log"
	"os"

	"github.com/bluenviron/gortsplib/v4"
	"github.com/bluenviron/gortsplib/v4/pkg/base"
)

func main() {
	if len(os.Args) < 1 {
		log.Fatalf("No RTSP_URL provided")
	}
	parsedURL, err := base.ParseURL(os.Args[1])
	if err != nil {
		log.Fatalf("parse rtsp url error, err=%s", err)
	}

	log.Printf("url scheme=%s, host=%s", parsedURL.Scheme, parsedURL.Host)

	c := gortsplib.Client{}
	c.Start(parsedURL.Scheme, parsedURL.Host)

	_, res, err := c.Describe(parsedURL)
	if err != nil {
		log.Fatalf("send describe to the server error, err=%s", err)
	}
	log.Println(res)
}
