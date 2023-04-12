# coding: utf-8

"""
    Description: python client
    Generated by: ecloud-sdk
"""

import sys
from dataclasses import dataclass
from typing import Optional, List

from ecloudsdkcore.config.config import Config
from ecloudsdkvpc.v1.model import *
from ecloudsdkvpc.v1.client import Client

import yaml


@dataclass
class Content:
    created_time: str
    default_rule: bool
    description: None
    direction: str
    ether_type: str
    id: str
    protocol: str
    secgroup_id: str
    status: None
    aim_sgid: Optional[str] = None
    max_port_range: Optional[int] = None
    min_port_range: Optional[int] = None
    remote_ip_prefix: Optional[str] = None


@dataclass
class Body:
    content: List[Content]
    empty: bool
    total: int


@dataclass
class ListSecurityGroupRuleResp:
    """ListSecurityGroupRuleResp"""

    body: Body
    request_id: str
    state: str


def create_client(access_key: str, access_secret: str, pool_id: str) -> Client:
    """
    使用AK&SK初始化账号Client
    @param access_key:
    @param access_secret:
    @param pool_id:
    @return: Client
    @throws Exception
    """
    config = Config(access_key=access_key, access_secret=access_secret, pool_id=pool_id)
    return Client(config)


@dataclass
class VPCHandler:
    @staticmethod
    def ListSecurityGroupRule(args: List[str], client: Client) -> None:
        request = ListSecurityGroupRuleRequest()
        list_security_group_rule_query = ListSecurityGroupRuleQuery()
        list_security_group_rule_query.security_group_id = (
            "2d64eeb1-d8f0-446a-92dc-79543fb7a027"
        )
        request.list_security_group_rule_query = list_security_group_rule_query
        resp: ListSecurityGroupRuleResp = client.list_security_group_rule(request)
        print("当前共有 {} 条规则".format(len(resp.body.content)))

    def CreateSecurityRule(args: List[str], client: Client) -> None:
        request = CreateSecurityRuleRequest()
        create_security_rule_body = CreateSecurityRuleBody()
        create_security_rule_body.security_group_id = (
            "2d64eeb1-d8f0-446a-92dc-79543fb7a027"
        )
        create_security_rule_body.remote_type = "cidr"
        create_security_rule_body.protocol = "TCP"
        create_security_rule_body.min_port_range = 22
        create_security_rule_body.ether_type = "IPv4"
        create_security_rule_body.remote_ip_prefix = "1.1.1.1/32"
        create_security_rule_body.direction = "ingress"
        create_security_rule_body.max_port_range = 22
        request.create_security_rule_body = create_security_rule_body
        resp = client.create_security_rule(request)
        print(resp)


if __name__ == "__main__":
    with open("pkg/config/my_config.yaml", "r") as file:
        yaml_data = yaml.safe_load(file)

    client = create_client(
        yaml_data["ak"],
        yaml_data["sk"],
        "CIDC-RP-25",
    )

    VPCHandler.ListSecurityGroupRule(sys.argv[1:], client)
    VPCHandler.CreateSecurityRule(sys.argv[1:], client)
