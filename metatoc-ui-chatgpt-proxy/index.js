import { ChatGPTAPI } from "chatgpt";
import express from "express";
import bodyParser from "body-parser";
import proxy from "https-proxy-agent";
import nodeFetch from "node-fetch";

const apiKey = process.env.API_KEY;
const proxyUrl = process.env.PROXY_URL;

const app = express();
const port = 3001;
app.use(bodyParser.json()); // 添加json解析
app.use(bodyParser.urlencoded({ extended: false }));
//设置跨域访问
// app.all('*', function(req, res, next) {
//   res.header("Access-Control-Allow-Origin", "*");
//   res.header("Content-Type", "application/json;charset=utf-8");
//   next();
// });

app.all("*", function (req, res, next) {
  res.header("Access-Control-Allow-Origin", req.headers.origin);
  res.header("Access-Control-Allow-Credentials", "true");
  res.header("Access-Control-Allow-Headers", "*");
  res.header("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS");
  res.header("X-Powered-By", " 3.2.1");
  if (req.method == "OPTIONS") res.send(200); /*让options请求快速返回*/
  else next();
});

const chatGPTApi = new ChatGPTAPI({
  apiKey,
  fetch: (url, options = {}) => {
    const defaultOptions = {
      agent: proxy(proxyUrl),
    };

    const mergedOptions = {
      ...defaultOptions,
      ...options,
    };

    return nodeFetch(url, mergedOptions);
  },
});

app.post("/sendMsg", async (req, res) => {
  console.log("req.body==>", req.body);
  try {
    let result = await chatGPTApi.sendMessage(req.body.msg);
    console.log("res==>", result);
    res.send(`${result.text}`);
  } catch (e) {
    console.error(e);
  }
});

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
});
