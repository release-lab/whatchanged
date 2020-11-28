import { Button, Form, Input, Spin, Tabs } from "ant-design-vue";
import "ant-design-vue/dist/antd.css";
// import { MessageOutlined } from "ant-design-vue/icons-vue";
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
app.use(Spin);
// app.use(MessageOutlined);

app.mount("#app");
