FROM python:3.8.13-alpine3.16

WORKDIR /usr/src/app

COPY . .
RUN pip install --no-cache-dir -r requirements.txt -i https://mirrors.aliyun.com/pypi/simple/

ENV FLASK_APP=flaskr
ENV FLASK_ENV=development
RUN flask init-db

CMD ["flask", "run", "-h", "0.0.0.0"]
