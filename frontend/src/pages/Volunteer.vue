<script setup lang="ts">
import {
  Agreement,
  Allergies,
  ApplicationHeader,
  Availability,
  PositionPreferences,
} from '../components/volunteer/index.ts'
import Button from '../components/common/ui/Button.vue'
import InputField from '../components/common/ui/InputField.vue'
import InputTextArea from '../components/common/ui/InputTextArea.vue'
import type { IVolunteerFormState } from '../models/volunteer-form.ts'
import { computed, reactive } from 'vue'

const formatPhoneNumber = (value: string | number | null): string => {
  if (!value) return ''
  const digits = String(value).replace(/\D/g, '').substring(0, 10)
  if (digits.length === 0) return ''
  if (digits.length <= 3) return `(${digits}`
  if (digits.length <= 6) return `(${digits.slice(0, 3)})${digits.slice(3)}`
  return `(${digits.slice(0, 3)})${digits.slice(3, 6)}-${digits.slice(6)}`
}

const sanitizeName = (value: string | number | null): string => {
  if (!value) return ''
  return String(value).replace(/[^a-zA-Z0-9 ]/g, '')
}

const sanitizeCity = (value: string | number | null): string => {
  if (!value) return ''
  return String(value).replace(/[^a-zA-Z0-9 -]/g, '')
}

const sanitizeZip = (value: string | number | null): string => {
  if (!value) return ''
  return String(value).replace(/[^a-zA-Z0-9 -]/g, '')
}

const sanitizeAddress = (value: string | number | null): string => {
  if (!value) return ''
  // Strict: Alphanumeric + Space + Dash only.
  return String(value).replace(/[^a-zA-Z0-9 -]/g, '')
}

const formState = reactive<IVolunteerFormState>({
  firstName: '',
  lastName: '',
  address: '',
  city: '',
  zip: '',
  phoneNumber: '',
  birthday: '',
  age: null as number | null,
  allergies: '',
  emergencyContactName: '',
  emergencyContactPhone: '',
  volunteerExperience: '',
  interestReason: '',
  positionPreferences: [] as string[],
  availability: [] as string[],
  nameFull: '',
  signatureData: null as string | null,
  signatureDate: '',
  parentName: '',
  parentSignatureData: null as string | null,
  parentSignatureDate: '',
})

const isFormValid = computed(() => {
  // Required Personal Info
  if (!formState.firstName || !formState.lastName || !formState.address || !formState.city || !formState.zip || !formState.phoneNumber || !formState.birthday || formState.age === null || !formState.allergies || !formState.emergencyContactName || !formState.emergencyContactPhone) {
    return false
  }

  // Experience & Interests
  if (!formState.interestReason || formState.positionPreferences.length === 0 || formState.availability.length === 0) {
    return false
  }

  // Agreement
  if (!formState.nameFull || !formState.signatureDate || !formState.signatureData) {
    return false
  }

  // Parental Consent (Under 21)
  if (formState.age !== null && formState.age < 21) {
    if (!formState.parentName || !formState.parentSignatureDate || !formState.parentSignatureData) {
      return false
    }
  }

  return true
})

const handleSubmit = () => {
  if (!isFormValid.value) return

  console.log('Form submitted successfully:', formState)
  alert('Application submitted successfully!')
  // TODO: Add API call here
}
</script>

<template>
  <section class="page-shell">
    <section class="form-card" aria-labelledby="form-title">
      <ApplicationHeader
        header-title="Volunteer"
        header-text="I Dream of Home Rescue (IDOHR) is an all-volunteer, nonprofit dedicated to helping homeless cats
    and kittens find loving, permanent homes. We’re looking for responsible volunteers to help with
    feeding and cleaning, socializing cats and kittens, supporting adoptions, and light
    administrative tasks. Volunteers must be 21 or older. If under 21, a parent or guardian must
    sign the waiver below. Join us and help change a life. You’ll connect with amazing animals, work
    alongside caring volunteers, and make a meaningful impact."
      />
      <fieldset class="grid" aria-labelledby="pi">
        <legend id="pi" class="section-title">Personal Information</legend>

        <InputField
          :modelValue="formState.firstName"
          @update:modelValue="(val) => (formState.firstName = sanitizeName(val))"
          label="First Name"
          placeholder="First name"
          autocomplete="given-name"
          maxlength="50"
        />
        <InputField
          :modelValue="formState.lastName"
          @update:modelValue="(val) => (formState.lastName = sanitizeName(val))"
          label="Last Name"
          placeholder="Last name"
          autocomplete="family-name"
          maxlength="50"
        />
        <InputField
          :modelValue="formState.address"
          @update:modelValue="(val) => (formState.address = sanitizeAddress(val))"
          label="Address"
          placeholder="Address"
          autocomplete="street-address"
          maxlength="100"
        />
        <InputField
          :modelValue="formState.city"
          @update:modelValue="(val) => (formState.city = sanitizeCity(val))"
          label="City"
          placeholder="City"
          autocomplete="address-level2"
          maxlength="50"
        />
        <InputField
          :modelValue="formState.zip"
          @update:modelValue="(val) => (formState.zip = sanitizeZip(val))"
          label="Zip"
          placeholder="Zip"
          type="text"
          autocomplete="postal-code"
          maxlength="10"
        />
        <InputField
          :modelValue="formState.phoneNumber"
          @update:modelValue="(val) => (formState.phoneNumber = formatPhoneNumber(val))"
          label="Phone Number"
          placeholder="Phone"
          type="tel"
          autocomplete="tel"
          maxlength="15"
        />
        <InputField
          v-model="formState.birthday"
          label="Birthday"
          placeholder="mm/dd/yyyy"
          type="date"
          autocomplete="bday"
        />
        <InputField
          v-model="formState.age"
          label="If under 21, Age"
          placeholder="Age"
          type="number"
          min="0"
          max="100"
        />

        <Allergies v-model="formState.allergies" />

        <InputField
          :modelValue="formState.emergencyContactName"
          @update:modelValue="(val) => (formState.emergencyContactName = sanitizeName(val))"
          label="Emergency Contact Person"
          placeholder="Name"
          maxlength="100"
        />
        <InputField
          :modelValue="formState.emergencyContactPhone"
          @update:modelValue="(val) => (formState.emergencyContactPhone = formatPhoneNumber(val))"
          label="Phone Number"
          placeholder="Phone Number"
          type="tel"
          maxlength="15"
        />
      </fieldset>

      <fieldset aria-labelledby="exp">
        <legend id="exp" class="section-title">Experience & Interests</legend>

        <InputTextArea
          v-model="formState.volunteerExperience"
          label="Volunteer Experience (if any):"
          placeholder="Briefly describe"
          maxlength="500"
        />

        <InputTextArea
          v-model="formState.interestReason"
          label="Why are you interested in being a volunteer:"
          placeholder="Your reason"
          maxlength="500"
        />

        <PositionPreferences v-model="formState.positionPreferences" />
      </fieldset>

      <Availability v-model="formState.availability" />

      <Agreement
        :formState="formState"
        :name="formState.firstName + ' ' + formState.lastName"
        :fullName="formState.nameFull"
        @update:fullName="(val: string) => (formState.nameFull = sanitizeName(val))"
        :age="formState.age!"
        :signature="formState.signatureData"
        :signature-date="formState.signatureDate"
        :parent-name="formState.parentName"
        @update:parentName="(val: string) => (formState.parentName = sanitizeName(val))"
        :parent-signature="formState.parentSignatureData"
        :parent-date="formState.parentSignatureDate"
      />

      <div class="actions">
        <Button
          @click="handleSubmit"
          type="submit"
          title="Submit Application"
          color="green"
          size="large"
          :disabled="!isFormValid"
        />
      </div>
    </section>
  </section>
</template>

<style scoped>
.page-shell {
  min-height: 100vh;
  background-color: var(--green);
  padding: 9rem var(--layout-padding-side) 64px;
  .form-card {
    max-width: 1600px;
    margin: 0 auto;
    background: var(--white);
    color: var(--font-color-dark);
    border-radius: 24px;
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
    padding: 48px 48px 32px;
    fieldset {
      border: 0;
      margin: 24px 0;
      padding: 0;
      .section-title {
        font-weight: 700;
        font-size: 18px;
        margin: 18px 0 12px;
      }
    }
    .actions {
      display: flex;
      justify-content: center;
      gap: 16px;
      margin-top: 20px;
    }
    .grid {
      display: grid;
      grid-template-columns: repeat(2, minmax(0, 1fr));
      gap: 16px;
    }
  }
  @media (min-width: 321px) and (max-width: 430px) {
    padding: 6rem 16px 32px;
  }
  @media (min-width: 431px) and (max-width: 768px) {
  }
}
</style>
