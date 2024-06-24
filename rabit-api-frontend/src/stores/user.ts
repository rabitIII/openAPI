import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { getUser, register } from '@/apis/auth'

export const useUserStore = defineStore('user', {
  state() {
    return {
      user: getUser() || {},
    }
  },
  getters: {

  },
  actions: {
    // async registerUser( { userAccount, userPassword }: any) {
    //   const user = await register(userAccount, userPassword);
    // }
  },
})
