<script setup lang="ts">
import { computed } from 'vue'
import type { IPet } from '../../../models/common'
import { formatDate } from '../../../utils/common'

const props = defineProps<{
  pet: IPet
}>()

const isSpayedOrNeutered = (pet: IPet) => {
  return pet.physicalTraits?.sex === 'Male' ? 'Neutered' : 'Spayed'
}

const goodWithText = computed(() => {
  const traits = props.pet.behavioralTraits
  if (!traits || (!traits.goodWithCats && !traits.goodWithDogs && !traits.goodWithKids)) {
    return 'N/A'
  }

  const goodWith: string[] = []
  if (traits.goodWithCats) goodWith.push('Other Cats')
  if (traits.goodWithDogs) goodWith.push('Other Dogs')
  if (traits.goodWithKids) goodWith.push('Kids')

  return goodWith.join(', ')
})
const houseTrainedText = () => {
  if (props.pet.behavioralTraits?.houseTrained === undefined) {
    return 'N/A'
  }
  if (props.pet.behavioralTraits?.houseTrained) {
    return 'Yes'
  } else {
    return 'No'
  }
}
</script>

<template>
  <div class="adopt-detail__additional-info">
    <div class="adopt-detail__additional-info__item">
      <p>Breed</p>
      <p>{{ pet.physicalTraits?.breed ?? 'N/A' }}</p>
    </div>
    <div class="adopt-detail__additional-info__item">
      <p>Color</p>
      <p>{{ pet.physicalTraits?.color ?? 'N/A' }}</p>
    </div>
    <div class="adopt-detail__additional-info__item">
      <p>Size</p>
      <p>{{ pet.physicalTraits?.size ?? 'N/A' }}</p>
    </div>
    <div class="adopt-detail__additional-info__item">
      <p>Birthday</p>
      <p>{{ formatDate(pet?.physicalTraits?.age ?? '') }}</p>
    </div>
    <div class="adopt-detail__additional-info__item">
      <p>House-trained</p>
      <p>{{ houseTrainedText() }}</p>
    </div>
    <div class="adopt-detail__additional-info__item">
      <p>Health</p>
      <p>
        {{ pet.medicalHistory?.vaccinationsUpToDate ? 'Vaccinated' : 'Not Vaccinated' }},
        {{
          pet.medicalHistory?.spayedOrNeutered
            ? isSpayedOrNeutered(pet)
            : `Not ${isSpayedOrNeutered(pet)}`
        }},
        {{ pet.medicalHistory?.microchipped ? 'Microchipped' : 'Not Microchipped' }}
      </p>
    </div>
    <div class="adopt-detail__additional-info__item">
      <p>Good in a home with</p>
      <p>{{ goodWithText }}</p>
    </div>
    <div class="adopt-detail__additional-info__item">
      <p>Adoption Fee</p>
      <p>{{ pet.adoption.fee !== null ? '$' + pet.adoption.fee : 'N/A' }}</p>
    </div>
  </div>
</template>

<style scoped lang="css">
.adopt-detail__additional-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 8px;

  @media (max-width: 440px) {
    gap: 5px;
    flex-direction: column;
    align-items: flex-start;
    p {
      font-size: 0.9rem;
      line-height: 1.5;
      text-wrap: wrap;
    }
    p:last-child {
      text-wrap: wrap;
      width: 150px;
      @media (max-width: 404px) {
        width: 130px;
      }
    }
  }
}

.adopt-detail__additional-info__item {
  display: flex;
  flex-direction: row;
  p {
    margin-right: 8px;
    text-transform: capitalize;
  }
  & p:first-child {
    width: 225px;
    @media (max-width: 404px) {
      width: 170px;
    }
  }
  & p:last-child {
    font-weight: bold;
    width: 300px;
  }
}
</style>
