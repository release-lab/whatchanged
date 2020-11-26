<template>
  <div v-bind="$attrs">
    <div v-html="html"></div>
  </div>
</template>

<script setup>
import { computed, defineProps } from "vue";
import marked from "marked";

const props = defineProps({ content: String, default: () => "" });

const html = computed(() => marked(props.content));
</script>
