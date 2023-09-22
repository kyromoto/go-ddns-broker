<script setup lang="ts">
import { ref, watch } from 'vue';
import { RouterView, RouterLink } from 'vue-router'

import { useMessageStore } from '@/stores/messageStore';

import Messages from '@/components/Messages.vue'
import router from '@/router';

const messageStore = useMessageStore()

const onBtnClickedNewRandomMsg = () => {
  messageStore.pushMessage({ text: (new Date()).toISOString(), level: "info" })
}

</script>

<template>

  <div class="layout-container">
    
    <header class="top-row-shadow">
      <div class="head-left"></div>
      <div class="head-center"></div>
      <div class="head-right">
        <button class="button is-small" @click="onBtnClickedNewRandomMsg">Add random message</button>
        <button class="button is-small is-warning">Logout</button>
      </div>
    </header>

    <div class="logo bg-darker">
      <RouterLink to="/" class="is-size-5">DDNS Broker</RouterLink>
    </div>
    
    <aside class="menu bg-darker">
      <p class="menu-label">
        General
      </p>
      <ul class="menu-list">
        <li><RouterLink to="Dashboard">Dashboard</RouterLink></li>
        <li><RouterLink to="Clients">Clients</RouterLink></li>
      </ul>
    </aside>

    <main>
      <Messages></Messages>
      <h1 class="is-size-4">{{ router.currentRoute.value.name }}</h1>
      <RouterView></RouterView>
    </main>
  </div>

</template>

<style scoped lang="scss">
  @import "../assets/scss/main.scss";

  $bg-darker: darken($background, 0.5%);

  div.layout-container {
    width: 100vw;
    height: 100vh;

    display: grid;

    grid-template-columns: minmax(15rem, auto) 1fr;
    grid-template-rows: auto 1fr;

    grid-template-areas:
      "logo header"
      "sidebar main"
    ;

    main {
      grid-area: main;

      padding: 1rem;

      position: relative;
    }

    header {
      grid-area: header;

      padding: 1rem;

      display: grid;

      grid-template-columns: auto 1fr auto;
      grid-template-rows: auto;

      grid-template-areas: "left center right";

      gap: 0.5rem;

      & > * {
        display: flex;
        gap: 0.5rem;
      }

      .head-left {
        grid-area: left;
      }

      .head-center {
        grid-area: center;
      }

      .head-right {
        grid-area: right;
      }
    }

    aside {
      grid-area: sidebar;
    }

    div.logo {
      grid-area: logo;

      display: flex;

      justify-content: center;
      align-items: center;
    }

    .menu {
      padding: 1rem;
    }

    .top-row-shadow {
      box-shadow: 0px 4px 12px rgba(0, 0, 0, 0.1) ;
    }

    .bg-darker {
      background-color: $bg-darker;
    }
  }
</style>