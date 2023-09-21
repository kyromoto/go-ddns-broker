<script setup lang="ts">
import { RouterView } from 'vue-router'
</script>

<template>
  <div class="layout-container">
    <RouterView></RouterView>
  </div>
</template>

<style scoped lang="scss">
  .layout-container {
    width: 100vw;
    height: 100vh;

    display: grid;

    place-content: center;
  }
</style>