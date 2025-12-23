<script setup lang="ts">
import { useAttrs } from 'vue'

defineOptions({
  inheritAttrs: false,
})

const props = defineProps<{
  id?: string
  label: string
  placeholder: string
  type?: string
  modelValue: string | number | null
  fullWidth?: boolean
  required?: boolean
}>()

// eslint-disable-next-line no-unused-vars
const emit = defineEmits<{ (e: 'update:modelValue', value: string | number | null): void }>()

const attrs = useAttrs()
const inputId = props.id ?? `input-${Math.random().toString(36).slice(2, 9)}`

function onInput(e: Event) {
  const target = e.target as HTMLInputElement | null
  emit('update:modelValue', target?.value ?? null)
}
</script>

<template>
  <div class="control field" :class="{ 'is-fullwidth': props.fullWidth }">
    <label class="label" :for="inputId">{{ props.label }}</label>
    <div class="field">
      <input
        v-bind="attrs"
        :id="inputId"
        :class="{ 'is-empty': !props.modelValue }"
        :placeholder="props.placeholder"
        :type="props.type"
        :value="props.modelValue"
        @input="onInput"
        :required="props.required"
      />
    </div>
  </div>
</template>

<style scoped lang="css">
.input {
  width: 100%;
  padding: 0.5rem;
  font-size: 1rem;
  border: 1px solid var(--border-color);
  border-radius: 4px;
  box-sizing: border-box;
  transition: border-color 0.2s ease-in-out;
  color: var(--font-color-dark);
  background-color: var(--font-color-light);
}
.input:focus {
  border-color: var(--green);
  outline: none;
}
.input::placeholder {
  color: var(--font-color-dark);
  opacity: 0.6;
}
input {
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  background-color: #ffffff;
  border: 1px solid var(--border-color);
  border-radius: 4px;
  padding: 0.5rem;
}
/* Added styles for date and time inputs */
input[type='date'],
input[type='time'] {
  appearance: none; /* Removes default browser styles */
  -webkit-appearance: none;
  -moz-appearance: none;
  min-height: 2.5rem; /* Ensure height matches text inputs */
}

input[type='date'].is-empty,
input[type='time'].is-empty {
  color: var(--font-color-dark);
  opacity: 0.6;
}

input[type='date']::placeholder,
input[type='time']::placeholder {
  color: var(--font-color-dark);
  opacity: 0.6;
}

.field.is-fullwidth {
  width: 100%;
}
</style>
