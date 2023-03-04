<template>
  <div>
    <ManageHeader></ManageHeader>
    <div class="container">
      <div class="card">
        <header class="card-header">
          <p class="card-header-title">
            输入查询数据
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
              <button class="button is-primary" :class="loading?'is-loading':''" @click="SeleteData">查询</button>
              <button class="button is-light" @click="CleanSeleteData">清空</button>
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
        <div class="card events-card" v-if="data.length <= 0">
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

    </div>   
  </div>
</template>
<script>
import { reactive, toRefs, defineComponent, onBeforeMount } from 'vue'
import { useRouter } from 'vue-router'
import ManageHeader from '@/components/Other/Header'
import EmptyEd from '@/components/Other/Empty'
import FormaTime from '@/components/Other/FormaTime'
import LoadIng from '@/components/Other/Loading'

import Fetch from '@/helper/fetch'
import Config from '@/helper/config'
export default defineComponent({
  name: 'AccList',
  components: { ManageHeader, EmptyEd, FormaTime, LoadIng },
  setup() {
    const router = useRouter()
    let states = reactive({
      loading: false,
      data: [],
      total: 0,
      splitstr: "0",
      rootUrl: Config.RootU,
      seleteData: "",
      path:router.currentRoute.value.path,
      title: ""
    })
    onBeforeMount(async()=>{
      states.title = "首页"
      document.title = `${Config.GlobalTitle}-${states.title}`
    })

    const SeleteData = async() => {
      let url = Config.Api.FindAccount
      states.loading = true
      const params = {
        data: states.seleteData,
        splitstr: states.splitstr
      }
      const d = await Fetch(url, params, 'POST')
      if (d.status == 0) {
        states.loading = false
        states.data = d.data
      }else{
        states.data = []
        states.loading = false
      }
    }
    const CleanSeleteData = () => {
      states.seleteData = ""
    }

    return {
      ...toRefs(states),
      SeleteData,
      CleanSeleteData,
      FormaTime
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
