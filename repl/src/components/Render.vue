<template>
  <a-tabs class="render">
    <a-tab-pane key="1" tab="Result">
      <Markdown
        :content="content"
        class="markdown"
        style="padding-left: 30px"
      />
    </a-tab-pane>
    <a-tab-pane key="2" tab="Source">
      <CodeMirror ref="editor" :content="content" :readonly="true" />
    </a-tab-pane>
  </a-tabs>
</template>

<script setup>
import { ref, watch, defineProps } from "vue";
import Markdown from "./Markdown.vue";
import CodeMirror from "./CodeMirror.vue";

const editor = ref(null);

const { content } = defineProps({
  content: { type: String },
});

watch(content, () => {
  editor.update(val);
});
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
