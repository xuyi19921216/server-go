package service

import (
	"fmt"
	"testing"
)

func TestRemoteConfig(t *testing.T) {
	configFetcher, _ := NewConfig(
		"1111",
		WithTimeout(1),
		WithCluster("default"),
	)
	config, _ := configFetcher.FetchConfigByKey("key")
	fmt.Println(config)

}
