import sys
import yaml


def LoadConfig(configFile):
    try:
        with open(configFile, "r") as file:
            configData = yaml.safe_load(file)
    except FileNotFoundError:
        print("配置文件不存在")
        sys.exit(1)

    return configData
