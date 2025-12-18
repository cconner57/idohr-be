<script setup lang="ts">
import InputField from '../../common/ui/InputField.vue'
import type { SurrenderFormState } from '../../../models/common'

const { formState } = defineProps<{
  formState: SurrenderFormState
}>()
</script>

<template>
  <div>
    <h5>Feeding</h5>
    <InputField
      label="What type of food does the cat eat?"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="How many times a day is pet fed?"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="How much is fed per feeding?"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="Does the pet get treats?"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="If yes, what treats does the pet like?"
      placeholder="Email"
      :modelValue="formState.email"
    />
  </div>
</template>

<style scoped lang="css"></style>
