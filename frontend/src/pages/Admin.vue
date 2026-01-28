<template>
  <div class="flex min-h-screen bg-gray-100 font-sans">
    <!-- Sidebar -->
    <aside :class="['bg-gray-900 text-white flex-col fixed inset-y-0 z-30 transition-all duration-300 transform', isSidebarOpen ? 'translate-x-0 w-64' : '-translate-x-full w-64 md:translate-x-0 md:static md:flex']">
      <div class="p-6 font-bold text-xl tracking-wider border-b border-gray-800 flex justify-between items-center">
        <span>ADMIN</span>
        <button @click="isSidebarOpen = false" class="md:hidden text-gray-400 hover:text-white">✕</button>
      </div>
      <nav class="flex-1 p-4 space-y-2">
        <button @click="setView('dashboard')" :class="['w-full text-left px-4 py-3 rounded-lg transition', currentView === 'dashboard' ? 'bg-green-600 text-white' : 'text-gray-400 hover:bg-gray-800']">
          Dashboard
        </button>
        <button @click="setView('peserta')" :class="['w-full text-left px-4 py-3 rounded-lg transition', currentView === 'peserta' ? 'bg-green-600 text-white' : 'text-gray-400 hover:bg-gray-800']">
          Data Peserta
        </button>
        <button @click="setView('kegiatan')" :class="['w-full text-left px-4 py-3 rounded-lg transition', currentView === 'kegiatan' ? 'bg-green-600 text-white' : 'text-gray-400 hover:bg-gray-800']">
          Kegiatan
        </button>
        <button @click="setView('settings')" :class="['w-full text-left px-4 py-3 rounded-lg transition', currentView === 'settings' ? 'bg-green-600 text-white' : 'text-gray-400 hover:bg-gray-800']">
          Settings
        </button>
      </nav>
      <div class="p-4 border-t border-gray-800">
        <router-link to="/check-in" class="block w-full text-center py-3 bg-indigo-600 rounded-lg font-bold hover:bg-indigo-500">
          Open Scanner
        </router-link>
      </div>
    </aside>

    <!-- Overlay for mobile -->
    <div v-if="isSidebarOpen" @click="isSidebarOpen = false" class="fixed inset-0 bg-black/50 z-20 md:hidden"></div>

    <!-- Main Content -->
    <main class="flex-1 flex flex-col min-w-0 overflow-hidden">
      <!-- Top header for mobile -->
      <header class="flex items-center justify-between p-4 bg-white shadow-sm md:hidden">
        <div class="font-bold text-gray-800">Admin Panel</div>
        <button @click="isSidebarOpen = true" class="p-2 text-gray-600 rounded hover:bg-gray-100">
           ☰ Menu
        </button>
      </header>

      <div class="flex-1 overflow-auto p-4 md:p-8">
        
        <!-- DASHBOARD VIEW -->
        <div v-if="currentView === 'dashboard'" class="animate-fade space-y-8">
          <div class="flex justify-between items-center">
             <div class="flex items-center gap-3">
                 <h1 class="text-3xl font-bold text-gray-800">Dashboard Overview</h1>
                 <span class="flex items-center gap-1.5 px-3 py-1 bg-green-100 text-green-700 text-xs font-bold rounded-full animate-pulse">
                    <span class="w-2 h-2 bg-green-500 rounded-full"></span>
                    Live Updates
                 </span>
             </div>
          </div>
          
          <!-- Top Stats Cards -->
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
            <!-- Total Peserta -->
            <div class="bg-white p-6 rounded-2xl shadow-sm border border-gray-100 flex flex-col justify-between">
              <div>
                 <h3 class="text-gray-500 font-medium mb-1">Total Peserta</h3>
                 <p class="text-4xl font-bold text-gray-900">{{ dashboardStats.total_peserta }}</p>
              </div>
              <div class="text-right mt-2">
                 <span class="text-xs text-green-600 bg-green-50 px-2 py-1 rounded-full">+{{ dashboardStats.total_peserta_new || 0 }} minggu ini</span>
              </div>
            </div>

            <!-- Hadir Hari Ini -->
            <div class="bg-white p-6 rounded-2xl shadow-sm border border-gray-100 flex flex-col justify-between">
              <div>
                  <h3 class="text-gray-500 font-medium mb-1">Hadir Hari Ini</h3>
                  <p class="text-4xl font-bold text-green-600">{{ dashboardStats.total_checkin_today }}</p>
              </div>
              <div class="w-full bg-gray-100 rounded-full h-1.5 mt-4">
                 <div class="bg-green-500 h-1.5 rounded-full" style="width: 40%"></div>
              </div>
            </div>

             <!-- Dominasi Gender -->
             <div class="bg-white p-6 rounded-2xl shadow-sm border border-gray-100">
               <h3 class="text-gray-500 font-medium mb-2">Dominasi Gender</h3>
               <div class="flex items-end gap-4 mt-2">
                  <div class="flex-1">
                      <span class="text-3xl font-bold text-blue-600">{{ dashboardStats.gender_stats?.laki_laki }}</span>
                      <span class="text-xs text-gray-400 block font-bold uppercase tracking-wider">Laki-laki</span>
                  </div>
                  <div class="flex-1 text-right">
                      <span class="text-3xl font-bold text-pink-500">{{ dashboardStats.gender_stats?.perempuan }}</span>
                      <span class="text-xs text-gray-400 block font-bold uppercase tracking-wider">Perempuan</span>
                  </div>
               </div>
            </div>

            <!-- Event Aktif -->
            <div class="bg-gradient-to-br from-indigo-600 to-purple-700 p-6 rounded-2xl shadow-lg run border border-indigo-500 text-white relative overflow-hidden">
                <div v-if="dashboardStats.active_event">
                    <h3 class="text-indigo-100 text-sm font-bold uppercase tracking-wider mb-2">Event Hari Ini</h3>
                    <p class="text-xl font-bold leading-tight mb-2">{{ dashboardStats.active_event.nama_kegiatan }}</p>
                    <p class="text-indigo-100 text-xs mb-1">⏰ {{ dashboardStats.active_event.waktu }}</p>
                    <p class="text-indigo-100 text-xs">🎙️ {{ dashboardStats.active_event.pemateri }}</p>
                </div>
                <div v-else class="flex flex-col items-center justify-center h-full text-indigo-200">
                    <span class="text-2xl mb-2">💤</span>
                    <span class="text-sm font-bold">Tidak ada event aktif</span>
                </div>
                <!-- Decorative circle -->
                <div class="absolute -bottom-4 -right-4 w-24 h-24 bg-white opacity-10 rounded-full blur-xl"></div>
            </div>
          </div>

          <!-- Charts Section -->
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
              <!-- Age Distribution -->
              <div class="bg-white p-6 rounded-xl shadow-sm border border-gray-100">
                  <h3 class="font-bold text-gray-800 mb-4">Sebaran Usia Peserta</h3>
                  <div class="h-64">
                      <Bar v-if="ageChartData" :data="ageChartData" :options="chartOptions as any" />
                  </div>
              </div>

               <!-- Participant Trend -->
              <div class="bg-white p-6 rounded-xl shadow-sm border border-gray-100">
                  <h3 class="font-bold text-gray-800 mb-4">Tren Pendaftar (7 Hari Terakhir)</h3>
                  <div class="h-64">
                      <Line v-if="trendChartData" :data="trendChartData" :options="chartOptions as any" />
                  </div>
              </div>
          </div>

          <!-- Lists Section -->
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
              <!-- Popular Events -->
              <div class="bg-white p-6 rounded-xl shadow-sm border border-gray-100">
                  <h3 class="font-bold text-gray-800 mb-4">Kajian Teramai</h3>
                  <div class="space-y-4">
                      <div v-for="(event, idx) in dashboardStats.popular_events" :key="idx" class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
                          <span class="font-semibold text-gray-700">{{ idx + 1 }}. {{ event.nama_kegiatan }}</span>
                          <span class="px-3 py-1 bg-green-100 text-green-700 rounded-full text-xs font-bold">{{ event.total_hadir }} Hadir</span>
                      </div>
                      <p v-if="!dashboardStats.popular_events?.length" class="text-center text-gray-400 text-sm">Belum ada data</p>
                  </div>
              </div>

              <!-- Top Participants -->
              <div class="bg-white p-6 rounded-xl shadow-sm border border-gray-100">
                  <h3 class="font-bold text-gray-800 mb-4">Peserta Paling Rajin</h3>
                   <div class="space-y-4">
                      <div v-for="(p, idx) in dashboardStats.top_peserta" :key="idx" class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
                          <span class="font-semibold text-gray-700">{{ idx + 1 }}. {{ p.nama }}</span>
                          <span class="px-3 py-1 bg-blue-100 text-blue-700 rounded-full text-xs font-bold">{{ p.total_hadir }}x Hadir</span>
                      </div>
                      <p v-if="!dashboardStats.top_peserta?.length" class="text-center text-gray-400 text-sm">Belum ada data</p>
                  </div>
              </div>
          </div>

        </div>

        <!-- PESERTA VIEW -->
        <div v-else-if="currentView === 'peserta'" class="animate-fade">
          <div class="flex justify-between items-center mb-6">
             <h1 class="text-3xl font-bold text-gray-800">Data Peserta</h1>
             <button @click="exportCSV" class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700">Export CSV</button>
          </div>
          <div class="bg-white rounded-xl shadow-sm border border-gray-100 overflow-hidden">
             <div class="p-4 border-b border-gray-100">
                <input v-model="searchPeserta" @input="onSearchPeserta" type="text" placeholder="Cari Nama / No HP..." class="w-full md:w-80 px-4 py-2 border rounded-lg focus:ring-2 focus:ring-green-500 outline-none">
             </div>
             <div class="overflow-x-auto">
               <table class="w-full text-left whitespace-nowrap">
                  <thead class="bg-gray-50 text-gray-500 text-sm uppercase">
                    <tr>
                      <th class="px-6 py-4">Nama Lengkap</th>
                      <th class="px-6 py-4">Panggilan</th>
                      <th class="px-6 py-4">JK</th>
                      <th class="px-6 py-4">Umur</th>
                      <th class="px-6 py-4">Domisili</th>
                      <th class="px-6 py-4">No HP</th>
                      <th class="px-6 py-4">Aksi</th>
                    </tr>
                  </thead>
                  <tbody class="divide-y divide-gray-100 text-sm">
                    <tr v-if="pesertaList.length === 0">
                       <td colspan="7" class="px-6 py-8 text-center text-gray-400">Tidak ada data ditemukan.</td>
                    </tr>
                    <tr v-for="p in pesertaList" :key="p.id" class="hover:bg-gray-50">
                      <td class="px-6 py-4 font-bold text-gray-700">{{ p.nama_lengkap }}</td>
                      <td class="px-6 py-4">{{ p.panggilan }}</td>
                      <td class="px-6 py-4">{{ p.jenis_kelamin }}</td>
                      <td class="px-6 py-4">
                        <span v-if="p.tahun_lahir">{{ new Date().getFullYear() - p.tahun_lahir }} Thn</span>
                        <span v-else class="text-gray-400">-</span>
                      </td>
                      <td class="px-6 py-4 text-gray-500">{{ p.alamat_domisili }}</td>
                      <td class="px-6 py-4 text-gray-500">{{ p.no_hp }}</td>
                      <td class="px-6 py-4 flex gap-2">
                          <button @click="openEditPeserta(p)" class="text-blue-600 hover:text-blue-800 font-bold">Edit</button>
                          <button @click="deletePeserta(p.id)" class="text-red-600 hover:text-red-800 font-bold">Hapus</button>
                      </td>
                    </tr>
                  </tbody>
               </table>
             </div>
             <div class="p-4 border-t border-gray-100 text-center text-gray-400 text-sm">
                Showing {{ pesertaList.length }} participants
             </div>
          </div>

          <!-- Edit Participant Modal -->
           <div v-if="showEditPesertaModal" class="fixed inset-0 z-50 flex items-center justify-center bg-black/50 p-4">
              <div class="bg-white rounded-2xl w-full max-w-lg p-6 shadow-xl animate-scale-in max-h-[90vh] overflow-y-auto">
                 <h3 class="text-xl font-bold mb-4">Edit Data Peserta</h3>
                 <form @submit.prevent="updatePeserta" class="space-y-4">
                    <div class="grid grid-cols-2 gap-4">
                        <div>
                           <label class="block text-sm font-bold text-gray-700 mb-1">Nama Lengkap</label>
                           <input v-model="editPesertaForm.nama_lengkap" required class="w-full p-2 border rounded-lg focus:ring-2 focus:ring-green-500">
                        </div>
                        <div>
                           <label class="block text-sm font-bold text-gray-700 mb-1">Panggilan</label>
                           <input v-model="editPesertaForm.panggilan" class="w-full p-2 border rounded-lg focus:ring-2 focus:ring-green-500">
                        </div>
                    </div>
                    
                    <div class="grid grid-cols-2 gap-4">
                        <div>
                           <label class="block text-sm font-bold text-gray-700 mb-1">Jenis Kelamin</label>
                           <select v-model="editPesertaForm.jenis_kelamin" class="w-full p-2 border rounded-lg focus:ring-2 focus:ring-green-500">
                               <option value="Laki-laki">Laki-laki</option>
                               <option value="Perempuan">Perempuan</option>
                           </select>
                        </div>
                        <div>
                           <label class="block text-sm font-bold text-gray-700 mb-1">Tahun Lahir</label>
                           <input v-model="editPesertaForm.tahun_lahir" type="number" class="w-full p-2 border rounded-lg focus:ring-2 focus:ring-green-500">
                        </div>
                    </div>

                    <div>
                       <label class="block text-sm font-bold text-gray-700 mb-1">Domisili</label>
                       <textarea v-model="editPesertaForm.alamat_domisili" rows="2" class="w-full p-2 border rounded-lg focus:ring-2 focus:ring-green-500"></textarea>
                    </div>

                    <div class="grid grid-cols-2 gap-4">
                        <div>
                           <label class="block text-sm font-bold text-gray-700 mb-1">No HP</label>
                           <input v-model="editPesertaForm.no_hp" class="w-full p-2 border rounded-lg focus:ring-2 focus:ring-green-500">
                        </div>
                         <div>
                           <label class="block text-sm font-bold text-gray-700 mb-1">Email</label>
                           <input v-model="editPesertaForm.email" type="email" class="w-full p-2 border rounded-lg focus:ring-2 focus:ring-green-500">
                        </div>
                    </div>

                    <div class="flex justify-end gap-3 mt-6">
                       <button type="button" @click="closeEditPesertaModal" class="px-4 py-2 text-gray-600 hover:bg-gray-100 rounded-lg">Batal</button>
                       <button type="submit" class="px-4 py-2 bg-green-600 text-white font-bold rounded-lg hover:bg-green-700">Simpan Perubahan</button>
                    </div>
                 </form>
              </div>
           </div>
        </div>

        <!-- KEGIATAN VIEW -->
        <div v-else-if="currentView === 'kegiatan'" class="animate-fade">
           <div class="flex justify-between items-center mb-6">
              <h1 class="text-3xl font-bold text-gray-800">Manajemen Kegiatan</h1>
              <button @click="showAddKegiatanModal = true" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 font-bold">
                 + Tambah Kegiatan
              </button>
           </div>
           
           <div class="bg-white rounded-xl shadow-sm border border-gray-100 overflow-hidden">
              <table class="w-full text-left">
                 <thead class="bg-gray-50 text-gray-500 text-sm uppercase">
                    <tr>
                       <th class="px-6 py-4">Nama Kegiatan</th>
                       <th class="px-6 py-4">Waktu</th>
                       <th class="px-6 py-4">Pemateri</th>
                       <th class="px-6 py-4">Tempat</th>
                       <th class="px-6 py-4">Aksi</th>
                    </tr>
                 </thead>
                 <tbody class="divide-y divide-gray-100 text-sm">
                    <tr v-if="kegiatanList.length === 0">
                       <td colspan="5" class="px-6 py-8 text-center text-gray-400">Belum ada data kegiatan.</td>
                    </tr>
                    <tr v-for="item in kegiatanList" :key="item.id" class="hover:bg-gray-50">
                       <td class="px-6 py-4 font-bold">{{ item.nama_kegiatan }}</td>
                       <td class="px-6 py-4">{{ new Date(item.waktu).toLocaleString('id-ID') }}</td>
                       <td class="px-6 py-4">{{ item.pemateri }}</td>
                       <td class="px-6 py-4">{{ item.tempat }}</td>
                       <td class="px-6 py-4 flex gap-2">
                          <button @click="editKegiatan(item)" class="text-blue-600 hover:text-blue-800 font-bold">Edit</button>
                          <button @click="deleteKegiatan(item.id)" class="text-red-600 hover:text-red-800 font-bold">Hapus</button>
                       </td>
                    </tr>
                 </tbody>
              </table>
           </div>

           <!-- Modal Form -->
           <div v-if="showAddKegiatanModal" class="fixed inset-0 z-50 flex items-center justify-center bg-black/50 p-4">
              <div class="bg-white rounded-2xl w-full max-w-md p-6 shadow-xl animate-scale-in">
                 <h3 class="text-xl font-bold mb-4">{{ isEditing ? 'Edit Kegiatan' : 'Tambah Kegiatan Baru' }}</h3>
                 <form @submit.prevent="submitKegiatan" class="space-y-4">
                    <div>
                       <label class="block text-sm font-bold text-gray-700 mb-1">Nama Kegiatan</label>
                       <input v-model="kegiatanForm.nama_kegiatan" required class="w-full p-2 border rounded-lg focus:ring-2 focus:ring-green-500">
                    </div>
                    <div>
                       <label class="block text-sm font-bold text-gray-700 mb-1">Waktu</label>
                       <input v-model="kegiatanForm.waktu" type="datetime-local" required class="w-full p-2 border rounded-lg focus:ring-2 focus:ring-green-500">
                    </div>
                    <div>
                       <label class="block text-sm font-bold text-gray-700 mb-1">Pemateri</label>
                       <input v-model="kegiatanForm.pemateri" class="w-full p-2 border rounded-lg focus:ring-2 focus:ring-green-500">
                    </div>
                     <div>
                       <label class="block text-sm font-bold text-gray-700 mb-1">Tempat</label>
                       <input v-model="kegiatanForm.tempat" class="w-full p-2 border rounded-lg focus:ring-2 focus:ring-green-500">
                    </div>
                    <div class="flex justify-end gap-3 mt-6">
                       <button type="button" @click="closeModal" class="px-4 py-2 text-gray-600 hover:bg-gray-100 rounded-lg">Batal</button>
                       <button type="submit" class="px-4 py-2 bg-green-600 text-white font-bold rounded-lg hover:bg-green-700">Simpan</button>
                    </div>
                 </form>
              </div>
           </div>
        </div>

        <!-- SETTINGS VIEW -->
        <div v-else-if="currentView === 'settings'" class="animate-fade">
           <h1 class="text-3xl font-bold text-gray-800 mb-8">Pengaturan Aplikasi</h1>
           <div class="space-y-6 max-w-2xl">
              <div class="bg-white p-6 rounded-xl shadow-sm border border-gray-100">
                 <h2 class="text-xl font-bold mb-4">Event Aktif</h2>
                 <div class="space-y-4">
                    <div>
                       <label class="block text-sm font-medium text-gray-700 mb-1">Pilih Kegiatan untuk Check-in</label>
                       <select v-model="settingsForm.activeEventId" class="w-full p-3 border rounded-lg bg-gray-50">
                          <option value="" disabled>-- Pilih Event --</option>
                          <option v-for="bg in kegiatanList" :key="bg.id" :value="bg.id">
                             {{ bg.nama_kegiatan }} ({{ new Date(bg.waktu).toLocaleString() }})
                          </option>
                       </select>
                    </div>
                    <button @click="saveActiveEvent" class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">Simpan Perubahan</button>
                 </div>
              </div>

               <div class="bg-white p-6 rounded-xl shadow-sm border border-gray-100">
                 <h2 class="text-xl font-bold mb-4">Aset Gambar</h2>
                 <div class="space-y-4">
                    <div>
                        <span class="block text-sm font-medium text-gray-700 mb-1">Gambar QRIS</span>
                        <div class="flex items-center gap-4">
                            <div class="w-24 h-24 bg-gray-200 rounded-lg overflow-hidden flex items-center justify-center border">
                                <img v-if="settingsForm.qrisUrl" :src="settingsForm.qrisUrl" class="w-full h-full object-cover">
                                <span v-else class="text-xs text-gray-500">No Image</span>
                            </div>
                            <label class="cursor-pointer text-blue-600 font-bold hover:underline">
                                Upload Baru
                                <input type="file" @change="uploadQRIS" class="hidden" accept="image/*">
                            </label>
                        </div>
                    </div>

                     <!-- Background Gallery -->
                     <div>
                        <h3 class="block text-sm font-medium text-gray-700 mb-2">Background Gallery (Carousel)</h3>
                        
                        <!-- Gallery Grid -->
                        <div class="grid grid-cols-2 gap-4 mb-4">
                            <div v-for="bg in backgrounds" :key="bg.id" class="group relative aspect-video rounded-lg overflow-hidden border border-gray-200 bg-gray-100">
                                <img :src="`${api.defaults.baseURL}/backgrounds/${bg.id}`" class="w-full h-full object-cover">
                                
                                <!-- Status Badge -->
                                <div class="absolute top-2 left-2 z-10">
                                    <span v-if="bg.is_active" class="px-2 py-1 bg-green-500 text-white text-[10px] font-bold rounded-full shadow-sm">Active</span>
                                    <span v-else class="px-2 py-1 bg-gray-500 text-white text-[10px] font-bold rounded-full shadow-sm">Inactive</span>
                                </div>

                                <!-- Actions Overlay -->
                                <div class="absolute inset-0 bg-black/50 opacity-0 group-hover:opacity-100 transition flex items-center justify-center gap-2">
                                    <button @click="toggleBackground(bg)" class="p-2 bg-white rounded-full hover:bg-gray-100 text-gray-700 shadow-sm transition transform hover:scale-105" title="Toggle Active">
                                        <span v-if="bg.is_active">🚫</span>
                                        <span v-else>✅</span>
                                    </button>
                                    <button @click="deleteBackground(bg.id)" class="p-2 bg-red-500 rounded-full hover:bg-red-600 text-white shadow-sm transition transform hover:scale-105" title="Delete">
                                        🗑️
                                    </button>
                                </div>
                            </div>

                            <!-- Upload Card -->
                            <label class="border-2 border-dashed border-gray-300 rounded-lg flex flex-col items-center justify-center cursor-pointer hover:border-green-500 hover:bg-green-50 transition aspect-video group">
                                <span class="text-3xl mb-1 group-hover:scale-110 transition">➕</span>
                                <span class="text-xs text-gray-500 font-bold">Add Image</span>
                                <input type="file" class="hidden" accept="image/*" @change="uploadBackgroundGallery">
                            </label>
                        </div>
                        <p class="text-xs text-gray-500">
                           <span class="font-bold">Tips:</span> Gambar dengan status "Active" akan berganti secara otomatis (slide carousel) di halaman pendaftaran.
                        </p>
                     </div>
                 </div>
              </div>
           </div>
        </div>

      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, onUnmounted, watch } from 'vue'
import api from '../services/api'
import {
  Chart as ChartJS,
  Title,
  Tooltip,
  Legend,
  BarElement,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  type ChartData,
  type ChartOptions
} from 'chart.js'
import { Bar, Line } from 'vue-chartjs'

ChartJS.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend, PointElement, LineElement)

interface DashboardStats {
  total_peserta: number
  total_peserta_new: number
  total_checkin_today: number
  gender_stats: { laki_laki: number; perempuan: number }
  age_stats: { remaja: number; dewasa: number; lansia: number }
  registration_trend: { date: string; count: number }[]
  active_event?: { nama_kegiatan: string; waktu: string; pemateri: string }
  popular_events: any[]
  top_peserta: any[]
}

interface BackgroundImage {
    id: string
    is_active: boolean
    created_at?: string
}

interface Peserta {
    id: string
    nama_lengkap: string
    panggilan: string
    jenis_kelamin: string
    alamat_domisili: string
    no_hp: string
    email: string
    tahun_lahir: number
}

interface Kegiatan {
    id: string
    nama_kegiatan: string
    waktu: string
    pemateri: string
    tempat: string
}

const currentView = ref('dashboard')
const isSidebarOpen = ref(false)

const dashboardStats = ref<Partial<DashboardStats>>({})
const ageChartData = ref<ChartData<'bar'> | null>(null)
const trendChartData = ref<ChartData<'line'> | null>(null)
const backgrounds = ref<BackgroundImage[]>([])

const chartOptions: ChartOptions = {
    responsive: true,
    maintainAspectRatio: false
}

const fetchStats = async () => {
    try {
        const res = await api.get('/admin/stats')
        const data = res.data.data
        dashboardStats.value = data

        // ... existing chart logic ... (omitted for brevity in replacement if possible, but replace_file_content needs contigous block. 
        // I will just keep fetchStats as is mostly, but need to add fetchBackgrounds call or separate it)
        
        // Age Chart Logic...
        ageChartData.value = {
            labels: ['Remaja (<20)', 'Dewasa (20-50)', 'Lansia (>50)'],
            datasets: [
                {
                    label: 'Jumlah Peserta',
                    backgroundColor: ['#60A5FA', '#34D399', '#F472B6'],
                    data: [data.age_stats.remaja, data.age_stats.dewasa, data.age_stats.lansia]
                }
            ]
        }
        // Trend Chart Logic...
         if (data.registration_trend) {
             trendChartData.value = {
                labels: data.registration_trend.map((d: any) => d.date),
                datasets: [
                    {
                        label: 'Pendaftar Baru',
                        backgroundColor: '#10B981',
                        borderColor: '#10B981',
                        data: data.registration_trend.map((d: any) => d.count),
                        tension: 0.3
                    }
                ]
            }
        }

    } catch (e) {
        console.error("Failed to fetch stats", e)
    }
}

// Background Gallery Logic
const fetchBackgrounds = async () => {
    try {
        const res = await api.get('/admin/backgrounds')
        backgrounds.value = res.data.data
    } catch (e) { console.error(e) }
}

const uploadBackgroundGallery = async (e: Event) => {
    const input = e.target as HTMLInputElement
    const file = input.files?.[0]
    if (!file) return
    
    const formData = new FormData()
    formData.append('image', file)
    
    try {
        await api.post('/admin/backgrounds', formData)
        fetchBackgrounds()
        input.value = '' // reset input
    } catch (err) {
        alert("Upload failed")
    }
}

const toggleBackground = async (bg: BackgroundImage) => {
    try {
        await api.post(`/admin/backgrounds/${bg.id}/toggle`)
        fetchBackgrounds()
    } catch (e) { alert("Failed to toggle") }
}

const deleteBackground = async (id: string) => {
    if(!confirm("Hapus background ini?")) return
    try {
        await api.delete(`/admin/backgrounds/${id}`)
        fetchBackgrounds()
    } catch (e) { alert("Failed to delete") }
}

let pollingInterval: ReturnType<typeof setInterval> | null = null

onMounted(() => {
    fetchStats() 
    fetchBackgrounds() // Fetch gallery
    pollingInterval = setInterval(fetchStats, 2000)
})

onUnmounted(() => {
    if (pollingInterval) clearInterval(pollingInterval)
})

// Kegiatan Logic
const kegiatanList = ref<Kegiatan[]>([])
const showAddKegiatanModal = ref(false)
const isEditing = ref(false)
const editingId = ref<string | null>(null)

const kegiatanForm = reactive({
  nama_kegiatan: '',
  waktu: '',
  pemateri: '',
  tempat: 'Masjid Jogokariyan'
})

// Settings Logic
const settingsForm = reactive({
    activeEventId: '',
    qrisUrl: null as string | null
})

const fetchSettings = async () => {
    try {
        const [resParams, resEvents] = await Promise.all([
            api.get('/admin/settings'),
            api.get('/admin/kegiatan')
        ])
        
        kegiatanList.value = resEvents.data.data
        
        const activeEventSetting = resParams.data.data.find((s: any) => s.key === 'active_event_id')
        if (activeEventSetting) {
            settingsForm.activeEventId = activeEventSetting.value
        }
        
        const qrisSetting = resParams.data.data.find((s: any) => s.key === 'qris_image_path')
        if (qrisSetting) {
            // Serve from new endpoint with timestamp
            settingsForm.qrisUrl = `${api.defaults.baseURL}/images/qris?t=${Date.now()}`
        }

        if (qrisSetting) {
            // Serve from new endpoint with timestamp
            settingsForm.qrisUrl = `${api.defaults.baseURL}/images/qris?t=${Date.now()}`
        }
    } catch (err) {
        console.error("Failed settings load", err)
    }
}

const saveActiveEvent = async () => {
    try {
        await api.post('/admin/settings/active-event', { event_id: settingsForm.activeEventId })
        alert("Event aktif berhasil diperbarui!")
    } catch (err) {
        alert("Gagal update event aktif")
    }
}

const uploadQRIS = async (event: Event) => {
    const input = event.target as HTMLInputElement
    const file = input.files?.[0]
    if (!file) return
    
    const formData = new FormData()
    formData.append("qris_image", file)
    
    try {
        await api.post('/admin/settings/qris', formData, {
            headers: { 'Content-Type': 'multipart/form-data' }
        })
        // Refresh view
        settingsForm.qrisUrl = `${api.defaults.baseURL}/images/qris?t=${Date.now()}`
        alert("QRIS berhasil diupload ke Database")
    } catch (err: any) {
        console.error(err)
        alert("Gagal upload QRIS: " + (err.response?.data?.message || err.message))
    }
}




// Data Peserta Logic
const pesertaList = ref<Peserta[]>([])
const searchPeserta = ref('')
let searchTimeout: ReturnType<typeof setTimeout> | null = null

const fetchPeserta = async () => {
    try {
        const res = await api.get('/admin/peserta', {
            params: { search: searchPeserta.value }
        })
        pesertaList.value = res.data.data
    } catch (err) {
        console.error("Failed fetch peserta", err)
    }
}

const onSearchPeserta = () => {
    if (searchTimeout) clearTimeout(searchTimeout)
    searchTimeout = setTimeout(() => {
        fetchPeserta()
    }, 500)
}

// Edit & Delete Peserta Logic
const showEditPesertaModal = ref(false)
const editPesertaId = ref<string | null>(null)
const editPesertaForm = reactive({
    nama_lengkap: '',
    panggilan: '',
    jenis_kelamin: '',
    alamat_domisili: '',
    no_hp: '',
    email: '',
    tahun_lahir: ''
})

const openEditPeserta = (p: Peserta) => {
    editPesertaId.value = p.id
    editPesertaForm.nama_lengkap = p.nama_lengkap
    editPesertaForm.panggilan = p.panggilan
    editPesertaForm.jenis_kelamin = p.jenis_kelamin
    editPesertaForm.alamat_domisili = p.alamat_domisili
    editPesertaForm.no_hp = p.no_hp
    editPesertaForm.email = p.email
    editPesertaForm.tahun_lahir = String(p.tahun_lahir)
    showEditPesertaModal.value = true
}

const closeEditPesertaModal = () => {
    showEditPesertaModal.value = false
    editPesertaId.value = null
}

const updatePeserta = async () => {
    try {
        // Convert tahun_lahir to int
        const payload = { ...editPesertaForm, tahun_lahir: parseInt(editPesertaForm.tahun_lahir) }
        await api.put(`/admin/peserta/${editPesertaId.value}`, payload)
        alert("Data peserta berhasil diupdate")
        closeEditPesertaModal()
        fetchPeserta()
    } catch (err: any) {
        console.error(err)
        alert("Gagal update peserta: " + (err.response?.data?.message || err.message))
    }
}

const deletePeserta = async (id: string) => {
    if (!confirm("Yakin ingin menghapus peserta ini? Data tidak bisa dikembalikan.")) return
    try {
        await api.delete(`/admin/peserta/${id}`)
        fetchPeserta()
    } catch (err) {
        console.error(err)
        alert("Gagal menghapus peserta")
    }
}

const exportCSV = () => {
    // Simple CSV export logic from current list
    const headers = ['Nama Lengkap', 'Panggilan', 'JK', 'Domisili', 'No HP']
    const rows = pesertaList.value.map(p => [
        `"${p.nama_lengkap}"`,
        `"${p.panggilan}"`,
        p.jenis_kelamin,
        `"${p.alamat_domisili}"`,
        `'${p.no_hp}`
    ])
    
    const csvContent = [headers.join(','), ...rows.map(r => r.join(','))].join('\n')
    const blob = new Blob([csvContent], { type: 'text/csv;charset=utf-8;' })
    const link = document.createElement("a")
    link.href = URL.createObjectURL(blob)
    link.download = "data_peserta.csv"
    link.click()
}

const setView = (view: string) => {
  currentView.value = view
  isSidebarOpen.value = false
  if (view === 'kegiatan') fetchKegiatan()
  if (view === 'settings') fetchSettings()
  if (view === 'peserta') fetchPeserta()
}

const fetchKegiatan = async () => {
    try {
        const res = await api.get('/admin/kegiatan')
        kegiatanList.value = res.data.data
    } catch (err) {
        console.error("Failed fetch kegiatan", err)
    }
}

const submitKegiatan = async () => {
    try {
        // Format date to ISO for backend if needed, but datetime-local usually works fine
        // Note: datetime-local value is "YYYY-MM-DDTHH:mm", Backend expects ISO8601
        const payload = {
            ...kegiatanForm,
            waktu: new Date(kegiatanForm.waktu).toISOString()
        }

        if (isEditing.value) {
            await api.put(`/admin/kegiatan/${editingId.value}`, payload)
        } else {
            await api.post('/admin/kegiatan', payload)
        }
        
        closeModal()
        fetchKegiatan()
    } catch (err: any) {
        console.error(err)
        alert("Gagal menyimpan data: " + (err.response?.data?.message || err.message))
    }
}

const editKegiatan = (item: Kegiatan) => {
    isEditing.value = true
    editingId.value = item.id
    kegiatanForm.nama_kegiatan = item.nama_kegiatan
    // Adjust datetime for input type="datetime-local" (YYYY-MM-DDTHH:mm)
    const dt = new Date(item.waktu)
    dt.setMinutes(dt.getMinutes() - dt.getTimezoneOffset()) // adjust for local time display
    kegiatanForm.waktu = dt.toISOString().slice(0, 16)
    
    kegiatanForm.pemateri = item.pemateri
    kegiatanForm.tempat = item.tempat
    showAddKegiatanModal.value = true
}

const deleteKegiatan = async (id: string) => {
    if(!confirm("Yakin hapus kegiatan ini?")) return
    try {
        await api.delete(`/admin/kegiatan/${id}`)
        fetchKegiatan()
    } catch (err) {
        alert("Gagal menghapus")
    }
}

const closeModal = () => {
    showAddKegiatanModal.value = false
    isEditing.value = false
    editingId.value = null
    kegiatanForm.nama_kegiatan = ''
    kegiatanForm.waktu = ''
    kegiatanForm.pemateri = ''
    kegiatanForm.tempat = 'Masjid Jogokariyan'
}
</script>

<style scoped>
.animate-fade {
  animation: fadeIn 0.3s ease-out;
}
@keyframes fadeIn {
  from { opacity: 0; transform: translateY(5px); }
  to { opacity: 1; transform: translateY(0); }
}
</style>
