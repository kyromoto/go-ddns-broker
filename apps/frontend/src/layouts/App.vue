<script setup lang="ts">
import { RouterView, RouterLink } from 'vue-router'
</script>

<template>
  <div class="layout-container">
    
    <header class="top-row-shadow">
      <div class="head-left"></div>
      <div class="head-center"></div>
      <div class="head-right">
        <button class="button is-small is-warning">Logout</button>
      </div>
    </header>

    <div class="bg-darker">
      <RouterLink to="/" class="is-size-5">DDNS Broker</RouterLink>
    </div>
    
    <aside class="menu bg-darker">
      <p class="menu-label">
        General
      </p>
      <ul class="menu-list">
        <li><RouterLink to="Dashboard">Dashboard</RouterLink></li>
      </ul>
      
    </aside>

    <main>
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
    }

    header {
      grid-area: header;

      padding: 1rem;

      display: grid;

      grid-template-columns: auto 1fr auto;
      grid-template-rows: auto;

      grid-template-areas: "left center right";

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

    div {
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