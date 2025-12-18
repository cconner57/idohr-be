<script setup lang="ts">
import InputField from '../../common/ui/InputField.vue'
import type { SurrenderFormState } from '../../../models/common'

const { formState } = defineProps<{
  formState: SurrenderFormState
}>()
</script>

<template>
  <div>
    <h5>Medical History</h5>
    <InputField
      label="Please list all veterniarians that have seen the cat, including address and number"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="Does the cat see a veterinarian at least once a year?"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="Is the cat spayed/neutered?"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="What vaccinations has the cat received?"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="Are vaccinations current?"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="Has the cat been tested for heartworm?"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="If yes, what were the test results"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="Is the pet taking heartworm preventative?"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="If yes, when is the next dosage due?"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="Is the pet microchipped?"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="If yes, please provide microchip number"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="Check if the cat has ever shown any of the following behaviors when handled by a vet or groomer"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="Does the cat have to be muzzled at the veterinarians?"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="Does the pet have any past or present health problems"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField label="If yes, please describe" placeholder="Email" :modelValue="formState.email" />
    <InputField
      label="Is the cat currently taking any medications?"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="If yes, please list medications"
      placeholder="Email"
      :modelValue="formState.email"
    />
  </div>
</template>

<style scoped lang="css"></style>
