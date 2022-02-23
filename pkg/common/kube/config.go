package kube

import (
	"fmt"
	"k8s.io/client-go/rest"
	"time"
)

const (
	Namespace              = "default"
	DefaultKubeHostEnvKey  = "KUBE_API_SERVER" // in debug model, replace incluster config
	DefaultKubeTokenEnvKey = "KUBE_SA_TOKEN"   // in debug model, replace incluster config
	DefaultSyncTime        = 30 * time.Minute  // reflush time
	defaultTimeOut         = 10 * time.Second
)

func GetKubeClusterConfig() (cfg *rest.Config, err error) {

	if cfg == nil {
		cfg, err = rest.InClusterConfig()
		if err != nil {
			return nil, fmt.Errorf("Get InClusterConfig Error: %s ", err.Error())
		}
	}
	cfg.Timeout = defaultTimeOut
	return cfg, nil
}
