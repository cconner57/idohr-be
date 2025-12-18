<script setup lang="ts">
import InputField from '../../common/ui/InputField.vue'
import Toggle from '../../common/ui/Toggle.vue'

const { modelValue, catName } = defineProps<{
  catName: string | undefined
  modelValue: {
    firstName: string | null
    lastName: string | null
    age: number | null
    spouseFirstName: string | null
    spouseLastName: string | null
    roommatesNames: string[] | null
    childrenNamesAges: string[] | null
    email: string | null
    address: string | null
    streetAddressLine2: string | null
    city: string | null
    state: string | null
    zip: string | null
    phoneNumber: string | null
    cellPhoneNumber: string | null
    adultMembersAgreed: string | null
  }
}>()
console.log('Cat name in form:', catName)

function addRoommate() {
  console.log('Add roommate button clicked')
  // Logic to add additional roommate information can be implemented here
}
</script>

<template>
  <div class="cat-adoption-form">
    <div class="cat-name-display">
      <h2>Adopting Cat:</h2>
      <p>{{ catName }}</p>
    </div>
    <div>
      <h2>Applicant Information</h2>
      <div class="name-fields">
        <InputField
          v-model="modelValue.firstName"
          label="First Name of applicant"
          name="firstName"
          placeholder="Enter your first name"
          required
        />
        <InputField
          v-model="modelValue.lastName"
          label="Last Name of applicant"
          name="lastName"
          placeholder="Enter your last name"
          required
        />
      </div>
      <div class="other-fields">
        <InputField
          v-model="modelValue.age"
          label="Age of applicant"
          name="firstName"
          placeholder="Enter the age of the applicant"
          required
        />
        <InputField
          v-model="modelValue.email"
          label="Email of applicant"
          name="email"
          placeholder="Enter your email address"
          required
        />
      </div>
      <div class="address-fields">
        <InputField
          v-model="modelValue.address"
          label="Address of applicant"
          name="address"
          placeholder="Enter your street address"
          required
        />
        <InputField
          v-model="modelValue.streetAddressLine2"
          label="Street Address Line 2"
          name="streetAddressLine2"
          placeholder="Enter your address line 2"
          required
        />
      </div>
      <div class="address2-fields">
        <InputField
          v-model="modelValue.city"
          label="City of applicant"
          name="city"
          placeholder="Enter your city"
          required
        />
        <InputField
          v-model="modelValue.state"
          label="State of applicant"
          name="state"
          placeholder="Enter your state"
          required
        />
        <InputField
          v-model="modelValue.zip"
          label="Zip Code of applicant"
          name="zip"
          placeholder="Enter your zip code"
          required
        />
      </div>
      <div class="phone-numbers">
        <InputField
          v-model="modelValue.phoneNumber"
          label="Phone Number of applicant"
          name="phoneNumber"
          placeholder="Enter your phone number"
          required
        />
        <InputField
          v-model="modelValue.firstName"
          label="Cell Phone Number of applicant"
          name="firstName"
          placeholder="Enter your cell phone number"
          required
        />
      </div>
      <div class="spouse-fields">
        <InputField
          v-model="modelValue.spouseFirstName"
          label="Spouse/Parnter First Name"
          name="spouseFirstName"
          placeholder="Enter the first name of your spouse/partner"
          required
        />
        <InputField
          v-model="modelValue.spouseLastName"
          label="Spouse/Partner Last Name"
          name="spouseLastName"
          placeholder="Enter the last name of your spouse/partner"
          required
        />
        <div class="roommates">
          <div class="roommates-input">
            <InputField
              v-model="modelValue.firstName"
              label="Name of roommmate(s) and other adults in the household"
              name="firstName"
              placeholder="Enter the names of roommates/other adults"
              required
            />
          </div>
          <button class="add-roommate-btn" @click="addRoommate">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="currentColor"
              width="24"
              height="24"
            >
              <path
                d="M12 4.5c.414 0 .75.336.75.75v6h6c.414 0 .75.336.75.75s-.336.75-.75.75h-6v6c0 .414-.336.75-.75.75s-.75-.336-.75-.75v-6h-6c-.414 0-.75-.336-.75-.75s.336-.75.75-.75h6v-6c0-.414.336-.75.75-.75z"
              />
            </svg>
          </button>
        </div>
        <div class="children">
          <div class="children-input">
            <InputField
              v-model="modelValue.firstName"
              label="Name and ages of children (under 18) in the household"
              name="firstName"
              placeholder="Enter the names and ages of children in the household"
              required
            />
          </div>
          <button class="add-roommate-btn" @click="addRoommate">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="currentColor"
              width="24"
              height="24"
            >
              <path
                d="M12 4.5c.414 0 .75.336.75.75v6h6c.414 0 .75.336.75.75s-.336.75-.75.75h-6v6c0 .414-.336.75-.75.75s-.75-.336-.75-.75v-6h-6c-.414 0-.75-.336-.75-.75s.336-.75.75-.75h6v-6c0-.414.336-.75.75-.75z"
              />
            </svg>
          </button>
        </div>
      </div>
      <div class="all-agreed">
        <p>Have all adult members of the household agreed to this adoption?</p>
        <Toggle leftLabel="No" rightLabel="Yes" />
      </div>
    </div>
  </div>
</template>

<style scoped lang="css">
.cat-adoption-form {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
  padding: 0 0.5rem;
}
.cat-name-display {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
}
.cat-name-display h2 {
  font-size: 1.5rem;
  font-weight: bold;
  color: var(--font-color-dark);
}
.cat-name-display p {
  font-size: 1.25rem;
  font-weight: 600;
  color: var(--green);
}
.cat-adoption-form {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
  padding: 0 0.5rem;
  p {
    font-size: 1.75rem;
    font-weight: 600;
    text-align: center;
    color: var(--green);
  }
  .name-fields,
  .address-fields,
  .spouse-fields,
  .phone-numbers {
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    column-gap: 1rem;
  }
  .address2-fields {
    display: grid;
    grid-template-columns: 50% 27% 20%;
    column-gap: 1rem;
  }
  .other-fields {
    display: grid;
    grid-template-columns: 25% 73%;
    column-gap: 1rem;
  }
  .roommates,
  .children {
    display: flex;
    align-items: center;
    gap: 0.5rem;
  }
  .children-input,
  .roommates-input {
    flex-grow: 1;
  }
  .add-roommate-btn {
    width: 40px;
    background: none;
    border: 1px solid var(--border-color);
    border-radius: 10px;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    color: var(--green);
    padding: 0.5rem;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  }
  .add-roommate-btn:hover {
    color: var(--blue-hover);
    border-color: var(--blue-hover);
    transition:
      color 0.2s ease-in-out,
      border-color 0.2s ease-in-out;
  }
}
.all-agreed {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.5rem;
  text-align: center;
  margin: 1rem 0 2rem;
  p {
    font-size: 1.25rem;
    font-weight: 600;
    text-align: center;
    color: var(--green);
  }
}
</style>
