<template>
  <a-spin :spinning="loading">
    <div v-bind="$attrs">
      <div v-html="html"></div>
    </div>
  </a-spin>
</template>

<script setup>
import { computed, toRefs } from "vue";
import marked from "marked";

const props = defineProps({
  content: String,
  loading: Boolean,
});

const { content, loading } = toRefs(props);

const html = computed(() => marked(content.value));
</script>
