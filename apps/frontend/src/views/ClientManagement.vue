<script setup lang="ts">
import { ref } from 'vue';

const clients = ref([{ "description" : "k-rm-28" }, { "description" : "gm-m-10" }])
</script>

<template>
    <div class="client-mgmt-layout-container">
        <div class="left">
            <ul class="list">
                <li v-for="(client, index) of clients" :key="index">{{ client.description }}</li>
            </ul>
        </div>

        <div class="right">

        </div>
    </div>
</template>

<style scoped lang="scss">
    .client-mgmt-layout-container {
        display: grid;

        width: 100%;
        height: 100%;

        grid-template-columns: 1fr 2fr;
        grid-template-rows: 1fr;

        grid-template-areas: "left right";

        .left {
            grid-area: left;

            overflow-y: auto;
        }

        .right {
            grid-area: right;
        }
    }

    ul.list {
        // border: 1px solid grey;

        li {
            padding: 1rem;

            border: 1px solid grey;

            &:first-child {
                border-top-left-radius: 6px;
                border-top-right-radius: 6px;
            }

            &:last-child {
                border-bottom-left-radius: 6px;
                border-bottom-right-radius: 6px;
            }

            &:not(:last-child) {
                border-bottom: 0;
            }

            &:hover {
                background-color: grey;
            }

            &:nth-child(2) {
                background-color: #f6f6f6;
            }
        }
    }
</style>