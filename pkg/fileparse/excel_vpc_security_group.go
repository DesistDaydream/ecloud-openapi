package fileparse

import (
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
	"gitlab.ecloud.com/ecloud/ecloudsdkvpc/model"
)

type SecurityRuleBody struct {
	Protocol       *model.CreateSecurityRuleBodyProtocolEnum
	RemoteType     *model.CreateSecurityRuleBodyRemoteTypeEnum
	EtherType      *model.CreateSecurityRuleBodyEtherTypeEnum
	Description    string
	Direction      *model.CreateSecurityRuleBodyDirectionEnum
	RemoteIpPrefix string
}

func GetSecurityGroupRules(file string, addrGroupName string) (srb []SecurityRuleBody, err error) {
	f, err := excelize.OpenFile(file)
	if err != nil {
		logrus.Errorf("打开文件错误，原因: %v", err)
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			logrus.Errorln(err)
			return
		}
	}()

	// 逐行读取Excel文件
	rows, err := f.GetRows(addrGroupName)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"sheet": addrGroupName,
		}).Errorf("读取中sheet页异常: %v", err)
		return nil, err
	}

	remoteType := "cidr"

	for i := 1; i < len(rows); i++ {
		row := rows[i]
		srb = append(srb, SecurityRuleBody{
			RemoteType:     (*model.CreateSecurityRuleBodyRemoteTypeEnum)(&remoteType),
			Protocol:       (*model.CreateSecurityRuleBodyProtocolEnum)(&row[0]),
			EtherType:      (*model.CreateSecurityRuleBodyEtherTypeEnum)(&row[3]),
			Description:    row[5],
			RemoteIpPrefix: row[4],
			Direction:      (*model.CreateSecurityRuleBodyDirectionEnum)(&row[2]),
		})
	}

	return srb, nil
}
