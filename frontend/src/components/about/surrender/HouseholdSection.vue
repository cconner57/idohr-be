<script setup lang="ts">
import InputField from '../../common/ui/InputField.vue'
import type { SurrenderFormState } from '../../../models/common'
import InputTextArea from '../../common/ui/InputTextArea.vue'

const { formState } = defineProps<{
  formState: SurrenderFormState
}>()
</script>

<template>
  <div class="household-section">
    <h5>Cat & Household Information</h5>
    <section>
      <InputField label="First Name" placeholder="First Name" :modelValue="formState.firstName" />
      <InputField label="Last Name" placeholder="Last Name" :modelValue="formState.lastName" />
    </section>
    <section>
      <InputField
        label="Phone Number"
        placeholder="Phone Number"
        :modelValue="formState.phoneNumber"
      />
      <InputField label="Email" placeholder="Email" :modelValue="formState.email" />
    </section>
    <section>
      <InputField
        label="Street Address"
        placeholder="Street Address"
        :modelValue="formState.streetAddress"
      />
      <InputField label="City" placeholder="City" :modelValue="formState.city" />
    </section>
    <section>
      <InputField label="State" placeholder="State" :modelValue="formState.state" />
      <InputField label="Zip Code" placeholder="Zip Code" :modelValue="formState.zipCode" />
    </section>
    <section>
      <InputTextArea
        label="When do you need to surrender your cat"
        placeholder="When do you need to surrender your cat"
        :modelValue="formState.whenToSurrenderCat"
      />
      <InputField label="Cat's Name" placeholder="Cat's Name" :modelValue="formState.catName" />
      <InputField label="Age" placeholder="Age" :modelValue="formState.catAge" />
      <fieldset class="field col-span-2" aria-labelledby="cat-sex-legend">
        <legend id="cat-sex-legend" class="label">Sex</legend>
        <fieldset class="chips" aria-labelledby="cat-sex-legend">
          <label class="chip"><input type="checkbox" /> <span>Male</span></label>
          <label class="chip"><input type="checkbox" /> <span>Female</span></label>
          <label class="chip"><input type="checkbox" /> <span>Unknown</span></label>
        </fieldset>
      </fieldset>
    </section>
    <section>
      <InputTextArea
        label="How long have you had your cat?"
        placeholder="How long have you had your cat?"
        :modelValue="formState.catOwnershipDuration"
      />
      <InputTextArea
        label="Where did you get your cat?"
        placeholder="Where did you get your cat?"
        :modelValue="formState.catLocationFound"
      />
      <InputTextArea
        label="Why are you surrendering your cat?"
        placeholder="Why are you surrendering your cat?"
        :modelValue="formState.catWhySurrendered"
      />
    </section>
    <div>
      <h5>Including yourself, how many people of the following ages live in your home?</h5>
      <section>
        <InputField label="Female - Ages 0-3" placeholder="" :modelValue="formState.email" />
        <InputField label="Male - Ages 0-3" placeholder="" :modelValue="formState.email" />
        <InputField label="Female - Ages 4-9" placeholder="" :modelValue="formState.email" />
        <InputField label="Male - Ages 4-9" placeholder="" :modelValue="formState.email" />
        <InputField label="Female - Ages 10-17" placeholder="" :modelValue="formState.email" />
        <InputField label="Male - Ages 10-17" placeholder="" :modelValue="formState.email" />
        <InputField label="Female - Ages 18-29" placeholder="" :modelValue="formState.email" />
        <InputField label="Male - Ages 18-29" placeholder="" :modelValue="formState.email" />
        <InputField label="Female - Ages 30-59" placeholder="" :modelValue="formState.email" />
        <InputField label="Male - Ages 30-59" placeholder="" :modelValue="formState.email" />
        <InputField label="Female - Ages 60+" placeholder="" :modelValue="formState.email" />
        <InputField label="Male - Ages 60+" placeholder="" :modelValue="formState.email" />
      </section>
    </div>
    <section>
      <fieldset class="field col-span-2" aria-labelledby="other-animals-legend">
        <legend id="other-animals-legend" class="label">
          What other animals did the cat live with?
        </legend>
        <fieldset class="chips" aria-labelledby="other-animals-legend">
          <label class="chip"><input type="checkbox" /> <span>Dogs</span></label>
          <label class="chip"><input type="checkbox" /> <span>Cats</span></label>
          <label class="chip"><input type="checkbox" /> <span>Other</span></label>
          <label class="chip"><input type="checkbox" /> <span>No other animals</span></label>
        </fieldset>
      </fieldset>
    </section>
  </div>
</template>

<style scoped lang="css">
.household-section {
  display: flex;
  flex-direction: column;
  gap: 16px;

  section {
    display: grid;
    grid-template-columns: 1fr 1fr;
    row-gap: 0px;
    column-gap: 16px;

    fieldset.field {
      border: 0;
      padding: 0;
      margin: 0;
      @media (max-width: 440px) {
        .col-span-2 {
          grid-column: span 1;
        }
      }
    }

    .field {
      @media (max-width: 440px) {
        margin-bottom: 8px;
      }
    }
    .col-span-2 {
      grid-column: span 2;
      @media (max-width: 440px) {
        grid-column: span 1;
      }
    }

    .label {
      margin-bottom: 8px;
    }
    .chips {
      display: flex;
      flex-wrap: wrap;
      gap: 8px;
      border: none;
      @media (max-width: 440px) {
        margin-top: 8px;
        flex-direction: column;
        text-align: center;
        gap: 8px;
      }
    }
    .chip {
      position: relative;
      display: inline-flex;
      align-items: center;
      gap: 8px;
      padding: 8px 12px;
      border-radius: 999px;
      border: 1px solid #e7ebf0;
      background-color: #fff;
      cursor: pointer;
      user-select: none;
      font-size: 1rem;
      transition:
        background 0.2s,
        border-color 0.2s,
        box-shadow 0.2s;

      span {
        font-weight: 600;
        color: var(--text-900);
        line-height: 1.5;
      }

      &:hover {
        border-color: #d7e2f2;
        background: #f2f7ff;
      }
    }
    .chip > input {
      position: absolute;
      opacity: 0;
      width: 1px;
      height: 1px;
      pointer-events: none;
    }
    .chip:has(> input:checked) {
      background: var(--green-weak);
      border-color: var(--green);
      box-shadow: 0 0 0 3px #1e63d92e inset;
      color: var(--font-color-dark);
    }
    .chip:has(> input:focus-visible) {
      box-shadow: 0 0 0 3px var(--ring);
    }

    @supports not (selector(:has(*))) {
      .chip > input:checked + span {
        background: #e8f1ff;
        border-radius: 999px;
        padding: 6px 10px;
        margin: -6px -10px;
        box-shadow: 0 0 0 2px #bfd0ff inset;
      }
      .chip > input:focus-visible + span {
        box-shadow: 0 0 0 3px var(--ring);
      }
    }
  }
}
</style>
