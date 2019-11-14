package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type myjar struct {
	jar map[string][]*http.Cookie
}

func (p *myjar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	fmt.Printf("The URL is : %s\n", u.String())
	fmt.Printf("The cookie being set is : %s\n", cookies)
	p.jar[u.Host] = cookies
}

func (p *myjar) Cookies(u *url.URL) []*http.Cookie {
	fmt.Printf("The URL is : %s\n", u.String())
	fmt.Printf("Cookie being returned is : %s\n", p.jar[u.Host])
	return p.jar[u.Host]
}

func main() {
	link := &url.URL{Host: "http://madmoa.d092.cp:8082/madrhas/identification.ex"}

	proxyUrl, err := url.Parse("http://proxy.bercy.cp:3128")
	// handle err
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}

	// cookie := http.Cookie{Name: "JSESSIONID", Value: "138752C24DF4B4D6848CAD28C2E03BF5"}

	// jar := &myjar{}
	// jar.jar = make(map[string][]*http.Cookie)
	// var expectedCookies = []*http.Cookie{&cookie}
	// jar.SetCookies(link, expectedCookies)
	// client.Jar = jar

	req, err := http.NewRequest("Get", link.Host, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// req.Header.Set("Host", "madmoa.d092.cp:8082")
	// req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11")
	// req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	// req.Header.Set("Accept-Language", "fr-FR,fr;q=0.5")
	// req.Header.Set("Accept-Encoding", "gzip, deflate")
	// req.Header.Set("Referer", "")
	// //req.Header.Set("Cookie", "JSESSIONID=138752C24DF4B4D6848CAD28C2E03BF5")
	// req.Header.Set("Accept-Charset", "ISO-8859-1,utf-8;q=0.7,*;q=0.3")
	// req.Header.Set("Connection", "keep-alive")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))

}
