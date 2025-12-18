<script setup lang="ts">
import type { SurrenderFormState } from '../models/common'
import { reactive, ref, computed } from 'vue'
import {
  AggressiveSection,
  FeedingSection,
  HouseholdSection,
  MedicalSection,
  BehaviorSection,
  OtherSection,
} from '../components/about/surrender'
import Button from '../components/common/ui/Button.vue'
import SurrenderSteps from '../components/about/surrender/SurrenderSteps.vue'
import PetSelectSection from '../components/about/surrender/PetSelectSection.vue'

const formState = reactive<SurrenderFormState>({
  // Cat & Household Information
  firstName: '',
  lastName: '',
  phoneNumber: '',
  email: '',
  streetAddress: '',
  city: '',
  state: '',
  zipCode: '',
  whenToSurrenderCat: '',
  catName: '',
  catSex: '',
  catAge: '',
  catOwnershipDuration: '',
  catLocationFound: '',
  catWhySurrendered: '',
  agesOfHouseholdMembers: '',
  otherPetsInHousehold: '',

  // Behavior
  catsBehaviorTowardsKnownPeople: '',
  catsBehaviorTowardsStrangers: '',
  catsBehaviorTowardsKnownAnimals: '',
  commentsOnBehavior: '',
  catsReactionToNewPeople: '',
  catHouseTrained: '',
  catSpendMajorityOfTime: '',
  catLeftAloneDuration: '',
  catWhenLeftAlone: '',
  catLeftAloneBehaviors: '',
  catHowItPlays: '',
  catToysItLikes: '',
  catGamesItLikes: '',
  catScaredOfAnything: '',
  catScaredOfAnythingExplanation: '',
  catBadHabits: '',
  catAllowedOnFurniture: '',
  catSleepAtNight: '',
  catBehaviorFoodOthers: '',
  catBehaviorToysOthers: '',
  catProblemsRidingInCar: '',
  catProblemsRidingInCarExplanation: '',
  catEscapedBefore: '',
  catEscapedBeforeExplanation: '',

  // Aggressive Behavior
  catEverAttackedPeople: '',
  catEverAttackedPeopleExplanation: '',
  catEverAttackedOtherCats: '',
  catEverAttackedOtherCatsExplanation: '',

  // Medical History
  catVeterinarianList: '',
  catVeterinarianYearlyVisits: '',
  catSpayedNeutered: '',
  catVaccinationHistory: '',
  catVaccinationsCurrent: '',
  catTestedHeartworm: '',
  catTestedHeartwormExplanation: '',
  catHeartwormPrevention: '',
  catHeartwormPreventionExplanation: '',
  catMicrochipped: '',
  catMicrochippedExplanation: '',
  catVetOrGroomerBehavior: '',
  catVetMuzzled: '',
  catPastOrPresentHealthProblems: '',
  catPastOrPresentHealthProblemsExplanation: '',
  catCurrentMedications: '',
  catCurrentMedicationsExplanation: '',

  // Feeding
  catTypeOfFood: '',
  catEatingFrequency: '',
  catAmountOfFood: '',
  catFoodTreats: '',
  catFoodTreatsExplanation: '',

  // Other
  additionalInformation: '',
  fullBodyPhotoOfCat: '',
  closeUpPhotoOfCatsFace: '',
  copiesOfRecords: '',
})

const formStep = ref(0)
const selectedAnimal = ref<'dog' | 'cat' | null>(null)
const formError = ref<boolean>(false)

const handleSubmit = () => {
  if (formStep.value === 0 && !selectedAnimal.value) {
    formError.value = true
    console.log('Form submitted with state:', selectedAnimal.value)
    return
  }
  console.log('Form submitted with state:', formState)
  // Add form validation and submission logic here
  if (formStep.value < 6) {
    formStep.value += 1
  } else {
    // Final submission logic
    alert('Form submitted successfully!')
  }
}

const headerText = computed(() => {
  if (!selectedAnimal.value || formStep.value === 0) {
    return 'Surrender Pet'
  }
  return selectedAnimal.value === 'cat' ? 'Feline Surrender' : 'Canine Surrender'
})
</script>

<template>
  <section class="page-shell">
    <section class="form-card" aria-labelledby="form-title">
      <div class="form-header">
        <img
          v-if="selectedAnimal === 'cat' && formStep > 0"
          src="../../public/images/cat.png"
          alt="cat"
          height="50"
          width="100"
        />
        <img
          v-if="selectedAnimal === 'dog' && formStep > 0"
          src="../../public/images/dog.png"
          alt="cat"
          height="50"
          width="100"
        />
        <h1>{{ headerText }}</h1>
      </div>
      <SurrenderSteps
        v-if="selectedAnimal && formStep > 0"
        :formStep="formStep"
        :selectedAnimal="selectedAnimal"
      />
      <content>
        <PetSelectSection
          v-if="formStep === 0"
          :formError="formError"
          :selectedAnimal="selectedAnimal"
          @update:selectedAnimal="
            (value) => {
              selectedAnimal = value
            }
          "
        />
        <HouseholdSection v-if="formStep === 1" :formState="formState" />
        <BehaviorSection v-if="formStep === 2" :formState="formState" />
        <AggressiveSection v-if="formStep === 3" :formState="formState" />
        <MedicalSection v-if="formStep === 4" :formState="formState" />
        <FeedingSection v-if="formStep === 5" :formState="formState" />
        <OtherSection v-if="formStep === 6" :formState="formState" />
      </content>
      <div class="actions">
        <Button
          @click="handleSubmit"
          type="submit"
          :title="formStep === (selectedAnimal === 'dog' ? 6 : 5) ? 'Submit' : 'Next'"
          color="green"
          size="large"
          :disabled="formStep === 0 && !selectedAnimal"
        />
      </div>
    </section>
  </section>
</template>

<style scoped lang="css">
.page-shell {
  min-height: 100vh;
  background-color: var(--green);
  padding: 9rem 16px 64px;
  .form-card {
    max-width: 1100px;
    margin: 0 auto;
    background: var(--white);
    color: var(--font-color-dark);
    border-radius: 24px;
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
    padding: 48px 48px 32px;
    .form-header {
      display: flex;
      justify-content: center;
      align-items: center;
      gap: 16px;
      margin-bottom: 4px;
      color: var(--green);
      & h1 {
        font-size: 4.25rem;
        line-height: 1.2;
        letter-spacing: 0.2px;
      }
      img {
        width: 100px;
      }
    }
    .actions {
      display: flex;
      justify-content: center;
      gap: 16px;
      margin-top: 20px;
    }
  }
  @media (max-width: 440px) {
    .form-header {
      flex-direction: column;
      align-items: center;
      gap: 0px;
      margin-bottom: 1rem;
      & h1 {
        font-size: 1.75rem;
        text-align: center;
      }
    }
  }
}
</style>
