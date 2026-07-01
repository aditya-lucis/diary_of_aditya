<template>
  <div id="app" class="min-h-screen bg-navy-900">
    <div class="flex items-center justify-center min-h-screen">
      <div class="text-center">
        <!-- Logo -->
        <div class="mb-8">
          <div class="inline-flex items-center justify-center w-24 h-24 bg-navy-800 rounded-2xl shadow-2xl mb-4">
            <span class="text-5xl font-bold text-gold-400">A</span>
          </div>
        </div>
        
        <!-- Title -->
        <h1 class="text-4xl font-bold text-white mb-2">
          Diary Of Aditya
        </h1>
        <p class="text-gray-400 mb-8">
          IT Consultant Management System
        </p>

        <!-- Status -->
        <div class="inline-flex items-center gap-2 px-4 py-2 bg-navy-800 rounded-full">
          <div class="w-2 h-2 bg-green-500 rounded-full animate-pulse"></div>
          <span class="text-sm text-gray-300">System Ready</span>
        </div>

        <!-- Quick Links -->
        <div class="mt-8 flex gap-4 justify-center">
          <button 
            @click="navigateTo('/login')"
            class="px-6 py-3 bg-primary-600 hover:bg-primary-700 text-white font-semibold rounded-lg transition-colors"
          >
            Login
          </button>
          <button 
            @click="navigateTo('/register')"
            class="px-6 py-3 bg-transparent border-2 border-primary-600 text-primary-400 hover:bg-primary-600 hover:text-white font-semibold rounded-lg transition-colors"
          >
            Register
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()

const navigateTo = (path) => {
  router.push(path)
}
</script>

<style scoped>
#app {
  background: linear-gradient(135deg, #0f172a 0%, #1e293b 100%);
}
</style>

<style>
#app {
  background: linear-gradient(135deg, #020617 0%, #0f172a 50%, #1e293b 100%);
  min-height: 100vh;
}
</style>