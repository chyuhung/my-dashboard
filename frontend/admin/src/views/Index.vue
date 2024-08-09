<template>
  <div class="container">
    <div class="topBar">
      <div class="spacer">
        <a-button type="primary" @click="logout">退出</a-button>
      </div>
    </div>
    <div class="title">
      <h1>mydashboard</h1>
    </div>
    <div class="searchBox">
      <a-input
        v-model="searchText"
        placeholder="输入虚拟机IP或名称进行搜索"
        suffix-icon="search"
        @pressEnter="search"
        style="width: 300px"
      />
      <a-button type="primary" @click="search" style="margin-left: 10px">
        搜索
      </a-button>
    </div>
    <div class="searchResults" v-if="searchResults.length">
      <table>
        <thead>
          <tr>
            <th>名称</th>
            <th>主机</th>
            <th>规格</th>
            <th>镜像</th>
            <th>操作</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="server in searchResults" :key="server.id">
            <td>{{ server.name }}</td>
            <td>{{ server.host }}</td>
            <td>{{ server.flavor }}</td>
            <td>{{ server.image }}</td>
          </tr>
        </tbody>
      </table>
    </div>
    <div v-else class="no-results">
      <p>未找到相关服务器</p>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      searchText: '',
      searchResults: []
    }
  },
  methods: {
    logout() {
      window.sessionStorage.clear()
      this.$router.push('/login')
    },
    async search() {
      if (!this.searchText) {
        this.$message.warning('查询参数不能为空')
        return
      }
      try {
        const response = await this.$http.get(
          `servers/search?query=${this.searchText}`
        )
        const data = response.data

        if (data.message) {
          this.$message.error(data.message)
        } else {
          this.searchResults = data.servers || []
        }
      } catch (error) {
        console.error(error)
        this.$message.error('搜索服务器时出错')
      }
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
  font-size: 48px;
  font-weight: bold;
  margin-bottom: 2rem;
  font-family: 'Courier New', Courier, monospace;
}

.searchBox {
  display: flex;
  align-items: center;
  margin-bottom: 20px;
}

.searchResults {
  margin-top: 20px;
  width: 100%;
}

.no-results {
  margin-top: 20px;
  color: red;
}
</style>
