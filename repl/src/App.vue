<template>
  <div>
    <div class="toolbar">
      <a-form
        layout="inline"
        :model="form"
        @submit="onSubmit"
        @submit.native.prevent
      >
        <a-form-item label="Username">
          <a-input v-model:value="form.username"> </a-input>
        </a-form-item>
        <a-form-item label="Repo">
          <a-input v-model:value="form.repo"> </a-input>
        </a-form-item>
        <a-form-item label="Version">
          <a-input v-model:value="form.version"> </a-input>
        </a-form-item>
        <a-form-item>
          <a-button type="primary" html-type="submit" :loading="loading">
            Generate
          </a-button>
        </a-form-item>
      </a-form>
    </div>
    <div class="container">
      <div class="left">
        <div style="display: flex; height: 100%; padding-left: 10px">
          <CodeMirror
            class="my-editor2"
            :content.sync="template"
            :readonly="true"
            @update:content="template = $event"
          />
        </div>
      </div>
      <div class="right">
        <Render :content="content" :loading="loading"></Render>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { notification } from "ant-design-vue";
import Render from "./components/Render.vue";
import CodeMirror from "./components/CodeMirror.vue";
import TEMPLATE_DEFAULT from "./template/default";

const loading = ref(false);
const template = ref(TEMPLATE_DEFAULT);
const form = ref({
  username: "denoland",
  repo: "deno",
  version: "HEAD~",
});

const content = ref("");

function onSubmit() {
  const tpl = encodeURIComponent(template.value);

  loading.value = true;

  const key = Math.random() + "";

  notification.info({
    key: key,
    message: "Generating",
    description: "This may take a few minutes...",
  });

  fetch(
    `${import.meta.env.VITE_API_HOST}/?username=${
      form.value.username || ""
    }&repo=${form.value.repo || ""}&version=${
      form.value.version || ""
    }&template=${tpl || ""}`
  )
    .then((res) => res.text())
    .then((markdown) => {
      notification.close(key);
      content.value = markdown;
    })
    .catch((err) => {
      notification.close(key);
      notification.error({
        message: "Error",
        description: err.message,
      });
    })
    .finally(() => {
      loading.value = false;
    });
}
</script>

<style lang="scss" scoped>
$height: 60px;
.toolbar {
  height: $height;
  background: #eeda7c;
  display: flex;
  align-items: center;
  padding: 0 10px;
}

.container {
  display: flex;
  height: calc(100vh - #{$height});
  width: 100vw;

  .left {
    // background-color: green;
    flex: 1;
    overflow: scroll;
    border-right: 1px solid #e2e2;
  }

  .right {
    // background-color: red;
    flex: 1;
    overflow: hidden;
  }
}
</style>
