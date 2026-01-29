import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'
import Registration from '../pages/Registration.vue'
import CheckIn from '../pages/CheckIn.vue'
import Admin from '../pages/Admin.vue'

import Login from '../pages/Login.vue'

const routes: RouteRecordRaw[] = [
    { path: '/', component: Registration },
    { path: '/check-in', component: CheckIn },
    { path: '/admin', component: Admin },
    { path: '/login', component: Login },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

router.beforeEach((to, from, next) => {
    const isAuthenticated = !!localStorage.getItem('token');

    if (to.path.startsWith('/admin') && !isAuthenticated) {
        next('/login');
    } else if (to.path === '/login' && isAuthenticated) {
        next('/admin');
    } else {
        next();
    }
});

export default router
