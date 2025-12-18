<script setup lang="ts">
import type { IPet } from '../../../models/common'
import PetItem from '../../common/pet-item/PetItem.vue'
import { formatDate } from '../../../utils/common'

defineProps<{
  pets: IPet[]
}>()
</script>

<template>
  <div class="adopt-summary">
    <PetItem
      v-for="pet in pets"
      :capsules="[
        pet?.physicalTraits?.species ?? '',
        pet?.physicalTraits?.sex ?? '',
        pet?.physicalTraits?.age ? formatDate(pet?.physicalTraits?.age ?? '', true) : '',
      ]"
      :description="pet.descriptions?.physicalDescription ?? ''"
      :id="pet.id"
      :key="pet.id"
      :name="pet.name"
      :photo="pet.photos?.primaryPhoto"
    />
  </div>
</template>

<style scoped lang="css">
.adopt-summary {
  display: flex;
  flex-wrap: wrap;
  row-gap: 30px;
  column-gap: 20px;
  @media (min-width: 321px) and (max-width: 430px) {
    justify-content: center;
  }
  @media (min-width: 431px) and (max-width: 768px) {
    justify-content: center;
  }
  @media (min-width: 769px) and (max-width: 1024px) {
  }
  @media (min-width: 1025px) and (max-width: 1440px) {
    justify-content: center;
    align-items: center;
    max-width: 1275px;
  }
}
</style>
