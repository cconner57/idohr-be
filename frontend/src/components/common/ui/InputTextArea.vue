<script setup lang="ts">
const props = defineProps<{
  label: string
  placeholder: string
}>()
</script>

<template>
  <div class="field col-span-2">
    <label class="label">{{ props.label }}</label>
    <div class="control">
      <textarea
        rows="3"
        :placeholder="props.placeholder"
        style="box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1); resize: none"
      />
    </div>
  </div>
</template>

<style scoped lang="css">
.col-span-2 {
  grid-column: span 2;
}
</style>
