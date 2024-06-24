import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useInterfaceInfoStore = defineStore('interfaceinfo', {
  state() {
    return {
      name: "",
      descrition: "",
      url: "",
      method: "",
      requestHeader: "",
      responseHeader: "",
      requestParams: "",
      status: "",
    }
  },
  actions: {

  },
  getters: {
    
  }
})
