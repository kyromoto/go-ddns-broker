<script setup lang="ts">
import { useMessageStore } from '@/stores/messageStore'

const messageStore = useMessageStore()

</script>

<template>
    <div class="message-grid-container">
        <div class="message-container">
            <div class="notification is-expanded" v-for="[id, message] of messageStore.messages" :key="id" :class="{ 'is-info': message.level == 'info', 'is-warning': message.level == 'warning', 'is-error': message.level == 'error' }" >
                <button class="delete" @click="messageStore.deleteMessage(id)"></button>
                {{ message.text }}
            </div>
        </div>
    </div>
</template>

<style scoped lang="scss">
.message-grid-container {

    z-index: 100;

    position: absolute;

    left: 0;
    right: 0;
    top: 0;
    bottom: 0;

    display: grid;
    grid-template-columns: 1fr 30rem;
    grid-template-rows: 1fr;

    grid-template-areas: ". messages";

    .message-container {

        grid-area: messages;

        padding: 1rem;

        overflow-y: auto;
        .notification {

            box-shadow: rgba(0, 0, 0, 0.24) 0px 3px 8px;

            width: 100%;
        }
    }
}
</style>