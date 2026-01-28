import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'
import Registration from '../pages/Registration.vue'
import CheckIn from '../pages/CheckIn.vue'
import Admin from '../pages/Admin.vue'

const routes: RouteRecordRaw[] = [
    { path: '/', component: Registration },
    { path: '/check-in', component: CheckIn },
    { path: '/admin', component: Admin },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router
