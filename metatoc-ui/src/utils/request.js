import axios from "axios";

const service = axios.create({
  baseURL: "/api",
  timeout: 30000,
});

service.interceptors.request.use(
  (config) => {
    console.log("service.interceptors.request");
    console.log(config);
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

service.interceptors.response.use(
  (response) => {
    console.log("service.interceptors.response");
    console.log(response);
    if (response.status == 200) {
      return Promise.resolve(response);
    } else {
      return Promise.reject(response);
    }
  },
  (error) => {
    return Promise.reject(error);
  }
);

/**
 * GET请求
 */
function $get(url, params, headers) {
  return service.get(url, {
    params,
    headers,
  });
}

/**
 * POST请求
 */
function $post(url, data, headers) {
  return service.post(url, data, { headers });
}

/**
 * PUT请求
 */
function $put(url, data, headers) {
  return service.put(url, data, { headers });
}

export { $get, $post, $put };
