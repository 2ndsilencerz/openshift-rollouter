package openshift

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"openshift-rollouter/config"
	"openshift-rollouter/model"
	"strings"
	"time"
)

func Rollout(c *gin.Context) {
	rolloutResponse := "Resource not exist"
	namespace := c.Param("namespace")
	kind := c.Param("kind")
	name := c.Param("name")
	tokenAuth := Auth()
	apiUrl := config.NewConfig().Viper.GetString("openshift.api.apply.uri")
	apiUrl = strings.Replace(apiUrl, "/api/", "/apis/", -1)
	apiUrl = strings.Replace(apiUrl, "<namespace>", namespace, -1)
	apiUrl = strings.Replace(apiUrl, "<resource>", kind, -1)
	apiUrl = strings.Replace(apiUrl, "<apiVersion>", "apps/v1", -1)
	testUrl := apiUrl
	apiUrl = apiUrl + "/" + name
	yamlUpdate := struct {
		Metadata model.Metadata `yaml:"metadata" json:"metadata"`
	}{}
	if kind[len(kind)-1:] == "s" {
		kind = kind[:len(kind)-1]
	}
	var kindToTest string
	switch kind {
	case "deployment":
		kindToTest = "Deployment"
		break
	case "statefulset":
		kindToTest = "StatefulSet"
		break
	}

	exist, _, yamlResp := TestIfExist(testUrl, tokenAuth, name, kindToTest)
	if exist {
		timeNow := time.Now()
		//randSource := rand.NewSource(timeNow.UnixMilli())
		//randGenerator := rand.New(randSource)
		//rng := strconv.Itoa(randGenerator.Intn(10))
		yamlUpdate.Metadata.Name = yamlResp.Metadata.Name
		yamlUpdate.Metadata.Labels = yamlResp.Metadata.Labels
		//yamlUpdate.Metadata.ResourceVersion = timeNow.Format("2006150105") + rng
		yamlUpdate.Metadata.ResourceVersion = yamlResp.Metadata.ResourceVersion
		//yamlUpdate.Metadata.Generation = yamlResp.Metadata.Generation + 1
		if yamlResp.Metadata.Annotations == nil {
			yamlUpdate.Metadata.Annotations = make(map[string]string, 1)
		} else {
			yamlUpdate.Metadata.Annotations = yamlResp.Metadata.Annotations
		}
		yamlUpdate.Metadata.Annotations["date"] = timeNow.Format("2006-01-02 15:04:05")
		//if len(yamlUpdate.Metadata.Annotations["deployment.kubernetes.io/revision"]) > 0 {
		//	currentRevision, _ := strconv.Atoi(yamlUpdate.Metadata.Annotations["deployment.kubernetes.io/revision"])
		//	currentRevision += 1
		//	yamlUpdate.Metadata.Annotations["deployment.kubernetes.io/revision"] = strconv.Itoa(currentRevision)
		//}
		jsonPatch, _ := json.Marshal(yamlUpdate)
		apiUrl = apiUrl + "/rollout"
		//apiUrl = strings.Replace(apiUrl, "/apps/", "/apps.openshift.io/", -1)
		sendToOpenshiftApi(apiUrl, tokenAuth, http.MethodPost, jsonPatch, true)
		rolloutResponse = "Resource deployed"
	}
	c.JSON(http.StatusOK, rolloutResponse)
}
