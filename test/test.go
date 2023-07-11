// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package main

import (
	"fmt"

	"gitlab.ecloud.com/ecloud/ecloudsdkcore/config"
	"gitlab.ecloud.com/ecloud/ecloudsdkvpc"
	"gitlab.ecloud.com/ecloud/ecloudsdkvpc/model"
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
	ak := "fd75ddd5017e4708a7e4e41e940f718b"
	sk := "f81bb639b5424c70b94b4a26ee0aaca4"
	client := createClient(ak, sk, "CIDC-RP-29")
	request := &model.ListSecurityGroupRespRequest{}
	response, err := client.ListSecurityGroupResp(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}
