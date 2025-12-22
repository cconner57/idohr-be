<script setup lang="ts">
import { useRouter } from 'vue-router'
import { onMounted } from 'vue'
import { usePets } from '../composables/usePets.ts'

import BannerButton from '../components/common/ui/BannerButton.vue'
import Footer from '../components/common/footer/Footer.vue'
import AdoptionSpotlight from '../components/home/spotlight/AdoptionSpotlight.vue'
import HeroSection from '../components/home/hero-section/HeroSection.vue'
import Mission from '../components/home/mission/Mission.vue'
import Impact from '../components/home/impact/Impact.vue'

const router = useRouter()

const { spotlightPets, loading, error, fetchSpotlight } = usePets()

onMounted(() => {
  fetchSpotlight()
})
</script>

<template>
  <div class="container">

    <HeroSection />
    <main class="section-1">
      <div class="content-wrapper">
        <AdoptionSpotlight :pets="spotlightPets" :loading="loading" :error="error" />
        <Mission />
      </div>
    </main>
    <main class="section-2">
      <div class="content-wrapper">
        <Impact />
        <section class="call-to-action">
          <BannerButton
            imgSrc="/images/paw.svg"
            title="Adopt a Pet"
            subtitle="Find your perfect companion"
            color="blue"
            @click="() => router.push('/adopt')"
          />
          <BannerButton
            imgSrc="/images/hand.svg"
            title="Get Involved"
            subtitle="Volunteer, foster, or support us"
            color="purple"
            @click="() => router.push('/volunteer')"
          />
          <BannerButton
            imgSrc="/images/heart.svg"
            title="Donate"
            subtitle="Help us save more lives"
            color="green"
            @click="() => router.push('/donate')"
          />
        </section>
      </div>
    </main>

    <Footer />
  </div>
</template>

<style scoped lang="css">
.container {
  width: 100%;
}

.content-wrapper {
  width: 100%;
  max-width: 1600px;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  gap: 64px;

  /* Media query gap adjustments... we can move them here or keep them on main */
  /* Actually main had the gap. Let's move gap logic to content-wrapper if we are wrapping EVERYTHING inside main */
}

main {
  display: flex;
  flex-direction: column;
  align-items: center; /* Center the wrapper */
  /* gap: 64px;  <-- This gap was between AdoptionSpotlight and Mission. moving to content-wrapper */
  width: 100%;
}

/* Move gap media queries to content-wrapper */
.content-wrapper {
  /* ... base styles ... */
  display: flex;
  flex-direction: column;
  gap: 64px;

  @media (min-width: 0px) and (max-width: 320px) {
    gap: 24px;
  }
  @media (min-width: 321px) and (max-width: 430px) {
    gap: 40px;
  }
  @media (min-width: 431px) and (max-width: 768px) {
    gap: 48px;
  }
  @media (min-width: 769px) and (max-width: 1024px) {
    gap: 56px;
  }
}

.section-1 {
  background-color: var(--background);
  padding: 0 var(--layout-padding-side) 220px;
  margin-top: -20px;
  min-height: auto;

  /* Media queries for padding/margin */
  /* 0px - 320px */
  @media (min-width: 0px) and (max-width: 320px) {
    margin-top: 0;
    padding-bottom: 40px;
  }
  @media (min-width: 321px) and (max-width: 430px) {
    margin-top: 0;
    padding-bottom: 60px;
  }
  @media (min-width: 431px) and (max-width: 768px) {
    margin-top: 0;
    padding-bottom: 80px;
  }
  @media (min-width: 769px) and (max-width: 1024px) {
    padding-bottom: 120px;
  }
}

.section-2 {
  background-color: var(--white);
  padding: 60px var(--layout-padding-side) 80px;
  height: auto;

   /* Media queries for padding/margin */
  @media (min-width: 0px) and (max-width: 320px) {
    padding-top: 40px;
    padding-bottom: 60px;
  }
  @media (min-width: 321px) and (max-width: 430px) {
    padding-top: 60px;
    padding-bottom: 80px;
  }
  @media (min-width: 431px) and (max-width: 768px) {
    padding-top: 80px;
    padding-bottom: 100px;
  }
  @media (min-width: 769px) and (max-width: 1024px) {
    padding-top: 100px;
    padding-bottom: 150px;
  }
}

.call-to-action {
  display: flex;
  margin: 20px 0 0;
  gap: 60px;
  width: 100%;
  justify-content: flex-start;

  & > * {
    flex: 1;
  }

  /* 0px - 320px */
  @media (min-width: 0px) and (max-width: 320px) {
    flex-direction: column;
    gap: 16px;
    margin: 20px 0 40px;
  }

  /* 321px - 430px */
  @media (min-width: 321px) and (max-width: 430px) {
    flex-direction: column;
    gap: 20px;
    margin: 20px 0 40px;
  }

  /* 431px - 768px */
  @media (min-width: 431px) and (max-width: 768px) {
    flex-direction: column;
    gap: 32px;
  }

  /* 769px - 1024px */
  @media (min-width: 769px) and (max-width: 1024px) {
    gap: 40px;
  }
}

.adopt-now-button {
  background-color: var(--green);
  height: 48px;
  min-width: 160px;
  padding: 0 24px;
  font-size: 1rem;
  font-weight: 600;
  border-radius: 6px;
}
</style>
