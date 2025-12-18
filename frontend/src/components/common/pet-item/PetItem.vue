<script setup lang="ts">
import { useRouter } from 'vue-router'
import { goToAdopt } from '../../../utils/navigate.ts'
import Button from '../ui/Button.vue'
import { ref, type PropType } from 'vue'
import Capsules from '../ui/Capsules.vue'

const props = defineProps({
  name: {
    type: String,
    required: true,
  },
  id: {
    type: String,
    required: true,
  },
  description: {
    type: String,
    required: false,
  },
  capsules: {
    type: Array as PropType<string[]>,
    required: false,
    default: () => [],
  },
  photo: {
    type: String as PropType<string | null>,
    required: false,
  },
  size: {
    type: String as PropType<'small' | 'medium' | 'large'>,
    required: false,
    default: 'medium',
  },
})
const router = useRouter()

const imgError = ref(false)

function onImgError() {
  imgError.value = true
}

function handleAdopt() {
  goToAdopt(router, props.id.toLowerCase())
}
</script>

<template>
  <div class="pet-item">
    <img
      v-if="!imgError"
      :src="`/images/${props.photo ?? ''}`"
      :alt="props.name"
      height="250"
      width="240"
      loading="lazy"
      @error="onImgError"
      @click="handleAdopt"
    />
    <div v-else class="img-fallback" aria-hidden="true" @click="handleAdopt"></div>
    <div class="info-section">
      <h5>{{ props.name }}</h5>
      <div v-if="props.capsules.length > 0" class="capsules">
        <template v-for="capText in props.capsules" :key="capText">
          <Capsules v-if="capText && capText !== 'Invalid Date'">{{ capText }}</Capsules>
        </template>
      </div>
      <p v-if="props.description">{{ props.description }}</p>
      <div class="adopt-button">
        <Button title="Adopt Me" color="blue" @click="handleAdopt" :fullWidth="true" />
      </div>
    </div>
  </div>
</template>

<style scoped lang="css">
.pet-item {
  display: flex;
  flex-direction: column;
  gap: 12px;
  width: 280px;
  border-radius: 8px;
  overflow: hidden;
  background-color: var(--white);
  color: var(--font-color-dark);
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.25);
  height: 400px;

  img {
    width: 100%;
    object-fit: cover;
    height: 180px;
    background: url('/images/paw.svg') 90px 60px/100px 100px no-repeat #add8e6;
    cursor: pointer;
  }

  .img-fallback {
    width: 100%;
    height: 180px;
    background: url('/images/paw.svg') 90px 60px/100px 100px no-repeat #add8e6;
  }

  .info-section {
    display: flex;
    flex-direction: column;
    padding: 0 20px 16px;
    height: 100%;
  }

  h5 {
    font-size: 1.5rem;
    margin-bottom: 8px;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  .capsules {
    display: flex;
    gap: 8px;
    flex-wrap: wrap;
    margin-bottom: 16px;
  }

  p {
    font-size: 1rem;
    flex-grow: 1;
    margin-bottom: 12px;
  }

  .adopt-button {
    margin-top: auto;
  }

  @media (min-width: 321px) and (max-width: 430px) {
    & .img-fallback {
      height: 400px;
      background: url('/images/paw.svg') 90px 50px/100px 100px no-repeat #add8e6;
    }
    & .info-section {
      & .capsules {
        margin-bottom: 8px;
        gap: 2px;
      }
    }
  }
  @media (min-width: 431px) and (max-width: 768px) {
  }
  @media (min-width: 769px) and (max-width: 1024px) {
  }
  @media (min-width: 1025px) and (max-width: 1440px) {
    width: 240px;
    height: 360px;
    & .img-fallback {
      height: 360px;
      background: url('/images/paw.svg') 70px 50px/100px 100px no-repeat #add8e6;
    }
    & .info-section {
      & .capsules {
        margin-bottom: 12px;
        gap: 2px;
      }
    }
  }
  @media (min-width: 1441px) {
    width: 260px;
    height: 380px;
    & .img-fallback {
      height: 380px;
      background: url('/images/paw.svg') 80px 55px/100px 100px no-repeat #add8e6;
    }
  }
}
</style>
