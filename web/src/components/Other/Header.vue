<template>
  <nav class="navbar is-white">
    <div class="container">
      <div class="navbar-brand">
        <span class="navbar-item brand-text">
          <img :src="logo">Account Manage
        </span>
      </div>
      <div id="navMenu" class="navbar-menu">
        <div class="navbar-start">
          <router-link class="navbar-item" :class="path === 'index' || path === 'index'?'is-active':''" to="/">
            首页
          </router-link>
          <router-link class="navbar-item" :class="path === 'postdata' || path === 'postdata'?'is-active':''" to="/postdata">
            上传数据
          </router-link>
          <router-link class="navbar-item" :class="path === 'userdata' || path === 'userdata'?'is-active':''" to="/userdata">
            用户数据
          </router-link>
        </div>
      </div>
    </div>
  </nav>
</template>
<script>
import Config from '@/helper/config'
import { reactive, toRefs, defineComponent } from 'vue'
import { useRouter } from 'vue-router'
export default defineComponent({
  name: 'CardList',
  setup() {
    const router = useRouter()
    let states = reactive({
      logo: Config.images[2],
      path: router.currentRoute.value.name
    })
    return {
      ...toRefs(states),
      router
    }
  }
})
</script>

<style scoped>
nav.navbar {
  border-top: 4px solid #276cda;
}
</style>