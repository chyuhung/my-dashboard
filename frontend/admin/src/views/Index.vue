<template>
  <div class="container">
    <!-- <div class="header"><Header></Header></div> -->
    <div class="topBar">
      <div class="spacer">
        <a-button type="primary" @click="logout">退出</a-button>
      </div>
    </div>
    <div class="title"><h1 class="title">mydashboard</h1></div>
    <div class="searchBox">
      <a-input
        v-model="searchText"
        placeholder="输入虚拟机IP或名称进行搜索"
        suffix-icon="search"
        @pressEnter="search"
        style="width: 300px"
      />
      <a-button type="primary" @click="search" style="margin-left: 10px"
        >搜索</a-button
      >
    </div>
    <div class="footer"><Footer></Footer></div>
  </div>
</template>

<script>
export default {
  methods: {
    logout() {
      window.sessionStorage.clear('token')
      this.$router.push('/login')
    },
    search() {
      this.$refs.loginFormRef.validate(async (valid) => {
        if (!valid) return this.$message.error('输入不合规，请按要求修改')
        const { data: res } = await this.$http.post('login', this.formdata)
        if (res.status !== 200) return this.$message.error(res.message)
        window.sessionStorage.setItem('token', res.token)
        this.$router.push('index')
        return this.$message.info('登录成功，欢迎回来')
      })
    }
  }
}
</script>

<style scoped>
.container {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.topBar {
  display: flex;
  align-items: center;
  width: 100%;
  padding: 20px;
}

.spacer {
  margin-left: auto;
}

.title {
  font-size: 90px;
  font-weight: bold;
  margin-bottom: 7%;
  font-family: 'Courier New', Courier, monospace;
}

.searchBox {
  display: flex;
  align-items: center;
}
</style>
