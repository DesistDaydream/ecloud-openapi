// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package main

import (
	"fmt"

	"gitlab.ecloud.com/ecloud/ecloudsdkcore/config"
	"gitlab.ecloud.com/ecloud/ecloudsdkvpc"
	"gitlab.ecloud.com/ecloud/ecloudsdkvpc/model"

	myconfig "github.com/DesistDaydream/ecloud-openapi/config"
)

// 使用AK&SK初始化账号Client
// @param accessKey
// @param secretKey
// @param poolId
// @return Client
func createClient(accessKey string, secretKey string, poolId string) *ecloudsdkvpc.Client {
	config := &config.Config{
		AccessKey: &accessKey,
		SecretKey: &secretKey,
		PoolId:    &poolId,
	}
	return ecloudsdkvpc.NewClient(config)
}

func main() {
	user := ""
	auth := myconfig.NewAuthInfo("config/my_config.yaml")

	ak := auth.AuthList[user].AccessKey
	sk := auth.AuthList[user].SecretKey
	client := createClient(ak, sk, "CIDC-RP-29")

	request := &model.CreateSecurityRuleRequest{}
	createSecurityRuleBody := &model.CreateSecurityRuleBody{}
	securityGroupId := "11bb4bcc-cd95-4bed-9def-d2df9d528768"
	createSecurityRuleBody.SecurityGroupId = &securityGroupId
	remoteType := "cidr"
	createSecurityRuleBody.RemoteType = (*model.CreateSecurityRuleBodyRemoteTypeEnum)(&remoteType)
	protocol := "ANY"
	createSecurityRuleBody.Protocol = (*model.CreateSecurityRuleBodyProtocolEnum)(&protocol)
	etherType := "IPv4"
	createSecurityRuleBody.EtherType = (*model.CreateSecurityRuleBodyEtherTypeEnum)(&etherType)
	description := "测试描述"
	createSecurityRuleBody.Description = &description
	remoteIpPrefix := "1.1.1.1/32"
	createSecurityRuleBody.RemoteIpPrefix = &remoteIpPrefix
	direction := "ingress"
	createSecurityRuleBody.Direction = (*model.CreateSecurityRuleBodyDirectionEnum)(&direction)
	request.CreateSecurityRuleBody = createSecurityRuleBody
	response, err := client.CreateSecurityRule(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}
