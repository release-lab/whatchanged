<template>
  <div class="my-editor" v-bind="$attrs">
    <textarea ref="input"></textarea>
  </div>
</template>

<script>
import { defineComponent } from "vue";
import CodeMirror from "https://cdn.jsdelivr.net/npm/codemirror/src/codemirror.js";
import "./CodeMirror.css";

export default defineComponent({
  props: {
    content: { type: String },
    readonly: { type: Boolean, default: () => false },
  },
  data() {
    return {};
  },
  methods: {
    update(value) {
      this.editor.doc.setValue(value);
    },
  },
  mounted() {
    this.editor = CodeMirror.fromTextArea(this.$refs.input, {
      lineNumbers: true,
      readonly: !!this.readonly,
    });
    if (typeof this.content === "string") {
      this.update(this.content);
    }
    this.editor.on("change", (instance, change) => {
      const val = this.editor.doc.getValue();
      this.$emit("update:content", val);
    });
  },
});
</script>

<style lang="scss">
.my-editor {
  width: 100%;
  .CodeMirror {
    height: calc(100vh - 60px - 46px);
  }
}

.my-editor2 {
  width: 100%;
  .CodeMirror {
    height: calc(100vh - 60px);
  }
}
</style>
