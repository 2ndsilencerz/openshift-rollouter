package openshift

import (
	"crypto/tls"
	"encoding/base64"
	"log"
	"net/http"
	"net/http/httputil"
	"openshift-rollouter/config"
	"regexp"
	"strings"
)

var tempUrl = ""

func Auth() string {
	conf := config.NewConfig().Viper
	oauthUrl := conf.GetString("openshift.auth.uri")

	// auth with username and password
	username := conf.GetString("openshift.auth.username")
	password := conf.GetString("openshift.auth.password")
	auth := username + ":" + password
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))

	req, _ := http.NewRequest("POST", oauthUrl, nil)
	//req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Authorization", basicAuth)
	req.Header.Set("User-Agent", "curl/7.64.1")
	return regexAndReplace(request(req))
}

func closeClient(response *http.Response) {
	_ = response.Body.Close()
}

type loggingTransport struct {
	rt http.RoundTripper
}

func (t *loggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	//dump, err := httputil.DumpRequestOut(req, true)
	//if err != nil {
	//	log.Println("Error dumping request:", err)
	//} else {
	//log.Println("Dumping Request")
	//log.Println(string(dump))
	//}

	//log.Println("Do RoundTrip")
	resp, err := t.rt.RoundTrip(req)

	if err != nil {
		log.Println("Error sending request:", err)
	} else {
		//dump, err := httputil.DumpResponse(resp, true)
		if strings.Contains(resp.Header.Get("Location"), "token=sha") {
			tempUrl = resp.Header.Get("Location")
			//log.Println(tempUrl)
		}
		tmpResp, _ := httputil.DumpResponse(resp, true)
		log.Println(string(tmpResp))
		//if err != nil {
		//	log.Println("Error dumping response:", err)
		//} else {
		//	log.Println("Dumping Response")
		//	log.Println(string(dump))
		//}
	}

	return resp, err
}

func request(req *http.Request) string {
	loggingTransport := &loggingTransport{
		rt: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	client := &http.Client{
		Transport: loggingTransport,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
	}

	defer closeClient(resp)
	return tempUrl
}

func regexAndReplace(url string) string {
	var result string
	pattern := regexp.MustCompile("access_token=(.*)&expire")
	res := pattern.FindAllString(url, -1)
	for _, v := range res {
		result = v
	}
	//log.Println(result)
	result = strings.Replace(result, "access_token=", "", -1)
	result = strings.Replace(result, "&expire", "", -1)
	log.Println(result)
	return result
}
