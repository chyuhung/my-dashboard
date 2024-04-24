<template>
  <div class="container">
    <div class="topBar">
      <div class="spacer">
        <a-button type="primary" @click="logout">退出</a-button>
      </div>
    </div>
    <div class="title">
      <h1 class="title">mydashboard</h1>
    </div>
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
    <div class="searchResults" v-if="searchResults">
      <table>
        <thead>
          <tr>
            <th>名称</th>
            <th>主机</th>
            <th>规格</th>
            <th>镜像</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>{{ searchResults.name }}</td>
            <td>{{ searchResults.host }}</td>
            <td>{{ searchResults.flavor }}</td>
            <td>{{ searchResults.image }}</td>
            <td>
              <a-button type="primary" @click="reinstall(searchResults)"
                >重装</a-button
              >
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      searchText: '',
      searchResults: null
    }
  },
  methods: {
    logout() {
      window.sessionStorage.clear('token')
      this.$router.push('/login')
    },
    async search() {
      try {
        const { data } = await this.$http.get(`instance/${this.searchText}`)
        if (data.message) {
          this.$message.error(data.message)
        } else {
          this.searchResults = data.data
        }
      } catch (error) {
        console.error(error)
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
  font-size: 90px;
  font-weight: bold;
  margin-bottom: 7%;
  font-family: 'Courier New', Courier, monospace;
}

.searchBox {
  display: flex;
  align-items: center;
}

.searchResults {
  margin-top: 20px;
}
</style>
