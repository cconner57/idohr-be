<script setup lang="ts">
import InputField from '../../common/ui/InputField.vue'
import type { SurrenderFormState } from '../../../models/common'

const { formState } = defineProps<{
  formState: SurrenderFormState
}>()
</script>

<template>
  <div>
    <h5>Aggressive Behavior</h5>
    <InputField
      label="Has the cat ever attacked or bit a person?"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField label="If yes, please explain" placeholder="Email" :modelValue="formState.email" />
    <InputField
      label="Has the cat ever attacked or bit a cat?"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField label="If yes, please explain" placeholder="Email" :modelValue="formState.email" />
  </div>
</template>

<style scoped lang="css"></style>
