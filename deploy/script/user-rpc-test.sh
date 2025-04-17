#!/bin/bash
reso_addr='crpi-l7tvpukszv54yebz.cn-hangzhou.personal.cr.aliyuncs.com/binghu-easy-chat/user-rpc-dev'
tag='latest'

container_name="binghu-easy-chat-user-rpc-dev"

docker stop ${container_name}

docker rm ${container_name}

docker rmi ${reso_addr}:${tag}

docker pull ${reso_addr}:${tag}


# 如果需要指定配置文件的
# docker run -p 10001:8080 --network imooc_easy-im -v /easy-im/config/user-rpc:/user/conf/ --name=${container_name} -d ${reso_addr}:${tag}
docker run -p 10000:10000  --name=${container_name} -d ${reso_addr}:${tag}