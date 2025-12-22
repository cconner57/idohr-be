<script setup lang="ts">
import Footer from '../components/common/footer/Footer.vue'
</script>

<template>
  <div class="donate-page">
    <div class="header">
      <h1>Help Us Save Lives</h1>
      <p class="mission-text">
        You can save a life today! With your donation, IDOHR can only continue its lifesaving efforts … finding loving homes for wonderful rescued shelter dogs. We are a small operation but have many expenses, veterinary bills being first and foremost along with food and care for all dogs in our rescue. Any amount helps. We are a 501(c)(3) non-profit organization, so your contribution is 100% tax-deductible.
      </p>
    </div>

    <main class="donate-container">
      <section class="donate-card cat-card">
        <h2>Support Our Cats</h2>
        <div class="illustration">
             <img src="/images/shadow.jpeg" alt="Cute Cat" />
        </div>

        <div class="payment-methods">
            <div class="qr-section">
                <div class="qr-placeholder">
                    <span class="username">lorraine marrero</span>
                    <img src="/images/cat-paypal.jpg" class="real-qr" alt="PayPal QR Code" />
                    <span>Scan for PayPal</span>
                </div>
            </div>

            <a href="https://www.paypal.com/qrcodes/p2pqrc/FZYN4ZRPJWGXQ" target="_blank" rel="noopener noreferrer" class="donate-btn btn-paypal">
               <i class="icon-paypal"></i> Donate via PayPal
            </a>

            <div class="qr-section">
                <div class="qr-placeholder">
                    <span class="username">Lorraine Marrero-Baeza</span>
                    <span class="handle">@Lorraine-Marrero-Baeza</span>
                    <img src="/images/cat-venmo.jpg" class="real-qr" alt="Venmo QR Code" />
                    <span>Scan for Venmo</span>
                </div>
            </div>

            <a href="https://venmo.com/code?user_id=3025264426614784345&created=1762028642" target="_blank" rel="noopener noreferrer" class="donate-btn btn-venmo">
               <i class="icon-venmo"></i> Donate via Venmo
            </a>
        </div>
      </section>

      <section class="donate-card dog-card">
        <h2>Support Our Dogs</h2>
        <div class="illustration">
             <img src="/images/paw.svg" alt="Dog Paw" />
        </div>

        <div class="payment-methods">
            <div class="qr-section">
                <div class="qr-placeholder">
                    <img src="/images/paw.svg" class="qr-icon" alt="QR Code" />
                    <span>Scan for PayPal</span>
                </div>
            </div>

            <a href="https://www.paypal.com" target="_blank" rel="noopener noreferrer" class="donate-btn btn-paypal">
               <i class="icon-paypal"></i> Donate via PayPal
            </a>

            <div class="qr-section">
                <div class="qr-placeholder">
                    <img src="/images/paw.svg" class="qr-icon" alt="QR Code" />
                    <span>Scan for Venmo</span>
                </div>
            </div>

            <a href="#" class="donate-btn btn-venmo">
               <i class="icon-venmo"></i> Donate via Venmo
            </a>
        </div>
      </section>
    </main>

    <Footer />
  </div>
</template>

<style scoped>
.donate-page {
  background-color: var(--green);
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.header {
  text-align: center;
  padding: 100px 20px 40px;
}

.header h1 {
  color: var(--font-color-light);
  font-size: 3rem;
  margin: 0;
}

.mission-text {
  color: var(--font-color-light);
  font-size: 1.15rem;
  line-height: 1.6;
  max-width: 800px;
  margin: 24px auto 0;
  opacity: 0.95;
  font-weight: 500;
}

.donate-container {
  display: flex;
  justify-content: center;
  gap: 40px;
  padding: 0 var(--layout-padding-side) 100px;
  flex-wrap: wrap;
  flex: 1;
}

/* Common Card Styles */
.donate-card {
  background: var(--white);
  border-radius: 24px;
  padding: 40px;
  width: 100%;
  max-width: 500px;
  display: flex;
  flex-direction: column;
  align-items: center;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
  transition: transform 0.3s ease;
}

.donate-card:hover {
  transform: translateY(-5px);
}

.donate-card h2 {
    font-size: 2rem;
    margin-bottom: 24px;
    font-weight: 700;
}

.illustration img {
    height: 200px;
    width: 200px;
    object-fit: cover;
    border-radius: 50%;
    margin-bottom: 32px;
    border: 6px solid;
}

.payment-methods {
    display: flex;
    flex-direction: column;
    gap: 16px;
    width: 100%;
}

.donate-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 16px 24px;
    border-radius: 50px;
    text-decoration: none;
    font-weight: 600;
    font-size: 1.1rem;
    transition: all 0.2s ease;
    color: white;
    box-shadow: 0 4px 6px rgba(50, 50, 93, 0.11), 0 1px 3px rgba(0, 0, 0, 0.08);
}

.donate-btn:hover {
    transform: translateY(-2px);
    box-shadow: 0 7px 14px rgba(50, 50, 93, 0.1), 0 3px 6px rgba(0, 0, 0, 0.08);
    filter: brightness(1.1);
}

.qr-section {
    display: flex;
    justify-content: center;
    margin: 8px 0;
}


.qr-placeholder {
    width: 220px;
    height: auto;
    padding: 24px;
    background: transparent;
    border-radius: 20px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    min-height: 280px;
    color: #666;
    font-size: 1rem;
    text-align: center;
    gap: 8px;
    border: none;
    box-shadow: inset 0 2px 6px rgba(0,0,0,0.03);
}

.qr-placeholder > span:not(.username):not(.handle) {
    font-weight: 600;
    color: #222;
    letter-spacing: 0.5px;
}

.username {
    font-weight: 700;
    color: #333;
    font-size: 0.9rem;
}

.handle {
    font-size: 0.85rem;
    color: #666;
}

.qr-icon {
    width: 32px;
    height: 32px;
    opacity: 0.5;
}

.real-qr {
    width: 180px;
    height: 180px;
    object-fit: contain;
    border-radius: 4px;
}

/* Cat Specific Styles */
.cat-card h2 {
    color: var(--purple);
}

.cat-card .illustration img {
    border-color: var(--purple-weak);
}

.cat-card .qr-placeholder {
    background-color: var(--purple-weak);
    color: var(--purple);
}

.btn-paypal {
    background-color: var(--purple);
}

/* Dog Specific Styles */
.dog-card h2 {
    color: var(--blue);
}

.dog-card .illustration img {
    border-color: var(--blue-weak);
    padding: 40px;
    box-sizing: border-box;
    background-color: var(--blue-weak);
    border-width: 6px;
    box-shadow: inset 0 0 20px rgba(0,0,0,0.05);
}

.dog-card .qr-placeholder {
    background-color: var(--blue-weak);
    color: var(--blue);
    padding: 24px;
}

.dog-card .qr-icon {
    width: 80px;
    height: 80px;
    opacity: 0.8;
    margin-bottom: 12px;
}

.btn-zelle {
    background-color: var(--blue); /* Or custom Zelle purple */
}

/* Adjusting Zelle to be distinct if reused */
.cat-card .btn-zelle {
     background-color: var(--purple);
}
.dog-card .btn-zelle {
     background-color: var(--blue);
}
.dog-card .btn-paypal {
     background-color: var(--blue);
}

.cat-card .btn-venmo {
     background-color: var(--purple);
}

.btn-venmo {
     background-color: var(--blue);
}

/* Mobile Responsiveness */
@media (max-width: 1024px) {
    .donate-container {
        flex-direction: column;
        align-items: center;
    }
}

@media (max-width: 480px) {
    .header h1 {
        font-size: 2rem;
    }
    .donate-card {
        padding: 24px;
    }
    .donate-card h2 {
        font-size: 1.5rem;
    }
    .illustration img {
        height: 150px;
    }
}
</style>
