import { defineStore } from "pinia";
import { ref, computed, type Ref } from "vue";

import { v4 as uuid } from 'uuid'

export type Message = {
    level : "info" | "warning" | "error"
    text: string
}

export const useMessageStore = defineStore('message', () => {
    const _messages :Ref<Map<string, Message>> = ref(new Map())

    const messages = computed( () => _messages.value)
    
    const pushMessage = (msg :Message) => {
        const key = uuid()

        _messages.value.set(key, msg)

        setTimeout(() => deleteMessage(key), 10 * 1000)
    }

    const deleteMessage = (key: string) => {
        _messages.value.delete(key)
    }
    
    return {
        messages, pushMessage, deleteMessage
    }
})