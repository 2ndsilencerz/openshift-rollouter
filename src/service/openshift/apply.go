package openshift

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"openshift-rollouter/model"
	"openshift-rollouter/service/openshift/utils"
	"os"
	"strings"
	"time"
)

//func Apply(c *gin.Context) {
//	namespace := c.Param("namespace")
//	tokenAuth := Auth()
//	apiUrl := config.NewConfig().Viper.GetString("openshift.api.apply.uri")
//	apiUrl = strings.Replace(apiUrl, "<namespace>", namespace, -1)
//
//	yamlContent := saveAndLoad(c)
//	contents := utils.ReadYaml(yamlContent)
//	//log.Println(contents)
//
//	resources := utils.Categorize(contents)
//	applyConfigMap(apiUrl, tokenAuth, "ConfigMap", resources.ConfigMap)
//	applySecret(apiUrl, tokenAuth, "Secret", resources.Secret)
//	applyService(apiUrl, tokenAuth, "Service", resources.Service)
//	//applyIngress(apiUrl, tokenAuth, "Ingress", resources.Ingress)
//	applyRoute(apiUrl, tokenAuth, "Route", resources.Route)
//	applyDeployment(apiUrl, tokenAuth, "Deployment", resources.Deployment)
//	applyStatefulSet(apiUrl, tokenAuth, "StatefulSet", resources.StatefulSet)
//	applyDaemonSet(apiUrl, tokenAuth, "DaemonSet", resources.DaemonSet)
//	applyCronJob(apiUrl, tokenAuth, "CronJob", resources.CronJob)
//	applyDeploymentConfig(apiUrl, tokenAuth, "DeploymentConfig", resources.DeploymentConfig)
//
//	c.JSON(http.StatusOK, contents)
//}

func deleteFile(location string) {
	err := os.Remove(location)
	if err != nil {
		log.Println(err)
	}
}

func saveAndLoad(c *gin.Context) []byte {
	fileName := saveFile(c)

	if len(fileName) > 0 {
		//log.Println("Uploading Success with Name " + fileName)
		defer deleteFile(fileName)
		yamlFile, err := os.ReadFile(fileName)
		if err != nil {
			log.Println(err)
			return nil
		}
		return yamlFile
	}
	return nil
}

func saveFile(c *gin.Context) string {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
		return ""
	}
	//log.Println(file.Filename)
	fileName := "temp" + time.Now().Format("150405.000") + ".yaml"

	err = c.SaveUploadedFile(file, fileName)
	if err != nil {
		log.Println(err)
		return ""
	}
	return fileName
}

//func applyResource(url string, tokenAuth string, kind string, contents []string) {
//	if len(contents) == 0 {
//		return
//	}
//	resourceKind := strings.ToLower(kind) + "s"
//	url = strings.Replace(url, "<resource>", resourceKind, -1)
//	apiVersion := ApiVersion([]byte(contents[0]))
//	url = strings.Replace(url, "<apiVersion>", apiVersion, -1)
//
//	if kind == "StatefulSet" ||
//		kind == "Deployment" ||
//		kind == "DaemonSet" ||
//		kind == "CronJob" ||
//		kind == "DeploymentConfig" {
//		url = strings.Replace(url, "/api/", "/apis/", -1)
//	}
//
//}

func urlReplace(url string, kind string, apiVersion string) string {
	resourceKind := strings.ToLower(kind) + "s"
	url = strings.Replace(url, "<resource>", resourceKind, -1)
	url = strings.Replace(url, "<apiVersion>", apiVersion, -1)
	return url
}

func applyDeployment(url string, tokenAuth string, kind string, contents []model.Deployment) {
	if len(contents) == 0 {
		return
	}
	apiVersion := contents[0].ApiVersion
	url = urlReplace(url, kind, apiVersion)
	url = strings.Replace(url, "/api/", "/apis/", -1)

	for _, v := range contents {
		method := http.MethodPost
		currentUrl := url
		exist, resourceVersion, _ := TestIfExist(url, tokenAuth, v.Metadata.Name, kind)
		if exist {
			method = http.MethodPut
			v.Metadata.ResourceVersion = resourceVersion
		}
		vMarshalled, _ := yaml.Marshal(v)
		sendToOpenshiftApi(currentUrl, tokenAuth, method, vMarshalled, false)
	}
}

func applyStatefulSet(url string, tokenAuth string, kind string, contents []model.StatefulSet) {
	if len(contents) == 0 {
		return
	}
	apiVersion := contents[0].ApiVersion
	url = urlReplace(url, kind, apiVersion)
	url = strings.Replace(url, "/api/", "/apis/", -1)

	for _, v := range contents {
		method := http.MethodPost
		currentUrl := url
		exist, resourceVersion, _ := TestIfExist(url, tokenAuth, v.Metadata.Name, kind)
		if exist {
			method = http.MethodPut
			v.Metadata.ResourceVersion = resourceVersion
		}
		vMarshalled, _ := yaml.Marshal(v)
		sendToOpenshiftApi(currentUrl, tokenAuth, method, vMarshalled, false)
	}
}

func applyCronJob(url string, tokenAuth string, kind string, contents []model.CronJob) {
	if len(contents) == 0 {
		return
	}
	apiVersion := contents[0].ApiVersion
	url = urlReplace(url, kind, apiVersion)
	url = strings.Replace(url, "/api/", "/apis/", -1)

	for _, v := range contents {
		method := http.MethodPost
		currentUrl := url
		exist, resourceVersion, _ := TestIfExist(url, tokenAuth, v.Metadata.Name, kind)
		if exist {
			method = http.MethodPut
			v.Metadata.ResourceVersion = resourceVersion
		}
		vMarshalled, _ := yaml.Marshal(v)
		sendToOpenshiftApi(currentUrl, tokenAuth, method, vMarshalled, false)
	}
}

func applyDaemonSet(url string, tokenAuth string, kind string, contents []model.DaemonSet) {
	if len(contents) == 0 {
		return
	}
	apiVersion := contents[0].ApiVersion
	url = urlReplace(url, kind, apiVersion)
	url = strings.Replace(url, "/api/", "/apis/", -1)

	for _, v := range contents {
		method := http.MethodPost
		currentUrl := url
		exist, resourceVersion, _ := TestIfExist(url, tokenAuth, v.Metadata.Name, kind)
		if exist {
			method = http.MethodPut
			v.Metadata.ResourceVersion = resourceVersion
		}
		vMarshalled, _ := yaml.Marshal(v)
		sendToOpenshiftApi(currentUrl, tokenAuth, method, vMarshalled, false)
	}
}

func applyDeploymentConfig(url string, tokenAuth string, kind string, contents []model.DeploymentConfig) {
	if len(contents) == 0 {
		return
	}
	apiVersion := contents[0].ApiVersion
	url = urlReplace(url, kind, apiVersion)
	url = strings.Replace(url, "/api/", "/apis/", -1)

	for _, v := range contents {
		method := http.MethodPost
		currentUrl := url
		exist, resourceVersion, _ := TestIfExist(url, tokenAuth, v.Metadata.Name, kind)
		if exist {
			method = http.MethodPut
			v.Metadata.ResourceVersion = resourceVersion
		}
		vMarshalled, _ := yaml.Marshal(v)
		sendToOpenshiftApi(currentUrl, tokenAuth, method, vMarshalled, false)
	}
}

func applyConfigMap(url string, tokenAuth string, kind string, contents []model.ConfigMap) {
	if len(contents) == 0 {
		return
	}
	apiVersion := contents[0].ApiVersion
	url = urlReplace(url, kind, apiVersion)

	for _, v := range contents {
		method := http.MethodPost
		currentUrl := url
		exist, resourceVersion, _ := TestIfExist(url, tokenAuth, v.Metadata.Name, kind)
		if exist {
			method = http.MethodPut
			v.Metadata.ResourceVersion = resourceVersion
		}
		vMarshalled, _ := yaml.Marshal(v)
		sendToOpenshiftApi(currentUrl, tokenAuth, method, vMarshalled, false)
	}
}

func applySecret(url string, tokenAuth string, kind string, contents []model.Secret) {
	if len(contents) == 0 {
		return
	}
	apiVersion := contents[0].ApiVersion
	url = urlReplace(url, kind, apiVersion)

	for _, v := range contents {
		method := http.MethodPost
		currentUrl := url
		exist, resourceVersion, _ := TestIfExist(url, tokenAuth, v.Metadata.Name, kind)
		if exist {
			method = http.MethodPut
			v.Metadata.ResourceVersion = resourceVersion
		}
		vMarshalled, _ := yaml.Marshal(v)
		sendToOpenshiftApi(currentUrl, tokenAuth, method, vMarshalled, false)
	}
}

func applyService(url string, tokenAuth string, kind string, contents []model.Service) {
	if len(contents) == 0 {
		return
	}
	apiVersion := contents[0].ApiVersion
	url = urlReplace(url, kind, apiVersion)

	for _, v := range contents {
		method := http.MethodPost
		currentUrl := url
		exist, resourceVersion, _ := TestIfExist(url, tokenAuth, v.Metadata.Name, kind)
		if exist {
			method = http.MethodPut
			v.Metadata.ResourceVersion = resourceVersion
		}
		vMarshalled, _ := yaml.Marshal(v)
		sendToOpenshiftApi(currentUrl, tokenAuth, method, vMarshalled, false)
	}
}

//func applyIngress(url string, tokenAuth string, kind string, contents []model.Ingress) {
//	if len(contents) == 0 {
//		return
//	}
//	apiVersion := contents[0].ApiVersion
//	url = urlReplace(url, kind, apiVersion)
//
//	for _, v := range contents {
//		method := http.MethodPost
//		currentUrl := url
//		exist, resourceVersion := TestIfExist(url, tokenAuth, v.Metadata.Name, kind)
//		if exist {
//			method = http.MethodPut
//			v.Metadata.ResourceVersion = resourceVersion
//		}
//		vMarshalled, _ := yaml.Marshal(v)
//		sendToOpenshiftApi(currentUrl, tokenAuth, method, vMarshalled)
//	}
//}

func applyRoute(url string, tokenAuth string, kind string, contents []model.Route) {
	if len(contents) == 0 {
		return
	}
	apiVersion := contents[0].ApiVersion
	url = urlReplace(url, kind, apiVersion)

	for _, v := range contents {
		method := http.MethodPost
		currentUrl := url
		exist, resourceVersion, _ := TestIfExist(url, tokenAuth, v.Metadata.Name, kind)
		if exist {
			method = http.MethodPut
			v.Metadata.ResourceVersion = resourceVersion
		}
		vMarshalled, _ := yaml.Marshal(v)
		sendToOpenshiftApi(currentUrl, tokenAuth, method, vMarshalled, false)
	}
}

func sendToOpenshiftApi(currentUrl string, tokenAuth string, method string, content []byte, rollout bool) {
	// Create an HTTP request with the YAML file contents as the request body
	req, err := http.NewRequest(method, currentUrl, bytes.NewBuffer(content))
	req.Header.Add("Authorization", "Bearer "+tokenAuth)
	if err != nil {
		log.Println(err)
		return
	}

	// Set the Content-Type header to "application/yaml"
	req.Header.Set("Content-Type", "application/yaml")
	if method == http.MethodPatch {
		req.Header.Set("Content-Type", "application/strategic-merge-patch+json")
	}
	if rollout {
		req.Header.Set("Content-Type", "application/json")
	}
	// Send the HTTP request
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	rawReq, _ := httputil.DumpRequest(req, true)
	log.Println("Request: ", string(rawReq))
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	dumpResp, _ := httputil.DumpResponse(resp, true)
	log.Println("Response: ", string(dumpResp))
	closeClient(resp)
}

type TestBody struct {
	Kind     string         `yaml:"kind"`
	Metadata model.Metadata `yaml:"metadata"`
}

func TestIfExist(url string, tokenAuth string, resourceName string, resourceKind string) (bool, string, TestBody) {
	jsonResp := TestBody{}
	url = url + "/" + resourceName
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Authorization", "Bearer "+tokenAuth)
	if err != nil {
		log.Println(err)
		return false, "", jsonResp
	}
	// Send the HTTP request
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	rawReq, _ := httputil.DumpRequest(req, true)
	log.Println("Request: ", string(rawReq))
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return false, "", jsonResp
	}

	rawRes, _ := httputil.DumpResponse(resp, true)
	log.Println("Response: ", string(rawRes))
	rawResp, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(rawResp, &jsonResp)
	if err != nil {
		log.Println(err)
		return false, "", jsonResp
	}
	//log.Println(jsonResp)
	if strings.TrimSpace(jsonResp.Kind) == resourceKind {
		return true, jsonResp.Metadata.ResourceVersion, jsonResp
	}
	return false, "", jsonResp
}

func Apply(c *gin.Context) {
	namespace := c.Param("namespace")
	fileName := saveFile(c)
	c.JSON(utils.Apply(namespace, fileName))
}
