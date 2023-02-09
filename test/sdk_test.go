package test

import (
	"encoding/json"
	"fmt"
	. "github.com/jdcloud-api/jdcloud-sdk-go/core"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/rds/apis"
	. "github.com/jdcloud-api/jdcloud-sdk-go/services/rds/client"
	"testing"
)

var (
	accessKey = "D4322972304B498C1323F3B9209E24B9"
	secretKey = "05F22DB8053E9FED651A1C8CD2B3AB98"
)

func TestDescribeParameterGroupParameters(t *testing.T) {
	regionId := "cn-north-1"
	parameterGroupId := "mysql-pg-tbar52k5l5"
	credentials := NewCredentials(accessKey, secretKey)
	client := NewRdsClient(credentials)
	req := apis.NewDescribeParameterGroupParametersRequest(regionId, parameterGroupId)
	resp, err := client.DescribeParameterGroupParameters(req)
	if err != nil {
		return
	}
	marshal, err := json.Marshal(resp)
	if err != nil {
		return
	}
	fmt.Println(len(resp.Result.Parameters))
	fmt.Println(string(marshal))
}

func TestDescribeParameters(t *testing.T) {
	regionId := "cn-north-1"
	pageSize := 100
	credentials := NewCredentials(accessKey, secretKey)
	client := NewRdsClient(credentials)
	req := apis.NewDescribeParameterGroupsRequest(regionId)
	req.PageSize = &pageSize
	resp, err := client.DescribeParameterGroups(req)

	if err != nil {
		return
	}
	marshal, err := json.Marshal(resp)
	if err != nil {
		return
	}

	fmt.Println(string(marshal))
}
