import { createApp } from "vue";
import Antd from "ant-design-vue";
import App from "./App.vue";
import "ant-design-vue/dist/antd.min.css"; // 更新此处

const app = createApp(App);

app.use(Antd);

app.mount("#app");
