<template>
  <div v-bind="$attrs">
    <div v-html="html"></div>
  </div>
</template>

<script>
import { defineComponent } from "vue";
import marked from "marked";

export default defineComponent({
  props: {
    content: { type: String, default: () => "" },
  },
  data() {
    return {};
  },
  computed: {
    html() {
      return marked(this.content);
    },
  },
});
</script>
