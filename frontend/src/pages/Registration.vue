<template>
  <div class="flex flex-col lg:flex-row min-h-screen">
    <!-- Left: Banner (Desktop 50%, Mobile Top) -->
    <!-- Left: Banner (Desktop 50%, Mobile Top) -->
    <div class="lg:w-1/2 bg-green-900 relative h-64 lg:h-auto overflow-hidden group">
      <!-- Sliding Backgrounds -->
      <div v-if="hasCustomBg" class="absolute inset-0 flex transition-transform duration-700 ease-in-out" 
           :style="{ transform: `translateX(-${activeBgIndex * 100}%)` }">
        <div v-for="id in bgImages" :key="id" class="min-w-full h-full relative">
             <img :src="`${api.defaults.baseURL}/backgrounds/${id}`" class="w-full h-full object-cover">
             <div class="absolute inset-0 bg-black/40"></div>
        </div>
      </div>
      
      <!-- Fallback Background -->
      <div v-else class="absolute inset-0 bg-green-700 opacity-20 pattern-grid-lg"></div>

      <!-- Static Text Overlay -->
      <div class="absolute inset-0 flex items-center justify-center p-8 text-white z-10 pointer-events-none">
        <div class="text-center">
           <h1 class="text-4xl font-bold mb-4 drop-shadow-lg">Masjid Jogokariyan</h1>
           <p class="text-xl opacity-90 drop-shadow-md">Mendoa untuk Indonesia Bahagia</p>
        </div>
      </div>

      <!-- Carousel Indicators -->
      <div v-if="hasCustomBg && bgImages.length > 1" class="absolute bottom-8 left-0 right-0 flex justify-center gap-2 z-20">
          <button v-for="(_, idx) in bgImages" :key="idx" 
                  @click="goToSlide(idx)"
                  class="h-1.5 rounded-full transition-all duration-300 shadow-sm cursor-pointer"
                  :class="idx === activeBgIndex ? 'w-8 bg-white' : 'w-2 bg-white/50 hover:bg-white/80'">
          </button>
      </div>
    </div>

    <!-- Right: Form Area -->
    <div class="lg:w-1/2 flex flex-col justify-center p-6 lg:p-12 bg-white">
      
      <!-- SUCCESS STATE -->
      <div v-if="successData" class="max-w-md mx-auto w-full text-center space-y-6">
        <div ref="cardRef" class="bg-green-50 p-6 rounded-2xl border border-green-100 shadow-sm">
          <h2 class="text-2xl font-bold text-green-700 mb-2">Selamat!</h2>
          <p class="text-gray-600 mb-6">Anda terdaftar sebagai peserta Kajian Masjid Jogokariyan.</p>
          
          <div class="flex justify-center mb-4">
             <canvas id="qrcode-canvas"></canvas>
          </div>
          
          <p class="font-mono text-lg font-bold text-gray-800">{{ successData?.nama_lengkap }}</p>
          <p class="text-sm text-gray-500">{{ successData?.id }}</p>
        </div>
        
        <button @click="downloadCard" class="w-full py-3 bg-blue-600 text-white rounded-xl font-semibold hover:bg-blue-700 transition flex justify-center items-center gap-2">
           ⬇️ Download E-Card
        </button>

        <div class="bg-blue-50 p-6 rounded-2xl border border-blue-100 text-left">
          <h3 class="font-bold text-blue-800 mb-2">Undangan Kajian</h3>
          <div v-if="activeEvent">
            <p class="text-xl font-bold text-gray-900 mb-1">{{ activeEvent.nama_kegiatan }}</p>
            <p class="text-gray-700 flex items-center gap-2">
               📅 {{ new Date(activeEvent.waktu).toLocaleDateString('id-ID', { weekday: 'long', day: 'numeric', month: 'long', year: 'numeric' }) }}
            </p>
            <p class="text-gray-700 flex items-center gap-2">
               ⏰ {{ new Date(activeEvent.waktu).toLocaleTimeString('id-ID', { hour: '2-digit', minute: '2-digit' }) }} WIB - Selesai
            </p>
            <p class="text-gray-700 mt-2">
               🎙️ Bersama <span class="font-bold">{{ activeEvent.pemateri }}</span>
            </p>
            <p class="text-gray-600 text-sm mt-1">📍 {{ activeEvent.tempat }}</p>
          </div>
          <p v-else class="text-gray-700">Silahkan datang ke Masjid Jogokariyan nanti.</p>
        </div>

        <button @click="resetFlow" class="w-full py-3 bg-gray-200 rounded-xl font-semibold hover:bg-gray-300 transition">
          Kembali
        </button>
      </div>

      <!-- INITIAL CHOICE -->
      <div v-else-if="flow === 'choice'" class="max-w-md mx-auto w-full space-y-8 animate-fade-in">
        <h2 class="text-3xl font-bold text-gray-900 leading-tight">Selamat Datang</h2>
        <p class="text-gray-600">Sudah pernah registrasi kajian di Masjid Jogokariyan?</p>
        
        <div class="space-y-4">
          <button @click="flow = 'existing'" class="w-full py-4 border-2 border-green-600 text-green-700 rounded-xl font-bold text-lg hover:bg-green-50 transition">
            Sudah Pernah
          </button>
          <button @click="flow = 'new'" class="w-full py-4 bg-green-600 text-white rounded-xl font-bold text-lg hover:bg-green-700 shadow-lg hover:shadow-xl transition transform active:scale-95">
            Belum Pernah (Daftar Baru)
          </button>
        </div>
      </div>

      <!-- EXISTING USER FLOW -->
      <div v-else-if="flow === 'existing'" class="max-w-md mx-auto w-full space-y-6">
        <button @click="flow = 'choice'" class="text-gray-400 hover:text-gray-600 mb-4">&larr; Kembali</button>
        <h2 class="text-2xl font-bold">Cari Data Peserta</h2>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">No. HP / ID Peserta</label>
            <input v-model="searchQuery" @keyup.enter="searchUser" type="text" class="w-full px-4 py-3 rounded-xl border border-gray-300 focus:ring-2 focus:ring-green-500 outline-none" placeholder="08123456789">
          </div>
          <button @click="searchUser" :disabled="loading" class="w-full py-3 bg-green-600 text-white rounded-xl font-bold hover:bg-green-700 transition">
            <span v-if="loading">Mencari...</span>
            <span v-else>Cari Data</span>
          </button>
        </div>
        
        <p v-if="searchError" class="text-red-500 bg-red-50 p-3 rounded-lg text-sm text-center">
          {{ searchError }}
          <button @click="flow = 'new'" class="underline font-bold ml-1">Registrasi Sekarang</button>
        </p>
      </div>

      <!-- NEW REGISTRATION FLOW -->
      <form v-else-if="flow === 'new'" @submit.prevent="submitRegister" class="max-w-md mx-auto w-full space-y-4">
        <div class="flex items-center justify-between mb-2">
           <button type="button" @click="flow = 'choice'" class="text-gray-400 hover:text-gray-600">&larr; Kembali</button>
           <h2 class="text-xl font-bold">Formulir Pendaftaran</h2>
        </div>

        <div class="grid grid-cols-2 gap-4">
           <div>
             <label class="block text-xs font-bold text-gray-500 uppercase">Nama Lengkap</label>
             <input v-model="form.nama_lengkap" required class="w-full p-3 rounded-lg border border-gray-300 focus:ring-green-500">
           </div>
           <div>
             <label class="block text-xs font-bold text-gray-500 uppercase">Panggilan</label>
             <input v-model="form.panggilan" class="w-full p-3 rounded-lg border border-gray-300 focus:ring-green-500">
           </div>
        </div>

        <div>
             <label class="block text-xs font-bold text-gray-500 uppercase">Jenis Kelamin</label>
             <select v-model="form.jenis_kelamin" required class="w-full p-3 rounded-lg border border-gray-300">
               <option value="Laki-laki">Laki-laki</option>
               <option value="Perempuan">Perempuan</option>
             </select>
        </div>

        <div>
           <label class="block text-xs font-bold text-gray-500 uppercase">Tahun Lahir</label>
           <input v-model="form.tahun_lahir" type="number" required min="1900" :max="new Date().getFullYear()" class="w-full p-3 rounded-lg border border-gray-300 focus:ring-green-500">
        </div>

        <div>
           <label class="block text-xs font-bold text-gray-500 uppercase">Alamat Domisili</label>
           <textarea v-model="form.alamat_domisili" rows="2" class="w-full p-3 rounded-lg border border-gray-300"></textarea>
        </div>

        <div>
           <label class="block text-xs font-bold text-gray-500 uppercase">No. Handphone (WA)</label>
           <input v-model="form.no_hp" type="tel" required @input="debouncedCheck" :class="{'border-red-500': validationErrors.no_hp}" class="w-full p-3 rounded-lg border border-gray-300">
           <p v-if="validationErrors.no_hp" class="text-xs text-red-500 mt-1">{{ validationErrors.no_hp }}</p>
        </div>

        <div>
           <label class="block text-xs font-bold text-gray-500 uppercase">Email</label>
           <input v-model="form.email" type="email" @input="debouncedCheck" :class="{'border-red-500': validationErrors.email}" class="w-full p-3 rounded-lg border border-gray-300">
           <p v-if="validationErrors.email" class="text-xs text-red-500 mt-1">{{ validationErrors.email }}</p>
        </div>
        
        <div class="pt-6 border-t border-gray-100 space-y-4">
           <!-- Campaign Text -->
           <div class="bg-green-50 p-4 rounded-xl border border-green-100">
              <h3 class="font-bold text-green-800 text-sm mb-1">Mari Tebarkan Kebaikan di Bulan Penuh Berkah</h3>
              <p class="text-xs text-green-700 mb-2">Ramadhan adalah saat terbaik untuk berbagi kebahagiaan 🌙</p>
              <p class="text-xs text-gray-600 leading-relaxed">
                 Masjid Jogokariyan membuka kesempatan bagi jamaah yang berkenan menyalurkan infaq untuk 
                 <span class="font-bold">Program Kegiatan Masjid Jogokariyan</span> melalui QRIS berikut (opsional):
              </p>
              
              <!-- Dynamic QRIS Image -->
              <div class="mt-4 bg-white p-2 rounded-lg border border-gray-200 shadow-sm">
                 <img :src="`${api.defaults.baseURL}/images/qris`" 
                      alt="QRIS Masjid Jogokariyan" 
                      class="w-full h-auto rounded-md"
                      @error="(e: Event) => (e.target as HTMLImageElement).style.display = 'none'">
              </div>
           </div>

           <div>
             <p class="text-sm font-bold text-gray-700 mb-2">Unggah bukti transfer di sini</p>
             <input type="file" ref="fileInput" accept="image/*" class="block w-full text-sm text-gray-500
                file:mr-4 file:py-2 file:px-4
                file:rounded-full file:border-0
                file:text-sm file:font-semibold
                file:bg-green-50 file:text-green-700
                hover:file:bg-green-100
              " />
             <p class="text-xs text-gray-400 mt-1">Maks. 1MB (Opsional)</p>
           </div>
        </div>

        <button type="submit" :disabled="loading || Object.keys(validationErrors).length > 0" class="w-full py-4 mt-6 bg-green-600 text-white rounded-xl font-bold shadow-lg hover:bg-green-700 transition disabled:opacity-50 disabled:cursor-not-allowed">
           {{ loading ? 'Mendaftarkan...' : 'Daftar Sekarang' }}
        </button>

      </form>

    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, nextTick, computed, onMounted } from 'vue'
import api from '../services/api'
import QRCode from 'qrcode'

interface Peserta {
    id: string
    nama_lengkap: string
    [key: string]: any
}

interface EventData {
    nama_kegiatan: string
    waktu: string
    pemateri: string
    tempat: string
}

type FlowState = 'choice' | 'existing' | 'new'

const flow = ref<FlowState>('choice') 
const loading = ref(false)
const searchQuery = ref('')
const searchError = ref('')
const successData = ref<Peserta | null>(null)
const activeBgIndex = ref(0)
const bgImages = ref<string[]>([])
const hasCustomBg = ref(false)
const activeEvent = ref<EventData | null>(null)

const bannerStyle = computed(() => ({})) // Deprecated but kept to avoid breakage if ref still exists, effectively unused
let carouselInterval: ReturnType<typeof setInterval> | null = null

const startCarousel = () => {
    if (carouselInterval) clearInterval(carouselInterval)
    carouselInterval = setInterval(() => {
        activeBgIndex.value = (activeBgIndex.value + 1) % bgImages.value.length
    }, 5000)
}

const goToSlide = (idx: number) => {
    activeBgIndex.value = idx
    startCarousel() // Reset timer on manual interaction
}

// Check if custom background exists
onMounted(async () => {
    // 1. Fetch Background Gallery
    try {
        const res = await api.get('/backgrounds')
        if (res.data.data && res.data.data.length > 0) {
            bgImages.value = res.data.data
            hasCustomBg.value = true
            
            // Start Carousel if more than 1 image
            if (bgImages.value.length > 1) {
                startCarousel()
            }
        }
    } catch (e) { /* ignore */ }

    // 2. Fetch Active Event to display on success
    try {
        const res = await api.get('/event/active')
        if (res.data.data && res.data.data.active) {
           activeEvent.value = res.data.data.event
        }
    } catch (e) {
        console.log("No active event found")
    }
})

const form = reactive({
  nama_lengkap: '',
  panggilan: '',
  jenis_kelamin: 'Laki-laki',
  alamat_domisili: '',
  no_hp: '',
  email: '',
  tahun_lahir: 2002
})

const fileInput = ref<HTMLInputElement | null>(null)
const validationErrors = reactive<Record<string, string>>({})

// Debounce logic
let timeout: ReturnType<typeof setTimeout> | null = null
const debouncedCheck = (e: Event) => {
  const target = e.target as HTMLInputElement
  if (timeout) clearTimeout(timeout)
  timeout = setTimeout(() => {
    checkUniqueness(target.value)
  }, 1000)
}

const checkUniqueness = async (val: string) => {
  if (!val) return
  try {
    const res = await api.get(`/check-existing?query=${val}`)
    // API returns 409 if exists, 200 if available
    // Actually my backend returns 200 for available, 409 for exists error json.
    // Axios throws on 409 usually
  } catch (err: any) {
    if (err.response && err.response.status === 409) {
      if (val.includes('@')) validationErrors.email = "Email sudah digunakan"
      else validationErrors.no_hp = "No HP sudah digunakan"
    }
  }
}
// Clean errors on input start? Maybe. Simplified for now.

const submitRegister = async () => {
  loading.value = true
  try {
    const formData = new FormData()
    // Type assertion for keys to avoid string indexing error if needed
    ;(Object.keys(form) as Array<keyof typeof form>).forEach(key => formData.append(key, String(form[key])))
    if (fileInput.value && fileInput.value.files && fileInput.value.files[0]) {
       // Check size
       if (fileInput.value.files[0].size > 1024 * 1024) {
         alert("File too large > 1MB")
         loading.value = false
         return
       }
       formData.append('bukti_transfer', fileInput.value.files[0])
    }

    const res = await api.post('/register', formData, {
      headers: { 'Content-Type': 'multipart/form-data' }
    })
    
    handleSuccess(res.data.data)
  } catch (err: any) {
    alert(err.response?.data?.message || "Registration failed")
  } finally {
    loading.value = false
  }
}

const searchUser = async () => {
  loading.value = true
  searchError.value = ''
  try {
    const res = await api.get(`/peserta?query=${searchQuery.value}`)
    handleSuccess(res.data.data)
  } catch (err) {
    searchError.value = "Data peserta tidak ditemukan."
  } finally {
    loading.value = false
  }
}

const handleSuccess = (data: Peserta) => {
  successData.value = data
  nextTick(() => {
    const canvas = document.getElementById('qrcode-canvas')
    if (canvas) {
        QRCode.toCanvas(canvas, data.id, function (error) {
           if (error) console.error(error)
        })
    }
  })
}

const resetFlow = () => {
  successData.value = null
  flow.value = 'choice'
  // reset form...
}

// Download Card Logic
import html2canvas from 'html2canvas'
const cardRef = ref(null)

const downloadCard = async () => {
  if (!cardRef.value || !successData.value) return
  try {
    const canvas = await html2canvas(cardRef.value, {
       scale: 2, // better quality
       backgroundColor: '#f0fdf4', // match bg-green-50
       logging: false
    })
    const link = document.createElement('a')
    link.download = `Kartu-Peserta-${successData.value.nama_lengkap}.png`
    link.href = canvas.toDataURL('image/png')
    link.click()
  } catch (err) {
    console.error("Failed to capture card", err)
    alert("Gagal mendownload kartu")
  }
}
</script>

<style scoped>
.animate-fade-in {
  animation: fadeIn 0.5s ease-out;
}
@keyframes fadeIn {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}
</style>
