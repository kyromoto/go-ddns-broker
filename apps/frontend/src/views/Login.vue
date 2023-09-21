<script setup lang="ts">
import { ref, watch } from 'vue';

import { useUserStore } from '@/stores/userStore'
import router from '@/router';

const userStore = useUserStore()

const isFormDisabled = ref(false)

const username = ref("")
const password = ref("")

const onFormSubmit = async () => {
  console.error("login not implemented")

  isFormDisabled.value = true

  await userStore.login(username.value, password.value)

  isFormDisabled.value = false
}

const onFormReset = () => {
  username.value = ""
  password.value = ""
}

</script>

<template>
  <form class="login-form-container" @submit.prevent="onFormSubmit" @reset.prevent="onFormReset">
    <div class="field">
      <label class="label">Username</label>
      <div class="control">
        <input class="input" type="text" placeholder="Your username" v-model="username" :disabled="isFormDisabled" required>
      </div>
    </div>

    <div class="field">
      <label class="label">Password</label>
      <div class="control">
        <input class="input" type="password" placeholder="Your password" v-model="password" :disabled="isFormDisabled" required>
      </div>
    </div>

    <div class="field is-grouped">
      <div class="control">
        <button type="submit" class="button is-link" :disabled="isFormDisabled">Login</button>
      </div>
      <div class="control">
        <button type="reset" class="button is-link is-light" :disabled="isFormDisabled">Reset</button>
      </div>
    </div>
  </form>
</template>

<style scoped lang="scss">
  .login-form-container {

    display: flex;

    flex-direction: column;

    align-items: center;
    justify-content: center;

    box-shadow: 0px 4px 12px rgba(0, 0, 0, 0.1);

    padding: 3rem 4rem;
  }

  .w-100 {
    width: 100%;
  }
</style>