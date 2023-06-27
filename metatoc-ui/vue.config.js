const { defineConfig } = require("@vue/cli-service");
module.exports = defineConfig({
  publicPath: "./",
  transpileDependencies: true,
  devServer: {
    proxy: {
      "/api/sendMsg": {
        target: "http://127.0.0.1:3001",
        changOrigin: true,
        pathRewrite: {
          "^/api": "",
        },
      },
      "/api/block": {
        target: "http://172.22.50.202:2929",
        // target: "http://172.22.50.211:5000",
        changOrigin: true,
        pathRewrite: {
          "^/api/block": "",
        },
      },
      "/api/metatoc-service/v1/": {
        // target: "http://172.16.31.118:8080",
        target: "http://172.22.50.211:8088",
        changOrigin: true,
        pathRewrite: {
          "^/api": "",
        },
      }
      // "/api/block/paths": {
      //   target: "http://172.22.50.202:2929",
      //   changOrigin: true,
      //   pathRewrite: {
      //     "^/api/block": "",
      //   },
      // },
    },
  },
  chainWebpack: (config) => {
    config.plugin("html").tap((args) => {
      args[0].title = "MetaTOC Demo";
      return args;
    });
  },
});
