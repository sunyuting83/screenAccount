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

      <div class="card events-card" v-if="loading">
        <header class="card-header">
          <p class="card-header-title">
            努力加载中.....
          </p>
        </header>
        <div class="content has-text-centered	min-heights" style="min-height: 11.3rem">
            <div class="com__box" v-if="loading" :style="loading? 'margin-top:5rem':''">
              <LoadIng></LoadIng>
            </div>
        </div>
      </div>
      <div v-else>
        <div class="card events-card mt-5" v-if="data.length <= 0">
          <header class="card-header">
            <p class="card-header-title">
              空空如也
            </p>
          </header>
          <div class="card-content">
            <div class="content has-text-centered	min-heights" style="min-height: 11.3rem">
              <EmptyEd></EmptyEd>
            </div>
          </div>
        </div>
        <div class="card events-card mt-5"  v-else>
          <header class="card-header">
            <p class="card-header-title">
              查询结果
            </p>
          </header>
          <div class="card-content">
            <div class="content has-text-centered	min-heights" style="min-height: 11.3rem">
              <table class="table is-striped is-hoverable is-fullwidth is-narrow has-text-left">
                <thead>
                  <tr>
                    <td>编号</td>
                    <td>帐号</td>
                    <td>手机</td>
                    <td>密码</td>
                    <td>金币</td>
                    <td>价格</td>
                    <td>用户</td>
                    <td>导入日期</td>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="(it,index) in data" :key="it.ID">
                    <td>{{index + 1}}</td>
                    <td>{{it.Account}}</td>
                    <td>{{it.Phone}}</td>
                    <td>{{it.Password}}</td>
                    <td>{{it.Gold}}</td>
                    <td>{{it.Price}}</td>
                    <td>{{it.User.UserName}}</td>
                    <td><FormaTime :DateTime="it.CreatedAt"></FormaTime></td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>

      <PaginAtion v-if="data.length > 0 && pageLoading === true" :total="total" :number="limit" :GetData="GetList"></PaginAtion>
    </div>
  </div>
</template>
<script>
import { reactive, toRefs, defineComponent, onBeforeMount } from 'vue'
import { useRouter } from 'vue-router'
import ManageHeader from '@/components/Other/Header'
import EmptyEd from '@/components/Other/Empty'
import LoadIng from '@/components/Other/Loading'
import FormaTime from '@/components/Other/FormaTime'
import PaginAtion from '@/components/Other/PaginAtion'

import Fetch from '@/helper/fetch'
import Config from '@/helper/config'
export default defineComponent({
  name: 'AccList',
  components: { ManageHeader, EmptyEd, LoadIng, FormaTime, PaginAtion },
  setup() {
    const router = useRouter()
    let states = reactive({
      loading: false,
      page: [],
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
      pageLoading: false,
      limit: 100
    })
    onBeforeMount(async()=>{
      states.loading = true
      states.title = "上传数据"
      document.title = `${Config.GlobalTitle}-${states.title}`
      await GetData()
    })

    const GetData = async() => {
      states.pageLoading = false
      let url = Config.Api.UserList
      const d = await Fetch(url, {}, 'GET')
      if (d.status == 0) {
        states.userlist = d.userList
        if (d.userList.length > 0) {
          states.currentSelected = d.userList[0].ID
          await GetList()
        }
      }else{
        states.userlist = []
      }
    }

    const GetList = async(page = 1) => {
      let url = Config.Api.AccountList
      states.loading = TextTrackCueList
      const params = {
        UserID: states.currentSelected,
        page: page,
        limit: states.limit,
      }
      const d = await Fetch(url, params, 'GET')
      if (d.status == 0) {
        states.data = d.data
        states.total = d.total
        states.pageLoading = true
        states.loading = false
      }else{
        states.data = []
        states.loading = false
      }
    }

    const setCurrentSelect = async(id) => {
      states.currentSelected = id
      states.pageLoading = false
      await GetList(1)
    }

    return {
      ...toRefs(states),
      setCurrentSelect,
      GetList
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
