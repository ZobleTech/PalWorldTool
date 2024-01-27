import {createRouter, createWebHistory} from "vue-router"

const routes = [
	{
		path: "/",
		name: "home",
		component: () => import("../components/IniGenerator.vue"),
	},

]

const router = createRouter({
	history: createWebHistory(),
	routes,
})

router.beforeEach((to, from, next) => {
	next()
})

export default router