FROM centos:centos7.9.2009 AS build
RUN mkdir -p /data/node
WORKDIR /data/node
RUN curl https://nodejs.org/dist/v16.18.0/node-v16.18.0-linux-x64.tar.xz > node-v16.18.0-linux-x64.tar.xz
RUN xz -d node-v16.18.0-linux-x64.tar.xz && tar -xvf node-v16.18.0-linux-x64.tar && rm -f node-v16.18.0-linux-x64.tar
RUN mv node-v16.18.0-linux-x64 /usr/local/nodejs && ln -s /usr/local/nodejs/bin/node /usr/local/bin/ && ln -s /usr/local/nodejs/bin/npm /usr/local/bin/
RUN npm config set registry https://registry.npm.taobao.org --global
COPY . /data/node/
WORKDIR /data/node/
ARG tag
RUN npm install && npm run build

FROM nginx:1.21.6
COPY --from=build /data/node/dist/ /usr/share/nginx/html
RUN mkdir -p /usr/share/nginx/html/src/assets
COPY nginx.template /etc/nginx/conf.d
EXPOSE 80
WORKDIR /etc/nginx/conf.d
ENTRYPOINT envsubst '$chatgpt_service_proxy_pass $blockchain_service_proxy_pass $metatoc_service_proxy_pass' < nginx.template > default.conf && cat default.conf && nginx -g 'daemon off;'
