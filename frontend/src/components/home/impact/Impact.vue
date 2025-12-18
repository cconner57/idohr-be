<script setup lang="ts">
import Candid from '../../common/candid-award/Candid.vue'
</script>

<template>
  <section class="impact">
    <h4>Our Impact</h4>
    <content>
      <div class="awards">
        <Candid type="Gold" year="2024" />
        <Candid type="Gold" year="2023" />
        <Candid type="Silver" year="2022" />
      </div>
      <div class="divider"></div>
      <div class="stats">
        <span>
          <h5>500+</h5>
          <p>pets rescued since 2020</p>
        </span>
        <span>
          <h5>500+</h5>
          <p>pets rescued since 2020</p>
        </span>
      </div>
    </content>
  </section>
</template>

<style scoped lang="css">
.impact {
  max-width: var(--desktop-breakpoint);
  background-color: var(--white);
  display: flex;
  flex-direction: column;
  gap: 16px;
  padding: 24px 50px 50px 50px;
  border-radius: 12px;
  margin-top: -200px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.25);

  & h4 {
    font-size: 2rem;
    color: var(--font-color-dark);
    margin-left: -15px;
  }
  & content {
    display: flex;
    align-items: center;
    width: 100%;
    color: var(--font-color-dark);
  }
  & .awards {
    display: flex;
    gap: 50px;
    width: 50%;
  }
  .divider {
    border-left: 5px solid gray;
    height: 150px;
    border-radius: 12px;
    margin: 0 90px 0 80px;
  }
  .stats {
    display: flex;
    gap: 100px;
    width: 45%;
    & span {
      display: flex;
      flex-direction: column;
      gap: 8px;
      & h5 {
        font-size: 3rem;
        color: var(--font-color-dark);
      }
      & p {
        font-size: 1.25rem;
        color: var(--font-color-dark);
      }
    }
  }

  @media (min-width: 321px) and (max-width: 430px) {
    margin-top: -80px;
    padding: 16px 20px 30px 20px;
    gap: 12px;
    overflow: hidden;
    & h4 {
      font-size: 1.5rem;
      margin-left: 0;
    }
    & content {
      flex-direction: column;
      align-items: center;
      gap: 20px;
    }
    & .awards {
      flex-direction: column;
      justify-content: center;
      margin-top: 1.5rem;
    }
    .divider {
      height: 5px;
      width: 250px;
      border-left: none;
      border-top: 5px solid gray;
      border-radius: 12px;
      margin: 2rem 0 0.5rem;
    }
    .stats {
      gap: 50px;
      justify-content: center;
      width: 100%;
      margin-left: 3rem;
      & span {
        display: flex;
        flex-direction: column;
        gap: 4px;
        & h5 {
          font-size: 2rem;
          margin-left: 0;
        }
        & p {
          font-size: 1rem;
          margin-left: 0;
        }
      }
    }
  }

  @media (min-width: 431px) and (max-width: 768px) {
  }
  @media (min-width: 769px) and (max-width: 1024px) {
  }
  @media (min-width: 1025px) and (max-width: 1440px) {
    max-width: 1120px;
    & h4 {
      font-size: 1.75rem;
    }
  }
}
</style>
