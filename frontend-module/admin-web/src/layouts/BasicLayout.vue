<template>
  <a-layout class="layout">
    <a-layout-sider class="sider" width="240">
      <div class="brand">
        <div class="brand__title">Ovo Gaming</div>
        <div class="brand__sub">Admin Console</div>
      </div>

      <a-menu mode="inline" theme="dark" :selectedKeys="[currentPath]">
        <a-menu-item key="/dashboard" @click="go('/dashboard')">仪表盘</a-menu-item>
        <a-menu-item key="/banners" @click="go('/banners')">轮播图管理</a-menu-item>
        <a-menu-item key="/notices" @click="go('/notices')">公告管理</a-menu-item>
        <a-menu-item key="/packages" @click="go('/packages')">套餐管理</a-menu-item>
        <a-menu-item key="/orders" @click="go('/orders')">订单管理</a-menu-item>
      </a-menu>
    </a-layout-sider>

    <a-layout>
      <a-layout-header class="header">
        <div class="header__title">电竞陪玩后台</div>
        <a-button type="primary">管理员</a-button>
      </a-layout-header>
      <a-layout-content class="content">
        <router-view />
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
const router = useRouter()
const currentPath = computed(() => route.path)

function go(path: string) {
  router.push(path)
}
</script>

<style scoped>
.layout {
  min-height: 100vh;
}

.sider {
  background: linear-gradient(180deg, #101828 0%, #1e293b 100%);
}

.brand {
  padding: 28px 24px;
  color: #fff;
}

.brand__title {
  font-size: 22px;
  font-weight: 800;
}

.brand__sub {
  margin-top: 6px;
  font-size: 12px;
  color: rgba(255, 255, 255, 0.7);
}

.header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 24px;
  background: #fff;
  box-shadow: 0 6px 24px rgba(15, 23, 42, 0.06);
}

.header__title {
  font-size: 18px;
  font-weight: 700;
}

.content {
  margin: 24px;
  padding: 24px;
  border-radius: 20px;
  background: #fff;
  min-height: calc(100vh - 120px);
}
</style>
