<template>
	<div>
		<n-config-provider
				:date-locale="dateZhCN"
				:locale="zhCN"
				:theme="darkTheme"
				class="appCon"
		>
			<n-loading-bar-provider>
				<n-message-provider>
					<router-view></router-view>
				</n-message-provider>
			</n-loading-bar-provider>
		</n-config-provider>
	</div>
</template>

<script setup>
import { provide } from 'vue';
import axios from "axios";
import {useRoute, useRouter} from "vue-router";
import {data} from "./store";
import {darkTheme, dateZhCN, NConfigProvider, NLoadingBarProvider, NMessageProvider, zhCN} from "naive-ui";

const route = useRoute()
const router = useRouter()

const request = axios.create({
	baseURL: "/",
	headers: {
		Authorization: localStorage.getItem("Token"),
	},
});
const datas = data()

provide("request", request)
provide('route', route)
provide('router', router)
provide('datas',datas)
</script>

<style>
.n-modal-mask{
	background-color: rgba(0, 0, 0, 0.5);
	backdrop-filter: blur(8px) !important;
}
</style>