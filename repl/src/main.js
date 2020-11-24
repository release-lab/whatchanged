import { Button, Form, Input, Tabs } from "ant-design-vue";
import "ant-design-vue/dist/antd.css";
import { createApp } from "vue";
import App from "./App.vue";
import "./index.css";

const app = createApp(App);

app.use(Button);
app.use(Input);
app.use(Input.TextArea);
app.use(Form);
app.use(Tabs);
app.use(Tabs.TabPane);
app.use(Tabs.TabContent);

app.mount("#app");
