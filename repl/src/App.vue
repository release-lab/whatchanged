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
            @update:content="template = $event"
            :readonly="true"
          />
        </div>
      </div>
      <div class="right">
        <Render :content="content"></Render>
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent } from "vue";
import Render from "./components/Render.vue";
import CodeMirror from "./components/CodeMirror.vue";
import TEMPLATE_DEFAULT from "./template/default";

export default defineComponent({
  name: "App",
  components: {
    Render,
    CodeMirror,
  },
  data() {
    return {
      loading: false,
      template: TEMPLATE_DEFAULT,
      content: "",
      form: {
        username: "axetroy",
        repo: "vscode-deno",
        version: "HEAD~",
      },
    };
  },
  methods: {
    onSubmit() {
      const template = encodeURIComponent(this.template);

      this.loading = true;
      fetch(
        `${import.meta.env.VITE_API_HOST}/?username=${
          this.form.username || ""
        }&repo=${this.form.repo || ""}&version=${
          this.form.version || ""
        }&template=${template || ""}`
      )
        .then((res) => {
          return res.text();
        })
        .then((markdown) => {
          this.content = markdown;
        })
        .finally(() => {
          this.loading = false;
        });
    },
  },
});
</script>

<style lang="scss" scoped>
$height: 60px;
.toolbar {
  height: $height;
  background: yellow;
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
