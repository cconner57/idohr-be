<script setup lang="ts">
defineProps<{ name: string }>()
</script>

<template>
  <div class="vaccination-item">
    <template v-if="vaccinationRecord">
      <li v-if="pet.medical?.vaccinations?.rabies">
            <p>Rabies</p>
            <p>
              {{
                pet.medical.vaccinations.rabies.dateAdministered
                  ? 'Administered on ' + formatDate(pet.medical.vaccinations.rabies.dateAdministered)
                  : 'Not administered'
              }}
            </p>
          </li>
    </template>
    <template v-if="vaccinationSeries">
      <li v-if="pet.medical?.vaccinations?.felineDistemper">
            <p>Feline Distemper ({{ pet.medical.vaccinations.felineDistemper.isComplete ? 'Complete' : 'Incomplete' }})</p>
            <p>
              <span v-if="pet.medical.vaccinations.felineDistemper.round1">
                Round 1: {{ formatDate(pet.medical.vaccinations.felineDistemper.round1.dateAdministered) }}
              </span>
              <span v-if="pet.medical.vaccinations.felineDistemper.round2">
                , Round 2: {{ formatDate(pet.medical.vaccinations.felineDistemper.round2.dateAdministered) }}
              </span>
              <span v-if="pet.medical.vaccinations.felineDistemper.round3">
                , Round 3: {{ formatDate(pet.medical.vaccinations.felineDistemper.round3.dateAdministered) }}
              </span>
            </p>
          </li>
        </template>
  </div>
</template>

<style scoped lang="css">
.vaccination-item {
}
</style>
