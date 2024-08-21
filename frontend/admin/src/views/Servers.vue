<template>
  <div class="container">
    <div class="topBar">
      <div class="spacer">
        <a-button type="default" @click="logout">退出</a-button>
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
        class="inputField"
      />
      <a-button type="primary" @click="search" class="btn">搜索</a-button>
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
            <td>
              <a-button type="link" @click="handleAction(server.id)"
                >操作</a-button
              >
            </td>
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
          `servers/search?query=${encodeURIComponent(this.searchText)}`
        )
        const data = response.data

        if (data.message) {
          this.$message.error(data.message)
        } else {
          this.searchResults = data.data || []
        }
      } catch (error) {
        console.error(error)
        this.$message.error('搜索服务器时出错')
      }
    },
    handleAction(serverId) {
      // 处理操作逻辑
      console.log(`操作服务器 ID: ${serverId}`)
    }
  }
}
</script>

<style scoped>
.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  background-color: #f0f2f5;
  padding: 20px;
  min-height: 100vh;
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
  margin: 2rem 0;
  font-family: 'Courier New', Courier, monospace;
}

.searchBox {
  display: flex;
  align-items: center;
  margin-bottom: 20px;
}

.inputField {
  border-radius: 5px;
  border: 1px solid #d9d9d9;
  transition: border-color 0.3s;
}

.inputField:focus {
  border-color: #40a9ff;
  box-shadow: 0 0 5px rgba(64, 169, 255, 0.5);
}

.btn {
  margin-left: 10px;
  border-radius: 5px;
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
