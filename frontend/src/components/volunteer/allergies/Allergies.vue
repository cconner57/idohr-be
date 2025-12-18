<script setup lang="ts">
import Toggle from '../../common/ui/Toggle.vue'
</script>

<template>
  <div class="field col-span-2">
    <label class="label">Allergies?</label>
    <Toggle leftLabel="No" rightLabel="Yes" />
  </div>
</template>

<style scoped lang="css">
.col-span-2 {
  grid-column: span 2;
}
</style>
