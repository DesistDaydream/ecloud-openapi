package fileparse

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestGetVpcIPaddrGroup(t *testing.T) {
	file := "/mnt/d/tmp/移动云安全组.xlsx"
	addrGroupName := "mgmt_tmp"
	srb, err := GetSecurityGroupRules(file, addrGroupName)
	if err != nil {
		logrus.Fatalf("%v", err)
	}
	fmt.Println(srb)
}
