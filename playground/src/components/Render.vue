<template>
  <a-tabs class="render">
    <a-tab-pane key="1" tab="Result">
      <Markdown
        :content="content"
        :loading="loading"
        class="markdown"
        style="padding-left: 30px"
      />
    </a-tab-pane>
    <a-tab-pane key="2" tab="Source">
      <CodeMirror :content="content" :loading="loading" :readonly="true" />
    </a-tab-pane>
  </a-tabs>
</template>

<script setup>
import { toRefs } from "vue";
import Markdown from "./Markdown.vue";
import CodeMirror from "./CodeMirror.vue";

const props = defineProps({
  content: String,
  loading: Boolean,
});

const { content, loading } = toRefs(props);
</script>

<style lang="scss">
.render {
  .ant-tabs-bar {
    margin-bottom: 0;
  }
  .markdown {
    height: calc(100vh - 60px - 46px);
    overflow: scroll;
  }
}
</style>
