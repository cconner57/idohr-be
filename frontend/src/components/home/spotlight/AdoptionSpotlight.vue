<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import PetItem from '../../common/pet-item/PetItem.vue'
import { useIsMobile } from '../../../utils/useIsMobile.ts'
import type { IPet } from '../../../models/common.ts'

const props = defineProps<{
  pets: IPet[]
  loading: boolean
  error: string | null
}>()

const isMobile = useIsMobile()

const randomPet = ref<IPet | null>(null)
console.log('AdoptionSpotlight props.pets:', props.pets)
watch(
  [() => props.pets, isMobile],
  ([newPets, newIsMobile], [oldPets, oldIsMobile]) => {
    if (!newIsMobile) return

    const hasPets = newPets && newPets.length > 0
    if (!hasPets) return

    const justSwitchedToMobile = !oldIsMobile
    const justLoadedPets = !oldPets || oldPets.length === 0

    if (justSwitchedToMobile || justLoadedPets) {
      randomPet.value = newPets[Math.floor(Math.random() * newPets.length)]
    }
  },
  { immediate: true },
)

const displayedPets = computed((): IPet[] => {
  if (isMobile.value) {
    return randomPet.value ? [randomPet.value] : []
  }
  return props.pets
})
</script>

<template>
  <section class="adoption-spotlight">
    <h4>Adoption Spotlight</h4>
    <div class="pet-list">
      <PetItem
        v-for="pet in displayedPets"
        :key="pet.id"
        :name="pet.name"
        :id="pet.id.toLowerCase()"
        :photo="pet.photos?.find((p) => p.isPrimary)?.url || null"
        :description="pet.descriptions?.spotlight || ''"
        :size="isMobile ? 'large' : 'medium'"
      />
    </div>
  </section>
</template>

<style scoped lang="css">
.adoption-spotlight {
  background-color: var(--white);
  border-radius: 12px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.25);
  display: flex;
  flex-direction: column;
  gap: 16px;
  margin: -120px 0 0;
  width: 100%;
  padding: 24px 50px 40px 50px;

  & h4 {
    font-size: 2rem;
    color: var(--font-color-dark);
  }

  .pet-list {
    display: flex;
    gap: 3rem;
    flex-wrap: nowrap;
    overflow-x: auto;
    padding-bottom: 8px;
    -webkit-overflow-scrolling: touch;
    justify-content: flex-start;
  }

  @media (max-width: 430px) {
    margin: 2rem 0 0;
    padding: 1rem 2rem;
    gap: 0.5rem;
    & h4 {
      font-size: 1.5rem;
    }
    .pet-list {
      gap: 1rem;
    }
  }
  @media (min-width: 431px) and (max-width: 768px) {
    margin: -20px 0 0;
    & h4 {
      font-size: 1.75rem;
    }
  }
  @media (min-width: 769px) and (max-width: 1024px) {
    & h4 {
      font-size: 1.75rem;
    }
  }
  @media (min-width: 1025px) and (max-width: 1440px) {
    width: 100%;
    padding: 24px 30px 30px;
    & h4 {
      font-size: 1.75rem;
    }
    .pet-list {
      gap: 2rem;
    }
  }
}
</style>
