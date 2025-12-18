<script setup lang="ts">
const props = defineProps<{
  label: string
  placeholder: string
}>()
</script>

<template>
  <div class="field col-span-2">
    <label class="label">{{ props.label }}</label>
    <div class="control"><textarea rows="3" :placeholder="props.placeholder" /></div>
  </div>
</template>

<style scoped lang="css">
.col-span-2 {
  grid-column: span 2;
}
</style>
