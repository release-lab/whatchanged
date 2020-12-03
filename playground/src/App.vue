<template>
  <div style="position: relative">
    <img
      src="/github.svg"
      style="position: fixed; right: 0; top: 0"
    />
    <div class="toolbar">
      <!-- <img :src="require('@/assets/logo.png')" /> -->
      <a-form
        layout="inline"
        :model="form"
        @submit="onSubmit"
        @submit.native.prevent
      >
        <a-form-item label="Username">
          <a-input v-model:value="form.username.value"> </a-input>
        </a-form-item>
        <a-form-item label="Repo">
          <a-input v-model:value="form.repo.value"> </a-input>
        </a-form-item>
        <a-form-item label="Version">
          <a-input v-model:value="form.version.value"> </a-input>
        </a-form-item>
        <a-form-item>
          <a-button type="primary" html-type="submit" :loading="loading">
            Generate
          </a-button>
          <a-button type="default" @click="copyURL" style="margin-left: 20px">
            <!-- <MessageOutlined :style="{ fontSize: '16px', color: '#08c' }" /> -->
            Share
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
import { ref, reactive, onMounted, watch, toRefs } from "vue";
import copy from "https://esm.sh/copy-to-clipboard";
import { message, notification } from "ant-design-vue";
import Render from "./components/Render.vue";
import CodeMirror from "./components/CodeMirror.vue";
import TEMPLATE_DEFAULT from "./template/default";

const loading = ref(false);
let template = ref(TEMPLATE_DEFAULT);

const formReactive = reactive({
  username: "axetroy",
  repo: "whatchanged",
  version: "HEAD~",
});

const form = toRefs(formReactive);

function watchAndUpdateQuery(field) {
  watch(
    () => formReactive[field],
    (val) => {
      syncQueryField(field, val);
    }
  );
}

function syncQueryField(field, val) {
  const url = new URL(window.location.href);
  if (url.searchParams.has(field)) {
    url.searchParams.set(field, val);
  } else {
    url.searchParams.append(field, val);
  }
  window.history.pushState(null, null, url);
}

watchAndUpdateQuery("username");
watchAndUpdateQuery("repo");
watchAndUpdateQuery("version");

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
      formReactive.username || ""
    }&repo=${formReactive.repo || ""}&version=${
      formReactive.version || ""
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

function copyURL() {
  const url = new URL(window.location.href);

  url.searchParams.append("tpl", template.value);

  copy(url.href);

  message.info("URL have been copy to clipboard.");
}

onMounted(() => {
  notification.warn({
    message: "IMPORTANT",
    description:
      "Hi ❤️ We are useing the free resources for backend and there is a limit to the memory size. so, it may fail to generate for large projects.",
    duration: 30,
  });

  const url = new URL(window.location.href);

  let params = 0;

  if (url.searchParams.has("username")) {
    formReactive.username = url.searchParams.get("username");
    params++;
  }

  if (url.searchParams.has("repo")) {
    formReactive.repo = url.searchParams.get("repo");
    params++;
  }

  if (url.searchParams.has("version")) {
    formReactive.version = url.searchParams.get("version");
  }

  if (url.searchParams.has("tpl")) {
    params++;
    template.value = url.searchParams.get("tpl");
  }

  if (params === 3 && formReactive.username && formReactive.repo) {
    onSubmit();
  }
});
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
