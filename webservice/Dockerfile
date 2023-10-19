# Copyright 2022-2023 The VNET Project Authors. All Rights Reserved.

# SPDX-License-Identifier: MIT

# FROM python:3.8.13-alpine3.16
FROM python:3.8.13

WORKDIR /usr/src/app

# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
# RUN apk add gcc linux-headers libc-dev libffi-dev musl-dev python3-dev openssl-dev cargo

COPY . .
RUN pip install --upgrade pip -i https://mirrors.aliyun.com/pypi/simple/
RUN pip install --no-cache-dir -r requirements.txt -i https://mirrors.aliyun.com/pypi/simple/

RUN flask --app flaskr init-db

EXPOSE 5000
# CMD ["flask", "run", "-h", "0.0.0.0"]
CMD ["waitress-serve", "--port", "5000", "--call", "flaskr:create_app"]
