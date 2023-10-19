# Copyright 2022-2023 The VNET Project Authors. All Rights Reserved.

# SPDX-License-Identifier: MIT

FROM ubuntu:20.04

RUN apt-get update && apt-get install -y wget unzip

WORKDIR /opt/diem
RUN wget https://github.com/violas-core/Violas/releases/download/v0.11.0/Violas.v0.11.0.zip && unzip Violas.v0.11.0.zip
COPY ./gen_config.sh /opt/diem/gen_config.sh
COPY ./genesis.yaml /opt/diem/genesis.yaml

ENTRYPOINT ["bash", "./gen_config.sh"]
