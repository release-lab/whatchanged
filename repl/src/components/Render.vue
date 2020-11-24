<template>
  <a-tabs class="render">
    <a-tab-pane key="1" tab="Result">
      <Markdown :content="content" class="markdown" />
    </a-tab-pane>
    <a-tab-pane key="2" tab="Source">
      <CodeMirror ref="CodeMirror" :content="content" :readonly="true" />
    </a-tab-pane>
  </a-tabs>
</template>

<script>
import { defineComponent } from "vue";
import Markdown from "./Markdown.vue";
import CodeMirror from "./CodeMirror.vue";

export default defineComponent({
  components: {
    Markdown,
    CodeMirror,
  },
  props: {
    content: { type: String },
  },
  watch: {
    content(val) {
      this.$refs.CodeMirror.update(val);
    },
  },
  data() {
    return {};
  },
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
