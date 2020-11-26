<template>
  <div class="my-editor" v-bind="$attrs">
    <textarea ref="input"></textarea>
  </div>
</template>

<script setup>
import { onMounted, ref, defineEmit, defineProps } from "vue";
import CodeMirror from "https://cdn.jsdelivr.net/npm/codemirror/src/codemirror.js";

const { content, readonly } = defineProps({
  content: { type: String },
  readonly: { type: Boolean, default: () => false },
});

const emit = defineEmit(["update:content"]);

let editor;

const input = ref(null);

onMounted(() => {
  editor = CodeMirror.fromTextArea(input.value, {
    lineNumbers: true,
    readonly: !!readonly,
  });

  if (typeof content === "string") {
    update(content);
  }
  editor.on("change", (instance, change) => {
    emit("update:content", editor.doc.getValue());
  });
});

function update(value) {
  editor.doc.setValue(value);
}
</script>

<style lang="scss">
@import "./CodeMirror.css";
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
