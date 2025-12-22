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
  width: 100%;
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
    justify-content: space-between;
    width: 100%;
    color: var(--font-color-dark);
    flex-wrap: wrap;
    gap: 40px;
  }
  & .awards {
    display: flex;
    gap: 50px;
    min-width: 300px;
    justify-content: center;
    flex-wrap: wrap;
  }
  .divider {
    border-left: 2px solid #e5e7eb;
    height: 120px;
    margin: 0 20px;
  }
  .stats {
    display: flex;
    gap: 60px;
    flex: 1;
    min-width: 300px;
    justify-content: center;

    & span {
      display: flex;
      flex-direction: column;
      gap: 8px;
      align-items: center;
      text-align: center;

      & h5 {
        font-size: 3rem;
        color: var(--font-color-dark);
      }
      & p {
        font-size: 1.1rem;
        color: var(--font-color-dark);
        max-width: 150px;
      }
    }
  }

  @media (min-width: 431px) and (max-width: 768px) {
    margin-top: -100px;
    .awards {
      margin-top: 20px;
      gap: 50px;
    }
    .divider {
      height: 1px;
      margin: 20px 0 0;
      border-left: none;
      border-top: 2px solid #e5e7eb;
      width: 80%;
      margin-left: auto;
      margin-right: auto;
    }
  }
  /* Mobile and Small Tablet */
  @media (min-width: 769px) and (max-width: 1024px) {
    margin-top: -80px;
    padding: 30px;
    height: auto;

    & h4 {
      margin-left: 0;
      text-align: center;
      width: 100%;
      margin-bottom: 20px;
    }

    & content {
      flex-direction: column;
      gap: 40px;
    }

    .divider {
      width: 80%;
      height: 2px;
      border-left: none;
      border-top: 2px solid #e5e7eb;
      margin: 0;
    }

    & .awards {
      width: 100%;
      gap: 50px;
    }

    .stats {
      width: 100%;
      gap: 40px;
    }
  }

  @media (max-width: 480px) {
    padding: 24px;

    .stats {
      flex-direction: column;
      gap: 30px;
    }
  }
}
</style>
