import { ref, computed, type Ref, type ComputedRef } from 'vue'
import { defineStore } from 'pinia'

export interface UserService {
  AuthenticateUserAsync(username: string, password: string) :Promise<boolean>
}

class UserServiceImpl implements UserService {
  AuthenticateUserAsync(username: string, password: string): Promise<boolean> {
    throw new Error('Method not implemented.')
  }
  
}

export const useUserStore = defineStore('user', () => {
  const _isLoggedIn :Ref<boolean> = ref(true)
  
  const isLoggedIn :ComputedRef<boolean> = computed(() => _isLoggedIn.value)

  const login = async (username :string, password :string) => {
    const userService = new UserServiceImpl()

    _isLoggedIn.value = await userService.AuthenticateUserAsync(username, password)
  }

  return { isLoggedIn, login }
})
