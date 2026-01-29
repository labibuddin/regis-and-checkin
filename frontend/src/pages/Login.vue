<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import api from '../services/api';

const username = ref('');
const password = ref('');
const error = ref('');
const loading = ref(false);
const router = useRouter();

const handleLogin = async () => {
    loading.value = true;
    error.value = '';
    
    try {
        const response = await api.post('/login', {
            username: username.value,
            password: password.value
        });

        localStorage.setItem('token', response.data.token);
        router.push('/admin');
    } catch (e: any) {
        if (e.response && e.response.status === 401) {
            error.value = 'Invalid username or password';
        } else {
            error.value = 'An error occurred. Please try again.';
        }
    } finally {
        loading.value = false;
    }
};
</script>

<template>
    <div class="min-h-screen flex items-center justify-center bg-slate-900 relative overflow-hidden">
        <!-- Background Elements -->
        <div class="absolute top-0 left-0 w-full h-full overflow-hidden z-0">
            <div class="absolute -top-[20%] -left-[10%] w-[70%] h-[70%] bg-emerald-500/20 rounded-full blur-[100px] animate-pulse"></div>
            <div class="absolute top-[30%] -right-[10%] w-[50%] h-[60%] bg-blue-500/10 rounded-full blur-[100px]"></div>
        </div>

        <!-- Login Card -->
        <div class="relative z-10 w-full max-w-md p-8 bg-white/5 backdrop-blur-xl border border-white/10 rounded-3xl shadow-2xl">
            <div class="text-center mb-10">
                <h1 class="text-3xl font-bold text-white mb-2">Welcome Back</h1>
                <p class="text-gray-400">Sign in to access the admin dashboard</p>
            </div>

            <form @submit.prevent="handleLogin" class="space-y-6">
                <div>
                    <label class="block text-sm font-medium text-gray-300 mb-2">Username</label>
                    <input 
                        v-model="username"
                        type="text" 
                        class="w-full px-4 py-3 bg-white/5 border border-white/10 rounded-xl text-white placeholder-gray-500 focus:outline-none focus:ring-2 focus:ring-emerald-500/50 focus:border-emerald-500/50 transition duration-200"
                        placeholder="Enter username"
                    />
                </div>

                <div>
                    <label class="block text-sm font-medium text-gray-300 mb-2">Password</label>
                    <input 
                        v-model="password"
                        type="password" 
                        class="w-full px-4 py-3 bg-white/5 border border-white/10 rounded-xl text-white placeholder-gray-500 focus:outline-none focus:ring-2 focus:ring-emerald-500/50 focus:border-emerald-500/50 transition duration-200"
                        placeholder="Enter password"
                    />
                </div>

                <div v-if="error" class="p-4 bg-red-500/10 border border-red-500/20 rounded-xl">
                    <p class="text-red-400 text-sm text-center">{{ error }}</p>
                </div>

                <button 
                    type="submit" 
                    :disabled="loading"
                    class="w-full py-3.5 bg-gradient-to-r from-emerald-500 to-teal-500 hover:from-emerald-400 hover:to-teal-400 text-white font-bold rounded-xl shadow-lg shadow-emerald-500/20 transform transition-all duration-200 hover:-translate-y-0.5 active:translate-y-0 disabled:opacity-50 disabled:cursor-not-allowed flex justify-center items-center"
                >
                    <svg v-if="loading" class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                    </svg>
                    <span>{{ loading ? 'Signing in...' : 'Sign In' }}</span>
                </button>
            </form>
        </div>
        
        <div class="absolute bottom-6 text-center w-full z-10">
             <a href="/" class="text-gray-500 hover:text-emerald-400 transition text-sm">← Back to Registration</a>
        </div>
    </div>
</template>
