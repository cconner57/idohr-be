<script setup lang="ts">
import { useRoute } from 'vue-router'
import { ApplicationHeader } from '../components/volunteer'
import Button from '../components/common/ui/Button.vue'
import { reactive, ref } from 'vue'
import AdoptionSteps from '../components/pet-adoption/adoption-steps/AdoptionSteps.vue'
import CatAdoptionForm from '../components/pet-adoption/cat-adoption/CatAdoptionForm.vue'
import { mockPetsData } from '../stores/mockPetData'

const formState = reactive({
  firstName: null as string | null,
  lastName: null as string | null,
  age: null as number | null,
  spouseFirstName: null as string | null,
  spouseLastName: null as string | null,
  roommatesNames: null as string[] | null,
  childrenNamesAges: null as string[] | null,
  email: null as string | null,
  address: null as string | null,
  streetAddressLine2: null as string | null,
  city: null as string | null,
  state: null as string | null,
  zip: null as string | null,
  phoneNumber: null as string | null,
  cellPhoneNumber: null as string | null,
  adultMembersAgreed: null as string | null,
})

const formStep = ref(0)

const handleSubmit = () => {
  console.log('Form submitted with state:', formState)
}

const getPetbyId = (id: string) => {
  return mockPetsData.find((pet) => pet.id === id)
}

const route = useRoute()
const petId = route.params.id as string

const selectedCat = getPetbyId(petId)
</script>

<template>
  <section class="page-shell">
    <section class="form-card" aria-labelledby="form-title">
      <ApplicationHeader
        header-title="Pet"
        header-text="This application is intended as a means to match the right cat with the right home. The more detail you provide, the better.  All of our adoptable pets are spayed/neutered, vaccinated, and microchipped. Typical adoption fees are $300 for kittens and $250 for adults. Adoption fees are tax-deductible donations, not purchase prices. Thank you for considering adoption!"
      />
      <AdoptionSteps :formStep="formStep" selectedAnimal="cat" />
      <CatAdoptionForm v-model="formState" :catName="selectedCat?.name" />
      <div class="actions">
        <Button
          @click="handleSubmit"
          type="submit"
          :title="formStep < 6 ? 'Next' : 'Submit Application'"
          color="green"
          size="large"
          :disabled="
            Object.values(formState).every(
              (value: string | number | null | string[]) =>
                value === '' || value === null || (Array.isArray(value) && value.length === 0),
            )
          "
        />
      </div>
    </section>
  </section>
</template>

<style scoped lang="css">
.page-shell {
  min-height: 100vh;
  background-color: var(--green);
  padding: 7rem 16px 64px;
  .form-card {
    max-width: 1100px;
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
