<template>
	<div class="iniGenCon">
		<div class="title" @click="openBili()">
		</div>
		<div class="iniContainer">
			<div class="iniItemOp uploadIni">
				<div class="left">修改现有的配置文件</div>
				<div class="right">
					<div @click="analyzeData.showPanel = true" class="analyze uploadFunc">
						分析文本
					</div>
					<n-upload ref="upload"
							  :accept="'.ini'"
							  :default-upload="false"
							  :show-file-list="false"
							  action="https://www.mocky.io/v2/5e4bafc63100007100d8b70f"
							  class="uploadFunc"
							  @change="handleUpload"
					>
						<div class="uploadBtn">选择文件</div>
					</n-upload>
				</div>
			</div>
			<n-scrollbar class="iniScroll" style="max-height: 60vh" trigger="none">
				<div class="iniItems">

					<div v-for="(item,index) in IniData" class="iniItem">
						<div v-if="item.toggleType==='slide'" class="iniItemOp inislider">
							<div class="left">{{ item.labelCN }}</div>
							<div class="center">{{ item.defaultValue }}</div>
							<n-slider v-model:value="IniData[index].defaultValue" :max="item.maxValue"
									  :min="item.minValue"
									  :step="getStepValue(item.minValue)" :tooltip="false" class="right"/>
						</div>
						<div v-else-if="item.toggleType==='toggle'" class="iniItemOp initoggle">
							<div class="left">{{ item.labelCN }}</div>
							<div class="center"></div>
							<div class="right">
								<div :class="{toggleActive:item.defaultValue === false}" class="btn"
									 @click="changeToggle(index,false)">OFF
								</div>
								<div :class="{toggleActive:item.defaultValue === true}" class="btn"
									 @click="changeToggle(index,true)">ON
								</div>

							</div>
						</div>
						<div v-else-if="item.toggleType==='input'" class="iniItemOp iniinput">
							<div class="left">{{ item.labelCN }}</div>
							<div class="center"></div>
							<input v-model="item.defaultValue" class="right" type="text">
						</div>
						<div v-else class="iniItemOp iniselect">
							<div class="left">{{ item.labelCN }}</div>
							<div class="center"></div>
							<div class="right">
								<div class="toLeft" @click="changeSelectAddorDec(index,-1)">
									<n-icon :size="18">
										<svg viewBox="0 0 24 24"
											 xmlns="http://www.w3.org/2000/svg"
											 xmlns:xlink="http://www.w3.org/1999/xlink">
											<path d="M17.77 3.77L16 2L6 12l10 10l1.77-1.77L9.54 12z"
												  fill="currentColor"></path>
										</svg>
									</n-icon>
								</div>
								<div class="middleInfo">
									<div class="mTop">{{ item.options[item.defaultValue].CN }}</div>
									<div class="mBottom">
										<div v-for="(litem,lindex) in item.options"
											 :class="{bLineActive:item.defaultValue==lindex}"
											 class="bLine"
											 @click="changeSelect(index,lindex)">
										</div>
									</div>
								</div>
								<div class="toRight" @click="changeSelectAddorDec(index,1)">
									<n-icon :size="18">
										<svg viewBox="0 0 24 24"
											 xmlns="http://www.w3.org/2000/svg"
											 xmlns:xlink="http://www.w3.org/1999/xlink">
											<path d="M17.77 3.77L16 2L6 12l10 10l1.77-1.77L9.54 12z"
												  fill="currentColor"></path>
										</svg>
									</n-icon>
								</div>
							</div>
						</div>
					</div>
				</div>
			</n-scrollbar>
		</div>
		<div class="ending">
			<div class="submit btn" @click="RandomIni()">随 机</div>
			<div class="preview btn" @click="previewIniFile()">预 览</div>
			<div class="submit btn" @click="downloadIniFile()">保 存</div>
		</div>
		<div class="background">
			<n-carousel :autoplay="true" :interval="10000" :loop="true" :show-dots="false" effect="fade">
				<div v-for="(item,index) in bgImgs" :style="{backgroundImage:`url(${item})`}"
					 class="bg"></div>
			</n-carousel>
		</div>
		<n-drawer v-model:show="iniPreview" :default-width="'35%'" :resizable="true"
				  style="background: rgba(10,10,10,0.8)">
			<n-drawer-content :native-scrollbar="false" closable title="PalWorldSettings.ini">
				<div style="white-space: pre-line;font-size: 18px;line-height: 1.2">
					{{ iniStrings }}
				</div>

			</n-drawer-content>
		</n-drawer>
		<n-modal v-model:show="analyzeData.showPanel">
			<div class="analyzeCard">
				<div class="analyzeContainer">
					<n-input
							v-model:value="analyzeData.data"
							type="textarea"
							placeholder="请粘贴配置文件内容至此处"
							:autosize="{ minRows: 12, maxRows: 24 }"
							class="inputarea"
					/>
					<div class="btn" @click="analyzeIni()">开始分析</div>
				</div>
			</div>
		</n-modal>
	</div>
</template>

<script setup>
import {NCarousel, NDrawer, NDrawerContent, NIcon, NScrollbar, NSlider, NUpload, NModal,NInput} from 'naive-ui'

import {inject, onMounted, reactive, ref, watch} from "vue";

const request = inject("request")
let iniPreview = ref(false)
const analyzeData = reactive({
	showPanel: false,
	data: ""
})

function getStepValue(baseValue) {
	if (baseValue % 1 === 0) {
		return 1
	} else {
		return 0.1
	}
}

const bgImgs = [
	"https://pic.imgdb.cn/item/65b2ac3c871b83018a41c3ef.jpg",
	"https://pic.imgdb.cn/item/65b2ac3a871b83018a41c036.jpg",
	"https://pic.imgdb.cn/item/65b2ac3b871b83018a41c1f5.png",
	"https://pic.imgdb.cn/item/65b2ac3c871b83018a41c2db.png",
	"https://pic.imgdb.cn/item/65b2ac3c871b83018a41c3bc.png",
	"https://pic.imgdb.cn/item/65b2ac48871b83018a41d9fc.png",
]

const iniStrings = ref('')

function openBili() {
	window.open('https://space.bilibili.com/7277347', '_blank')
}

function changeToggle(iniIndex, status) {
	IniData[iniIndex].defaultValue = status
}

function changeSelect(iniIndex, status) {
	IniData[iniIndex].defaultValue = status
}

function changeSelectAddorDec(iniIndex, n) {
	if ((IniData[iniIndex].defaultValue >= IniData[iniIndex].maxValue && n === 1) || (IniData[iniIndex].defaultValue <= IniData[iniIndex].minValue && n === -1)) {
		return
	}
	IniData[iniIndex].defaultValue += n
}

function formatString(n, is_float = false) {
	if (typeof n === 'boolean') {
		if (n) {
			return 'True'
		} else {
			return 'False'
		}
	} else if (typeof n === 'number') {
		if (!is_float) {
			return n
		} else {
			return n.toFixed(6)
		}
	} else {
		return n
	}
}


function analyzeIni(){

	let inistring = analyzeData.data.match(/\(([^)]+)\)/)[1].split(',')
	let iniMap = {}
	for (let i of inistring) {
		let tempIni = i.split("=")
		iniMap[tempIni[0]] = reformatString(tempIni[1])
	}
	for (let i in IniData){
		if (IniData[i].labelEN in iniMap) {
			IniData[i].defaultValue = iniMap[IniData[i].labelEN]
		}
	}
	analyzeData.data = ""
	analyzeData.showPanel = false
}
async function getIniStrings() {
	let res = {}
	for (let i of IniData) {
		if (i.toggleType === "select") {
			res[i.labelEN] = i.options[i.defaultValue].EN
			continue
		}
		if ('minValue' in i) {
			res[i.labelEN] = formatString(i.defaultValue, getStepValue(i.minValue) === 0.1)
		} else {
			res[i.labelEN] = formatString(i.defaultValue)
		}
		if (i.toggleType === 'input') {
			res[i.labelEN] = '"'+i.defaultValue+'"'
		}
	}
	await request.post("api/basic/create_ini", {"inidata": res}).then(res => {
		if (res.data.code === 100) {
			iniStrings.value = res.data.data
		}
	})
}

const iniOptions = {
	"None": 0,
	"Item": 1,
	"ItemAndEquipment": 2,
	"All": 3
}

function reformatString(n) {
	if (!isNaN(n) && n !== '') {
		if (n.includes('.')) {
			return parseFloat(n)
		} else {
			return parseInt(n)
		}
	}
	if (n === 'True') {
		return true
	} else if (n === 'False') {
		return false
	}
	if (n in iniOptions) {
		return iniOptions[n]
	}
	return n.replace(/"/g, '')
}

function handleUpload(e) {
	console.log(e)
	const iniForm = new FormData
	iniForm.append('inifile', e.file.file)
	request.post("api/basic/analyze_ini", iniForm).then(res => {
		if (res.data.code === 100) {
			for (let i in IniData) {
				console.log(i)
				IniData[i].defaultValue = reformatString(res.data.data[IniData[i].labelEN])
			}
		}
	})
}

async function previewIniFile() {
	if (iniStrings.value === '') {
		await getIniStrings()
	}
	iniPreview.value = true
}

async function downloadIniFile() {
	if (iniStrings.value === '') {
		await getIniStrings()
	}
	let iniBlob = new Blob([iniStrings.value], {type: 'application/octet-stream'})
	let iniUrl = URL.createObjectURL(iniBlob)
	let iniA = document.createElement('a')
	iniA.href = iniUrl
	iniA.download = 'PalWorldSettings.ini'
	iniA.click()
	URL.revokeObjectURL(iniUrl)
	iniBlob = null
	iniUrl = null
	iniA = null
}

function RandomIni() {
	for (let i in IniData) {
		if (typeof IniData[i].defaultValue === 'boolean') {
			IniData[i].defaultValue = generateRandomBoolean()
			continue
		}
		if (typeof IniData[i].defaultValue === 'number' && IniData[i].toggleType !== 'input') {
			IniData[i].defaultValue = generateRandomNumber(IniData[i].minValue, IniData[i].maxValue, getStepValue(IniData[i].minValue) === 0.1)
		}
	}
}

function generateRandomNumber(min, max, isFloat = false) {
	const randomDecimal = Math.random() * (max - min) + min;

	if (isFloat) {
		// 返回浮点数，并保留指定小数位数
		return parseFloat(randomDecimal.toFixed(1));
	} else {
		// 返回整数
		return Math.floor(randomDecimal);
	}
}

function generateRandomBoolean() {
	return Math.random() >= 0.5;
}

const presentIni = reactive({
	easy:{
		DayTimeSpeedRate:1.000000,
		NightTimeSpeedRate:1.000000,
		ExpRate:1.300000,
	}
})

const IniData = reactive([
	{
		minValue: 0.100000,
		maxValue: 5.000000,
		defaultValue: 1.000000,
		labelCN: "白天流逝速度",
		labelEN: "DayTimeSpeedRate",
		toggleType: "slide"
	},
	{
		minValue: 0.100000,
		maxValue: 5.000000,
		defaultValue: 1.000000,
		labelCN: "夜晚流逝速度",
		labelEN: "NightTimeSpeedRate",
		toggleType: "slide"
	},
	{
		minValue: 0.100000,
		maxValue: 20.000000,
		defaultValue: 1.000000,
		labelCN: "经验值倍率",
		labelEN: "ExpRate",
		toggleType: "slide"
	},
	{
		minValue: 0.200000,
		maxValue: 2.000000,
		defaultValue: 1.000000,
		labelCN: "捕获概率倍率",
		labelEN: "PalCaptureRate",
		toggleType: "slide"
	},
	{
		minValue: 0.100000,
		maxValue: 3.000000,
		defaultValue: 1.000000,
		labelCN: "帕鲁出现数量倍率",
		labelEN: "PalSpawnNumRate",
		toggleType: "slide"
	},
	{
		minValue: 0.100000,
		maxValue: 5.000000,
		defaultValue: 1.000000,
		labelCN: "帕鲁攻击伤害倍率",
		labelEN: "PalDamageRateAttack",
		toggleType: "slide"
	},
	{
		minValue: 0.100000,
		maxValue: 5.000000,
		defaultValue: 1.000000,
		labelCN: "帕鲁承受伤害倍率",
		labelEN: "PalDamageRateDefense",
		toggleType: "slide"
	},
	{
		minValue: 0.100000,
		maxValue: 5.000000,
		defaultValue: 1.000000,
		labelCN: "玩家攻击伤害倍率",
		labelEN: "PlayerDamageRateAttack",
		toggleType: "slide"
	},
	{
		minValue: 0.100000,
		maxValue: 5.000000,
		defaultValue: 1.000000,
		labelCN: "玩家承受伤害倍率",
		labelEN: "PlayerDamageRateDefense",
		toggleType: "slide"
	},
	{
		minValue: 0.100000,
		maxValue: 5.000000,
		defaultValue: 1.000000,
		labelCN: "玩家饱食度降低倍率",
		labelEN: "PlayerStomachDecreaceRate",
		toggleType: "slide"
	},
	{
		minValue: 0.100000,
		maxValue: 5.000000,
		defaultValue: 1.000000,
		labelCN: "玩家耐力降低倍率",
		labelEN: "PlayerStaminaDecreaceRate",
		toggleType: "slide"
	},
	{
		minValue: 0.100000,
		maxValue: 5.000000,
		defaultValue: 1.000000,
		labelCN: "玩家生命值自然回复倍率",
		labelEN: "PlayerAutoHPRegeneRate",
		toggleType: "slide"
	},
	{
		minValue: 0.100000,
		maxValue: 5.000000,
		defaultValue: 1.000000,
		labelCN: "玩家睡眠时生命值回复倍率",
		labelEN: "PlayerAutoHpRegeneRateInSleep",
		toggleType: "slide"
	},
	{
		minValue: 0.100000,
		maxValue: 5.000000,
		defaultValue: 1.000000,
		labelCN: "帕鲁饱食度降低倍率",
		labelEN: "PalStomachDecreaceRate",
		toggleType: "slide"
	},
	{
		minValue: 0.100000,
		maxValue: 5.000000,
		defaultValue: 1.000000,
		labelCN: "帕鲁耐力降低倍率",
		labelEN: "PalStaminaDecreaceRate",
		toggleType: "slide"
	},
	{
		minValue: 0.100000,
		maxValue: 5.000000,
		defaultValue: 1.000000,
		labelCN: "帕鲁生命值自然回复倍率",
		labelEN: "PalAutoHPRegeneRate",
		toggleType: "slide"
	},
	{
		minValue: 0.100000,
		maxValue: 5.000000,
		defaultValue: 1.000000,
		labelCN: "帕鲁终端中的生命值回复倍率",
		labelEN: "PalAutoHpRegeneRateInSleep",
		toggleType: "slide"
	},
	{
		minValue: 0.500000,
		maxValue: 3.000000,
		defaultValue: 1.000000,
		labelCN: "对建筑伤害倍率",
		labelEN: "BuildObjectDamageRate",
		toggleType: "slide"
	},
	{
		minValue: 0.000000,
		maxValue: 10.000000,
		defaultValue: 1.000000,
		labelCN: "建筑物的劣化速度倍率",
		labelEN: "BuildObjectDeteriorationDamageRate",
		toggleType: "slide"
	},
	{
		minValue: 0.500000,
		maxValue: 3.000000,
		defaultValue: 1.000000,
		labelCN: "道具采集量倍率",
		labelEN: "CollectionDropRate",
		toggleType: "slide"
	},
	{
		minValue: 0.500000,
		maxValue: 3.000000,
		defaultValue: 1.000000,
		labelCN: "可采集物品生命值倍率",
		labelEN: "CollectionObjectHpRate",
		toggleType: "slide"
	},
	{
		minValue: 0.500000,
		maxValue: 3.000000,
		defaultValue: 1.000000,
		labelCN: "可采集物品刷新间隔",
		labelEN: "CollectionObjectRespawnSpeedRate",
		toggleType: "slide"
	},
	{
		minValue: 0.100000,
		maxValue: 3.000000,
		defaultValue: 1.000000,
		labelCN: "道具掉落量倍率",
		labelEN: "EnemyDropItemRate",
		toggleType: "slide"
	},
	{
		minValue: 0,
		maxValue: 3,
		defaultValue: 2,
		labelCN: "死亡惩罚",
		labelEN: "DeathPenalty",
		toggleType: "select",
		options: {
			0: {
				CN: "不掉落任何东西",
				EN: "None",
			},
			1: {
				CN: "掉落装备以外的道具",
				EN: "Item",
			},
			2: {
				CN: "掉落所有道具",
				EN: "ItemAndEquipment",
			},
			3: {
				CN: "掉落所有道具及队伍内帕鲁",
				EN: "All",
			},
		},
	},
	{
		defaultValue: false,
		labelCN: "是否开启玩家间伤害",
		labelEN: "bEnablePlayerToPlayerDamage",
		toggleType: "toggle",
		unsure: true
	},
	{
		defaultValue: false,
		labelCN: "是否开启PVP",
		labelEN: "bIsPvP",
		toggleType: "toggle",
		unsure: true
	},
	{
		defaultValue: false,
		labelCN: "是否开启友军伤害",
		labelEN: "bEnableFriendlyFire",
		toggleType: "toggle",
		unsure: true
	},
	{
		defaultValue: true,
		labelCN: "是否会发生袭击事件",
		labelEN: "bEnableInvaderEnemy",
		toggleType: "toggle",
	},
	{
		defaultValue: true,
		labelCN: "是否启用辅助瞄准(手柄)",
		labelEN: "bEnableAimAssistPad",
		toggleType: "toggle",
	},
	{
		defaultValue: true,
		labelCN: "是否启用辅助瞄准(键鼠)",
		labelEN: "bEnableAimAssistKeyboard",
		toggleType: "toggle",
	},
	{
		minValue: 0,
		maxValue: 5000,
		defaultValue: 3000,
		labelCN: "世界内的掉落物上限",
		labelEN: "DropItemMaxNum",
		toggleType: "slide"
	},
	{
		minValue: 1,
		maxValue: 256,
		defaultValue: 128,
		labelCN: "最大终端数量",
		labelEN: "BaseCampMaxNum",
		toggleType: "slide"
	},
	{
		minValue: 1,
		maxValue: 20,
		defaultValue: 15,
		labelCN: "可分派至据点工作的帕鲁数量上限",
		labelEN: "BaseCampWorkerMaxNum",
		toggleType: "slide"
	},
	{
		minValue: 1,
		maxValue: 100,
		defaultValue: 20,
		labelCN: "公会人数上限",
		labelEN: "GuildPlayerMaxNum",
		toggleType: "slide"
	},
	{
		minValue: 0.100000,
		maxValue: 240.000000,
		defaultValue: 72.000000,
		labelCN: "巨大蛋孵化所需时间（小时）",
		labelEN: "PalEggDefaultHatchingTime",
		toggleType: "slide"
	},
	{
		minValue: 0.100000,
		maxValue: 3.000000,
		defaultValue: 1.000000,
		labelCN: "工作效率倍率",
		labelEN: "EnemyDropItemRate",
		toggleType: "slide"
	},
	{
		defaultValue: true,
		labelCN: "是否可拾取其他工会死亡掉落物",
		labelEN: "bCanPickupOtherGuildDeathPenaltyDrop",
		toggleType: "toggle",
	},
	{
		defaultValue: true,
		labelCN: "是否启用离线惩罚",
		labelEN: "bEnableNonLoginPenalty",
		toggleType: "toggle",
	},
	{
		defaultValue: true,
		labelCN: "是否启用快速旅行",
		labelEN: "bEnableFastTravel",
		toggleType: "toggle",
	},
	{
		defaultValue: true,
		labelCN: "是否根据地图选择初始位置",
		labelEN: "bIsStartLocationSelectByMap",
		toggleType: "toggle",
	},
	{
		minValue: 1,
		maxValue: 8,
		defaultValue: 4,
		labelCN: "领主塔可容纳玩家数",
		labelEN: "CoopPlayerMaxNum",
		toggleType: "slide"
	},
	{
		minValue: 2,
		maxValue: 32,
		defaultValue: 16,
		labelCN: "服务器最大玩家数",
		labelEN: "ServerPlayerMaxNum",
		toggleType: "slide"
	},
	{
		defaultValue: "Palworld Server",
		labelCN: "服务器名称",
		labelEN: "ServerName",
		toggleType: "input"
	},
	{
		defaultValue: "Happy Palworld！",
		labelCN: "服务器描述",
		labelEN: "ServerDescription",
		toggleType: "input"
	},
	{
		defaultValue: "",
		labelCN: "管理员密码",
		labelEN: "AdminPassword",
		toggleType: "input"
	},
	{
		defaultValue: "",
		labelCN: "服务器密码",
		labelEN: "ServerPassword",
		toggleType: "input"
	},
	{
		defaultValue: 8211,
		labelCN: "端口",
		labelEN: "PublicPort",
		toggleType: "input"
	},
	{
		defaultValue: "",
		labelCN: "IP地址",
		labelEN: "PublicIP",
		toggleType: "input"
	},
	{
		defaultValue: false,
		labelCN: "是否启用RCON",
		labelEN: "RCONEnabled",
		toggleType: "toggle",
	},
	{
		defaultValue: 25575,
		labelCN: "RCON端口",
		labelEN: "RCONPort",
		toggleType: "input"
	},
])
watch(IniData, () => {
	iniStrings.value = ''
})

onMounted(()=>{
	console.log(`
██████╗  ███████╗ ██╗ ██████╗   █████╗  ████████╗  ██████╗
██╔══██╗ ██╔════╝ ██║ ██╔══██╗ ██╔══██╗ ╚══██╔══╝ ██╔═══██╗
██████╔╝ █████╗   ██║ ██████╔╝ ███████║    ██║    ██║   ██║
██╔═══╝  ██╔══╝   ██║ ██╔══██╗ ██╔══██║    ██║    ██║   ██║
██║      ███████╗ ██║ ██║  ██║ ██║  ██║    ██║    ╚██████╔╝
╚═╝      ╚══════╝ ╚═╝ ╚═╝  ╚═╝ ╚═╝  ╚═╝    ╚═╝     ╚═════╝

https://space.bilibili.com/7277347
`);
})
</script>

<style scoped>
.iniContainer {
	color: white;
	width: 70%;
	position: relative;
	z-index: 999;
	display: flex;
	align-items: flex-end;
	flex-direction: column;
}

.iniGenCon {
	height: 100vh;
	overflow: hidden;
	display: flex;
	align-items: center;
	flex-direction: column;
}

.iniScroll {
	margin-right: 8px;
}

.iniItems {
	padding-right: 16px;
}

.iniItemOp {
	display: flex;
	align-items: center;
	background: rgba(10, 10, 10, 0.55);
	margin: 4px 0;
	padding: 8px 36px;
	border: 1px solid #00000000;
	transition: 200ms ease-in-out;
	font-size: 16px;
}

.iniItemOp:hover {
	border: 1px solid #77ccff;
}

.left {
	width: 100%;
	height: 38px;
	display: flex;
	align-items: center;
}

.right {
	width: 50%;
}

.center {
	text-align: center;
	width: 10%;
}

.initoggle .right {
	display: flex;
	justify-content: space-between;
}

.background {
	position: fixed;
	left: 0;
	right: 0;
	top: 0;
	bottom: 0;
	z-index: 1;
	filter: blur(10px) brightness(0.7);
}

.bg {
	height: 100%;
	width: 100%;
	background-position: center;
	background-size: cover;
}

:deep(.n-scrollbar-rail__scrollbar) {
	--n-scrollbar-border-radius: 0;
	--n-scrollbar-color: rgba(255, 255, 255, 0.6);
	--n-scrollbar-color-hover: rgba(255, 255, 255, 0.9);
}

:deep(.n-slider .n-slider-rail) {
	--n-fill-color: rgba(255, 255, 255, 0.9) !important;
	--n-fill-color-hover: rgba(255, 255, 255, 1) !important;
	--n-rail-height: 3px !important;
	border-radius: 0 !important;
}

:deep(.n-slider-handle) {
	background-color: rgba(0, 0, 0, 0) !important;
	border-radius: 0 !important;
	height: 26px !important;
	width: 26px !important;
	display: flex;
	justify-content: center;
	align-items: center;
	box-shadow: none !important;
}

:deep(.n-slider-handle::before) {
	content: "" !important;
	display: block;
	height: 15px !important;
	width: 15px !important;
	background-color: #ddd !important;
	border: white 4px solid;
	box-sizing: border-box;
	transform: rotate(45deg);
	transition: 300ms;
}

:deep(.n-slider-handle:hover::before) {
	border: #2fb4ff 4px solid !important;
	background-color: #92d4fb !important;
}

.right .btn, .uploadFunc {
	width: 49%;
	display: flex;
	justify-content: center;
	align-items: center;
	height: 32px;
	cursor: pointer;
	transition: 300ms;
	border-bottom: 1px solid rgba(255, 255, 255, 0.2);
	color: #ccc;
}

.right .btn:hover, .uploadFunc:hover {
	border-bottom: 1px solid #7bc6e6 !important;
	background-color: rgba(103, 157, 178, 0.35) !important;
	color: #eee;
}

.toggleActive {
	background-color: rgba(4, 137, 181, 0.7) !important;
	border-bottom: 1px solid #26BDF9 !important;
	color: #fff !important;
}

.iniinput .right {
	border: none;
	border-bottom: 1px solid rgba(255, 255, 255, 0.2);
	color: #eee;
	background-color: #00000000;
	padding: 3px 10px;
	box-sizing: border-box;
	transition: 300ms;
}

.iniinput .right:focus, .iniinput .right:hover {
	outline: none;
	border-bottom: 1px solid #7bc6e6 !important;
	background-color: rgba(103, 157, 178, 0.35) !important;
}

.mBottom {
	display: flex;
	justify-content: space-between;
	position: absolute;
	left: 0;
	right: 0;
	top: 0;
	bottom: 0;
}

.bLine {
	height: 150%;
	border-bottom: 2px solid #eeeeeeaa;
	width: 100%;
	margin: 0 2px;
	cursor: pointer;
}

.mTop {
	text-align: center;
	font-size: 16px;
}

.bLineActive {
	border-bottom: 2px solid #26BDF9 !important;
	transition: 300ms;
}

.middleInfo {
	position: relative;
	height: 100%;
	width: 100%;
	margin-bottom: 6px;
}

.iniselect .right {
	display: flex;
	align-items: center;
}

.toLeft, .toRight {
	width: 22px;
	display: flex;
	align-items: center;
	justify-content: center;
	color: #eee;
	cursor: pointer;
	transition: 300ms;
}

.toRight {
	transform: rotate(180deg);
}

.toLeft:hover, .toRight:hover {
	color: #26BDF9;
}

.title {
	z-index: 9;
	width: 90vw;
	height: 15vh;
	background-repeat: no-repeat;
	background-size: cover;
	background-position: top center;
	background-image: url("https://pic.imgdb.cn/item/65b2ad07871b83018a433e4d.png");
}

.ending {
	display: flex;
	position: relative;
	z-index: 9;
	justify-content: space-around;
	width: 50%;
	margin-top: 32px;
}

.ending .btn,.analyzeContainer .btn {
	color: #eee;
	background: rgba(0, 0, 0, 0.55);
	padding: 8px 5vw;
	font-size: 16px;
	border-left: 1px solid rgba(255, 255, 255, 0.8);
	border-right: 1px solid rgba(255, 255, 255, 0.8);
	border-top: 1px solid rgba(255, 255, 255, 0.3);
	border-bottom: 1px solid rgba(255, 255, 255, 0.3);
	transition: 150ms ease-in-out;
	cursor: pointer;
	text-align: center;
}

.ending .btn:hover,.analyzeContainer .btn:hover {
	color: #eee;
	background: rgba(0, 0, 0, 0.7);
	font-size: 16px;
	border: 1px solid #26BDF9;
}

.uploadIni {
	padding: 12px 36px;
	width: 98%;
}

.uploadIni .right {
	display: flex;
	width: 30%;
}

:deep(.n-upload-trigger) {
	width: 100%;
}

.uploadBtn {
	width: 100%;
	text-align: center;
}

.uploadFunc {
	margin: 0 6px;
}

.inputarea:focus-visible {
	outline: none;
}

.inputarea {
	background-color: #22222299;
	border: 1px solid #333;
	border-radius: 3px;
	font-size: 15px;
	color: #eeeeee;
	padding: 12px;
	max-height: 50vh;
}

:deep(.inputarea.n-input) {
	--n-caret-color: #26BDF9 !important;
	--n-border-hover: 1px solid #26BDF9 !important;
	--n-border-focus: 1px solid #26BDF9 !important;
	--n-box-shadow-focus: 0 0 8px 0 rgba(103, 157, 178, 0.35) !important;
	--n-color-focus: rgba(103, 157, 178, 0.1) !important;
}
.analyzeCard {
	background-color: rgba(0, 0, 0, 0.5);
	border: 1px solid #393939;
	border-radius: 3px;
	font-size: 22px;
	color: #eeeeee;
	padding: 5vh 2vw;
	width: 50vw;

}
.analyzeContainer .btn {
	margin-top: 3.5vh;
}
</style>