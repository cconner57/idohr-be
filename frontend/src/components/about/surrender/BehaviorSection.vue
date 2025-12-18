<script setup lang="ts">
import InputField from '../../common/ui/InputField.vue'
import type { SurrenderFormState } from '../../../models/common'

const { formState } = defineProps<{
  formState: SurrenderFormState
}>()
</script>

<template>
  <div>
    <h5>Behavior</h5>
    <InputField
      label="How does the cat usually behave towards people they do know?"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="How does the cat usually behave towards animals they do not know?"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="How does the cat usually behave towards animals they know"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="Comments on answers above"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="How does the cat usually react when an unfamiliar person approaches or enters the yard or house?"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="Is the cat housetrained?"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="Where does the cat spend the majority of their time?"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="How many hours a day is the cat left alone without a human?"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="When left alone, is he/she..."
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="When left alone, does the cat usually show any of the following behaviors?"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="When the cat plays, do they typically..."
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="What toys does the cat like?"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="What games does the cat like?"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="Is the cat scared of anything?"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField label="If yes, please explain" placeholder="Email" :modelValue="formState.email" />
    <InputField
      label="Please tell us the cats bad habits"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="Is the cat allowed on furniture?"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="Where does the pet usually sleep at overnight?"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="Describe the pets reaction if a person or other cat gets near or attempts to take food he/she is eating"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="Describe the pets reaction if a person or other cat attempts to take a toy they are playing with"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField
      label="Does the cat have problems riding in cars?"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField label="If yes, please explain" placeholder="Email" :modelValue="formState.email" />
    <InputField
      label="Has the pet escaped your property 2 or more times in the last 6 months?"
      placeholder="Email"
      :modelValue="formState.email"
    />
    <InputField label="If yes, please explain" placeholder="Email" :modelValue="formState.email" />
  </div>
</template>

<style scoped lang="css"></style>
