<template>
  <div>
    <ManageHeader></ManageHeader>
    <div class="container">
      <div class="card events-card">
        <header class="card-header">
          <p class="card-header-title">
            <span>用户列表</span>
          </p>
        </header>
        <div class="card-content">
          <div class="content has-text-centered	min-heights">
            <div class="buttons has-addons are-small">
              <button
                class="button"
                v-for="(it) in userlist"
                :key="it.ID"
                :class="currentSelected == it.ID?'is-success is-selected':''" @click="()=>{setCurrentSelect(it.ID)}">{{it.UserName}}</button>
            </div>
          </div>
        </div>
      </div>

      <div class="card">
        <header class="card-header">
          <p class="card-header-title">
            导入数据
          </p>
        </header>
        <div class="card-content">
          <div class="content">
            <div class="control">
              分隔符选择：
              <label class="radio">
                <input type="radio" v-model="splitstr" value="0" name="splitstr">
                Tab
              </label>
              <label class="radio">
                <input type="radio" v-model="splitstr" value="1" name="splitstr">
                ----
              </label>
              <label class="radio">
                <input type="radio" v-model="splitstr" value="2" name="splitstr">
                空格
              </label>
            </div>
            <div class="field has-addons">
              <textarea class="textarea" rows="10" v-model="seleteData" placeholder="粘贴要查询的帐号到这里"></textarea>
            </div>
            <div class="buttons">
              <button class="button is-primary" :class="loading?'is-loading':''" :disabled="userlist.length > 0?false:true" @click="PostData">导入</button>
              <button class="button is-light" @click="CleanSeleteData">清空</button>
            </div>
          </div>
        </div>
      </div>

      <div class="card events-card mt-5">
        <header class="card-header">
          <p class="card-header-title">
            <span>添加用户</span>
          </p>
        </header>
        <div class="card-content">
          <div class="content">
            <div class="field has-addons">
              <p class="control is-expanded">
                <input class="input" type="text" v-model="createUser" placeholder="用户名">
              </p>
              <p class="control">
                <button class="button is-info"
                  @click="onSubmit"
                  :class="loading ? 'is-loading' : ''">
                  添加
                </button>
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
    <NotIfication
      :showData="openerr" />
  </div>
</template>
<script>
import { reactive, toRefs, defineComponent, onBeforeMount } from 'vue'
import { useRouter } from 'vue-router'
import ManageHeader from '@/components/Other/Header'
import NotIfication from "@/components/Other/Notification"

import Fetch from '@/helper/fetch'
import Config from '@/helper/config'
export default defineComponent({
  name: 'AccList',
  components: { ManageHeader, NotIfication },
  setup() {
    const router = useRouter()
    let states = reactive({
      loading: false,
      temp: [],
      userlist: [],
      data: [],
      total: 0,
      createUser: "",
      currentSelected: 0,
      rootUrl: Config.RootU,
      seleteData: "",
      splitstr: "0",
      path:router.currentRoute.value.path,
      title: "",
      openerr: {
        active: false,
        message: "",
        color: "",
        newtime: 0,
      }
    })
    onBeforeMount(async()=>{
      states.loading = true
      states.title = "上传数据"
      document.title = `${Config.GlobalTitle}-${states.title}`
      await GetData()
    })

    const GetData = async() => {
      let url = Config.Api.UserList
      states.loading = true
      const d = await Fetch(url, {}, 'GET')
      if (d.status == 0) {
        states.userlist = d.userList
        states.loading = false
        if (d.userList.length > 0) {
          states.currentSelected = d.userList[0].ID
        }
      }else{
        states.userlist = []
        states.loading = false
      }
    }

    const onSubmit = async() => {
      let url = Config.Api.AddUser
      states.loading = true
      const params = {
        username: states.createUser
      }
      const d = await Fetch(url, params, 'POST')
      if (d.status == 0) {
        states.loading = false
        states.createUser = ""
        GetData()
      }else{
        states.createUser = ""
        states.loading = false
      }
    }
    const setCurrentSelect = (id) => {
      states.currentSelected = id
    }

    const CleanSeleteData = () => {
      states.seleteData = ""
    }
    const ShowMessage = (e) => {
      states.openerr = e
    }
    const PostData = async() => {
      let url = Config.Api.PostAccount
      states.loading = true
      const params = {
        userid: states.currentSelected.toString(),
        data: states.seleteData,
        splitstr: states.splitstr
      }
      const d = await Fetch(url, params, 'POST')
      let color = "is-success"
      if (d.status == 1) color = "is-danger"
      const e = {
        active: true,
        message: d.message,
        color: color,
        newtime: 0,
      }
      ShowMessage(e)
      states.loading = false
    }

    return {
      ...toRefs(states),
      onSubmit,
      setCurrentSelect,
      CleanSeleteData,
      PostData
    }
  },
})
</script>

<style scoped>
.statesimg {
  width: 150px;
  height: 30px;
  max-width: 150px;
  max-height: 30px;
  display: block;
}
.error {
  position:absolute;
  right: 0px;
  width: 240px;
  z-index: 9999999999;
}
.w350 {
  width: 350px;
  margin-top: 30px
}
.is-img {
  position: absolute;
  display: block;
  right: 10px;
  max-width: 340px;
  width: 340px;
}
.autohidden {
  max-height: 300px;
  overflow-x: auto;
}
</style>
