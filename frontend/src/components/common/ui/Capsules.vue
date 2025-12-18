<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  label?: string | null
  color?: string
  size?: 'sm' | 'md' | 'lg'
}>()

const sizeClass = computed(() => {
  if (props.size === 'sm') return 'capsule--sm'
  if (props.size === 'lg') return 'capsule--lg'
  return 'capsule--md'
})
</script>

<template>
  <span class="capsule" :class="sizeClass" :style="{ '--capsule-bg': props.color || '' }">
    <slot>{{ props.label }}</slot>
  </span>
</template>

<style scoped lang="css">
.capsule {
  background: var(--capsule-bg, #e8f6f4);
  color: #0f5b52;
  border-radius: 999px;
  padding: 0.375rem 0.625rem;
  border: 1px solid rgba(15, 91, 82, 0.12);
  font-weight: 600;
  font-size: 0.875rem;
  line-height: 1;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 1px 0 rgba(255, 255, 255, 0.6) inset;
  white-space: nowrap;
  text-overflow: ellipsis;
  max-width: 12rem;
  overflow: hidden;
  text-transform: capitalize;
}

.capsule--sm {
  padding: 0.25rem 0.5rem;
  font-size: 0.75rem;
}

.capsule--md {
  padding: 0.375rem 0.625rem;
  font-size: 0.875rem;
}

.capsule--lg {
  padding: 0.5rem 0.75rem;
  font-size: 1rem;
}
</style>
