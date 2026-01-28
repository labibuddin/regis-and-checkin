<template>
  <div v-if="isOpen" class="fixed inset-0 z-[60] flex items-center justify-center p-4">
    <!-- Backdrop with blur -->
    <div class="absolute inset-0 bg-black/40 backdrop-blur-sm transition-opacity" @click="closeOnBackdrop"></div>

    <!-- Modal Content -->
    <div class="bg-white rounded-2xl shadow-2xl w-full max-w-sm overflow-hidden transform transition-all animate-scale-in relative z-10">
      
      <!-- Header / Icon -->
      <div :class="['p-6 flex flex-col items-center justify-center pb-2', headerColorClass]">
         <div :class="['w-16 h-16 rounded-full flex items-center justify-center mb-4 shadow-sm', iconBgClass]">
            <span class="text-3xl">{{ icon }}</span>
         </div>
         <h3 class="text-xl font-bold text-gray-800 text-center">{{ title }}</h3>
      </div>

      <!-- Body -->
      <div class="px-6 py-2 text-center">
        <p class="text-gray-600 leading-relaxed">{{ message }}</p>
      </div>

      <!-- Footer / Actions -->
      <div class="p-6 flex gap-3 justify-center">
        <button v-if="type === 'confirm'" 
                @click="$emit('cancel')" 
                class="flex-1 px-4 py-3 bg-gray-100 text-gray-700 font-bold rounded-xl hover:bg-gray-200 transition">
          {{ cancelText }}
        </button>
        
        <button @click="$emit('confirm')" 
                :class="['flex-1 px-4 py-3 text-white font-bold rounded-xl shadow-lg hover:shadow-xl transition transform active:scale-95', confirmBtnClass]">
          {{ confirmText }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps({
  isOpen: Boolean,
  title: String,
  message: String,
  type: {
    type: String,
    default: 'success' // success, error, confirm
  },
  confirmText: {
    type: String,
    default: 'OK'
  },
  cancelText: {
    type: String,
    default: 'Batal'
  }
})

const emit = defineEmits(['close', 'confirm', 'cancel'])

const closeOnBackdrop = () => {
    if (props.type === 'confirm') return // Force user to choose for confirm
    emit('close')
}

const icon = computed(() => {
    if (props.type === 'error') return '❌'
    if (props.type === 'confirm') return '❓'
    return '✅'
})

const iconBgClass = computed(() => {
    if (props.type === 'error') return 'bg-red-50 text-red-500'
    if (props.type === 'confirm') return 'bg-blue-50 text-blue-500'
    return 'bg-green-50 text-green-500'
})

const headerColorClass = computed(() => {
    return ''
})

const confirmBtnClass = computed(() => {
    if (props.type === 'error') return 'bg-red-600 hover:bg-red-700'
    if (props.type === 'confirm') return 'bg-blue-600 hover:bg-blue-700'
    return 'bg-green-600 hover:bg-green-700'
})
</script>

<style scoped>
.animate-scale-in {
  animation: scaleIn 0.3s cubic-bezier(0.16, 1, 0.3, 1);
}

@keyframes scaleIn {
  from { opacity: 0; transform: scale(0.9) translateY(10px); }
  to { opacity: 1; transform: scale(1) translateY(0); }
}
</style>
