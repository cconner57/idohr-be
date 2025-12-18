<script setup lang="ts">
import SurrenderCat from '../components/about/surrender/SurrenderCat.vue'
import Candid from '../components/common/candid-award/Candid.vue'
</script>

<template>
  <main class="about">
    <section class="hero">
      <h1>About IDOHR</h1>
      <p class="lead">
        I Dream of Home Rescue (IDOHR) is a volunteer-powered nonprofit helping homeless cats and
        dogs heal, thrive, and find loving homes across Southern California.
      </p>
    </section>

    <section class="story">
      <content>
        <img src="/images/watercolor.jpeg" alt="IDOHR watercolor inspiration" />
      </content>
      <div class="story-text">
        <h2>Our Story</h2>
        <p>
          IDOHR began with a simple dream: to give every abandoned, neglected, and unwanted animal a
          safe place to call home. What started as a small circle of fosters has grown into a
          community of adopters, volunteers, and supporters who believe compassion changes
          lives—both animal and human.
        </p>
      </div>
    </section>

    <section class="mission">
      <content>
        <h2 class="text-on-teal">Our Mission</h2>
        <p class="text-on-teal">
          We rescue, rehabilitate, and rehome animals in need—one tail, one paw, one heart at a
          time. Every animal receives medical care, a safe foster home, love-based socialization,
          and a path to a forever family.
        </p>
        <ul class="ticks text-on-teal">
          <li>Vaccinations, microchipping, and spay/neuter</li>
          <li>Medical treatment and behavioral support</li>
          <li>Adopter education and lifetime support</li>
        </ul>
      </content>
      <img class="mission-image" src="/images/mission.png" alt="Happy adopted dog and cat" />
    </section>

    <section class="transparency">
      <div class="awards">
        <Candid type="Gold" year="2024" />
        <Candid type="Gold" year="2023" />
        <Candid type="Silver" year="2022" />
      </div>
      <div class="notice">
        <h2>Transparency</h2>
        <p>
          IDOHR is a 501(c)(3) nonprofit. Donations are tax-deductible as allowed by law. We publish
          annual updates and operate with financial transparency.
        </p>
        <p class="small">EIN: 81-0780050 • PO Box 7612, La Verne, CA 91750</p>
      </div>
    </section>
  </main>
  <SurrenderCat />
</template>

<style scoped lang="css">
.about {
  width: 100%;
  overflow: hidden;
  color: var(--font-color-dark);
  .hero {
    padding: 150px 180px 50px;
    color: var(--font-color-light);
    background-color: var(--green);
    h1 {
      font-size: 2.5rem;
    }
    p {
      font-size: 1.25rem;
      max-width: 600px;
      margin-top: 12px;
    }
  }
  .story,
  .mission {
    display: flex;
    justify-content: space-between;
    gap: 8rem;
    & h2 {
      font-size: 2.5rem;
    }
  }

  .story {
    display: grid;
    grid-template-columns: 50% 50%;
    background-color: var(--white);
    p {
      font-size: 1.15rem;
      max-width: 600px;
      margin-top: 12px;
    }
  }

  .mission {
    display: grid;
    grid-template-columns: 50% 50%;
    background-color: var(--green);
    color: var(--font-color-light);
    padding: 50px 180px 50px;
    p {
      font-size: 1.15rem;
      max-width: 600px;
      margin-top: 12px;
    }

    > div,
    > img {
      width: 80%;
      box-sizing: border-box;
    }

    img {
      align-self: center;
      justify-self: center;
      margin-right: 150px;
      width: 225px;
      max-width: 100%;
      height: auto;
      @media (max-width: 440px) {
        margin-right: 0px;
      }
    }

    & .ticks {
      list-style: none;
      padding: 0;
      margin: 12px 0 0;
    }
    & .ticks li {
      position: relative;
      padding-left: 24px;
      margin: 6px 0;
    }
    & .ticks li::before {
      content: '✓';
      position: absolute;
      left: 0;
      color: #bbf7d0;
      font-weight: 700;
    }
  }
  img {
    border-radius: 8px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.25);
  }
  .transparency {
    display: grid;
    grid-template-columns: 50% 50%;
    padding: 50px 180px 50px;
    color: var(--font-color-dark);
    background-color: var(--white);
    p {
      font-size: 1.15rem;
      max-width: 600px;
      margin-top: 12px;
    }
    .notice {
      margin-left: 8rem;
      & h2 {
        font-size: 2.5rem;
      }
      .small {
        font-size: 0.85rem;
        color: var(--font-color-medium);
        margin-top: 12px;
      }
    }
    .awards {
      display: flex;
      gap: 3rem;
      align-items: center;
      justify-content: center;
    }
  }
  @media (min-width: 0px) and (max-width: 320px) {
  }
  @media (min-width: 321px) and (max-width: 430px) {
    .hero {
      padding: 100px 30px 50px;
      h1 {
        font-size: 2rem;
      }
      p {
        font-size: 1.15rem;
      }
    }
    .story {
      padding: 50px 30px 50px;
      grid-template-columns: 100%;
      gap: 30px;
    }
    .mission {
      @media (max-width: 430px) {
        padding: 50px 30px 50px;
        grid-template-columns: 100%;
        gap: 30px;
      }
    }
    .transparency {
      padding: 50px 30px 50px;
      grid-template-columns: 100%;
      gap: 30px;
      .notice {
        margin-left: 0;
      }
      & .awards {
        flex-direction: column;
        justify-content: center;
        margin-top: 1.5rem;
      }
    }
  }
  @media (min-width: 431px) and (max-width: 768px) {
  }
  @media (min-width: 769px) and (max-width: 1024px) {
  }
  @media (min-width: 1025px) and (max-width: 1440px) {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    & .hero {
      padding: 120px 100px 50px;
      & p {
        font-weight: 400;
      }
    }
    & .story {
      width: 100%;
      display: flex;
      justify-content: center;
      padding: 50px 170px 50px;
      & content {
        display: grid;
        grid-template-columns: 50% 50%;
        max-width: 1120px;

        width: 100%vw;
        & img {
          max-width: 450px;
          height: auto;
        }
      }
      & .story-text {
        padding-left: 95px;
        & p {
          font-size: 1.15rem;
          margin-top: 12px;
        }
      }
    }
    & .mission {
      max-width: 1120px;
      padding: 50px 0px 50px;
      & p {
        font-size: 1.15rem;
        margin-top: 12px;
        font-weight: 400;
      }
    }
  }
  @media (min-width: 1441px) {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    & .hero {
      padding: 120px 100px 50px;
      & p {
        font-weight: 400;
      }
    }
    & .story {
      width: 100%;
      display: flex;
      justify-content: center;
      padding: 50px 170px 50px;
      & content {
        display: grid;
        grid-template-columns: 50% 50%;
        width: 100%vw;
        & img {
          max-width: 450px;
          height: auto;
        }
      }
      & .story-text {
        padding-left: 95px;
        & p {
          font-size: 1.15rem;
          margin-top: 12px;
        }
      }
    }
    & .mission {
      padding: 50px 0px 50px;
      & p {
        font-size: 1.15rem;
        margin-top: 12px;
        font-weight: 400;
      }
    }
    & .transparency {
      width: 100dvw;
    }
  }
}
</style>
