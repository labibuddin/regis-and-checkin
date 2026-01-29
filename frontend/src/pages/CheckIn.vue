<template>
  <div class="fixed inset-0 w-full h-[100dvh] bg-black flex flex-col z-50 overflow-hidden">
    <!-- Camera View -->
    <div class="relative flex-1 bg-black overflow-hidden flex items-center justify-center">
      <video ref="videoElem" class="absolute inset-0 w-full h-full object-cover"></video>
      
      <!-- Helpers -->
      <div v-if="!cameraError" class="absolute inset-0 flex items-center justify-center pointer-events-none z-10">
        <div class="w-64 h-64 border-2 border-white/50 rounded-3xl relative">
          <div class="absolute top-0 left-0 w-6 h-6 border-t-4 border-l-4 border-green-500 rounded-tl-xl"></div>
          <div class="absolute top-0 right-0 w-6 h-6 border-t-4 border-r-4 border-green-500 rounded-tr-xl"></div>
          <div class="absolute bottom-0 left-0 w-6 h-6 border-b-4 border-l-4 border-green-500 rounded-bl-xl"></div>
          <div class="absolute bottom-0 right-0 w-6 h-6 border-b-4 border-r-4 border-green-500 rounded-br-xl"></div>
        </div>
      </div>

      <!-- Error State -->
      <div v-if="cameraError" class="z-10 p-8 text-center text-white/80 max-w-sm mx-auto">
          <div class="text-4xl mb-4">📷🚫</div>
          <p class="text-lg font-medium breakdown-words">{{ cameraError }}</p>
      </div>

      <!-- Feedback Overlay -->
      <div v-if="feedback" :class="feedback.type === 'success' ? 'bg-green-500/90' : 'bg-red-500/90'" class="absolute inset-0 z-20 flex items-center justify-center flex-col text-white p-6 text-center animate-fade-in">
        <div class="text-6xl mb-4">{{ feedback.type === 'success' ? '✓' : '✕' }}</div>
        <h2 class="text-3xl font-bold mb-2">{{ feedback.title }}</h2>
        <p class="text-xl">{{ feedback.message }}</p>
      </div>
    </div>

    <!-- Manual Input Section -->
    <div class="bg-white p-6 pb-8 md:pb-6 rounded-t-3xl shadow-[0_-4px_20px_rgba(0,0,0,0.2)] z-10 safe-bottom">
      <form @submit.prevent="handleManualSubmit" class="flex gap-2 w-full">
        <input v-model="manualId" placeholder="Input ID / No HP" class="flex-1 min-w-0 bg-gray-100 border-none rounded-xl px-4 py-3 focus:ring-2 focus:ring-green-500 outline-none text-lg">
        <button type="submit" class="bg-green-600 active:scale-95 transition-transform text-white px-6 py-3 rounded-xl font-bold whitespace-nowrap shrink-0">
            Check-In
        </button>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import QrScanner from 'qr-scanner'
import api from '../services/api'

const videoElem = ref(null)
const scanner = ref(null)
const manualId = ref('')
const feedback = ref(null)
const isProcessing = ref(false)
const cameraError = ref('')

onMounted(() => {
  initScanner()
})

onUnmounted(() => {
  if (scanner.value) scanner.value.destroy()
})

const initScanner = async () => {
  try {
      scanner.value = new QrScanner(
        videoElem.value,
        result => handleScan(result.data),
        {
          highlightScanRegion: true,
          highlightCodeOutline: true,
          returnDetailedScanResult: true
        }
      )
      await scanner.value.start()
  } catch (err) {
      console.error(err)
      cameraError.value = "Kamera tidak dapat diakses. Pastikan menggunakan HTTPS atau localhost, dan izinkan akses kamera."
  }
}

const handleScan = async (data) => {
  if (isProcessing.value || feedback.value) return
  isProcessing.value = true
  
  await processCheckIn(data)
  
  isProcessing.value = false
}

const handleManualSubmit = async () => {
    if(!manualId.value) return
    isProcessing.value = true
    await processCheckIn(manualId.value)
    manualId.value = ''
    isProcessing.value = false
}

const processCheckIn = async (id) => {
  try {
    // Assuming backend handles ID lookup
    const res = await api.post('/check-in', { peserta_id: id })
    
    // Check 'new' flag to see if already checked in
    const isNew = res.data.data.new
    const name = res.data.data.peserta.panggilan
    const message = isNew ? `Selamat datang ${name}` : `${name} sudah check-in sebelumnya`
    
    showFeedback('success', 'Berhasil Check-in!', message)
  } catch (err) {
    showFeedback('error', 'Gagal', err.response?.data?.message || 'Peserta tidak ditemukan')
  }
}

const showFeedback = (type, title, message) => {
  feedback.value = { type, title, message }
  setTimeout(() => {
    feedback.value = null
  }, 3000)
}
</script>

<style scoped>
.animate-fade-in {
  animation: fadeIn 0.2s ease-out;
}
@keyframes fadeIn {
  from { opacity: 0; transform: scale(0.95); }
  to { opacity: 1; transform: scale(1); }
}
</style>
