<template>
  <div class="container">
    <h1 class="title">创建服务器</h1>
    <form @submit.prevent="handleSubmit" class="form">
      <div v-if="step === 1" class="step">
        <label for="hostname">主机名:</label>
        <input type="text" v-model="hostname" required class="inputField" />
        <button type="button" @click="checkHostname" class="btn">
          检查主机名
        </button>
        <button type="button" @click="nextStep" class="btn">下一步</button>
      </div>

      <div v-if="step === 2" class="step">
        <label for="image">选择镜像:</label>
        <select v-model="selectedImage" required class="selectField">
          <option v-for="image in images" :key="image" :value="image">
            {{ image }}
          </option>
        </select>
        <button type="button" @click="checkImage" class="btn">检查镜像</button>
        <button type="button" @click="prevStep" class="btn">上一步</button>
        <button type="button" @click="nextStep" class="btn">下一步</button>
      </div>

      <div v-if="step === 3" class="step">
        <label for="flavor">选择规格:</label>
        <select v-model="selectedFlavor" required class="selectField">
          <option v-for="flavor in flavors" :key="flavor" :value="flavor">
            {{ flavor }}
          </option>
        </select>
        <button type="button" @click="checkFlavor" class="btn">检查规格</button>
        <button type="button" @click="prevStep" class="btn">上一步</button>
        <button type="button" @click="nextStep" class="btn">下一步</button>
      </div>

      <div v-if="step === 4" class="step">
        <label for="volume">选择卷类型:</label>
        <select v-model="selectedVolumeType" required class="selectField">
          <option v-for="type in volumeTypes" :key="type" :value="type">
            {{ type }}
          </option>
        </select>
        <button type="button" @click="checkVolumeType" class="btn">
          检查卷类型
        </button>
        <button type="button" @click="prevStep" class="btn">上一步</button>
        <button type="button" @click="nextStep" class="btn">下一步</button>
      </div>

      <div v-if="step === 5" class="step">
        <label for="network">选择网络:</label>
        <select v-model="selectedNetwork" required class="selectField">
          <option v-for="network in networks" :key="network" :value="network">
            {{ network }}
          </option>
        </select>
        <button type="button" @click="checkNetwork" class="btn">
          检查网络
        </button>
        <button type="button" @click="prevStep" class="btn">上一步</button>
        <button type="button" @click="submit" class="btn submitBtn">
          创建服务器
        </button>
      </div>
    </form>

    <div class="result">
      <h2>检查结果:</h2>
      <p>{{ checkResult }}</p>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      step: 1,
      hostname: '',
      selectedImage: '',
      selectedFlavor: '',
      selectedVolumeType: '',
      selectedNetwork: '',
      images: [],
      flavors: [],
      volumeTypes: [],
      networks: [],
      checkResult: ''
    }
  },
  mounted() {
    this.fetchImages()
    this.fetchFlavors()
    this.fetchVolumeTypes()
    this.fetchNetworks()
  },
  methods: {
    nextStep() {
      this.step++
    },
    prevStep() {
      if (this.step > 1) {
        this.step--
      }
    },
    async fetchImages() {
      const response = await axios.get('/images')
      this.images = response.data.data
    },
    async fetchFlavors() {
      const response = await axios.get('/flavors')
      this.flavors = response.data.data
    },
    async fetchVolumeTypes() {
      const response = await axios.get('/volumetypes')
      this.volumeTypes = response.data.data
    },
    async fetchNetworks() {
      const response = await axios.get('/networks')
      this.networks = response.data.data
    },
    async checkHostname() {
      try {
        const response = await axios.get(
          `/check/hostname?name=${this.hostname}`
        )
        this.checkResult = response.data.message
      } catch (error) {
        this.checkResult = '检查主机名失败'
      }
    },
    async checkImage() {
      try {
        const response = await axios.get(
          `/check/image?image=${this.selectedImage}`
        )
        this.checkResult = response.data.message
      } catch (error) {
        this.checkResult = '检查镜像失败'
      }
    },
    async checkFlavor() {
      try {
        const response = await axios.get(
          `/check/flavor?flavor=${this.selectedFlavor}`
        )
        this.checkResult = response.data.message
      } catch (error) {
        this.checkResult = '检查规格失败'
      }
    },
    async checkVolumeType() {
      try {
        const response = await axios.get(
          `/check/volume?type=${this.selectedVolumeType}`
        )
        this.checkResult = response.data.message
      } catch (error) {
        this.checkResult = '检查卷类型失败'
      }
    },
    async checkNetwork() {
      try {
        const response = await axios.get(
          `/check/network?network=${this.selectedNetwork}`
        )
        this.checkResult = response.data.message
      } catch (error) {
        this.checkResult = '检查网络失败'
      }
    },
    async submit() {
      const serverData = {
        hostname: this.hostname,
        image: this.selectedImage,
        flavor: this.selectedFlavor,
        volumeType: this.selectedVolumeType,
        network: this.selectedNetwork
      }

      try {
        const response = await axios.post('/create', serverData)
        alert('服务器创建成功:' + response.data.id)
      } catch (error) {
        console.error('创建服务器失败:', error)
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
  background-color: #f0f2f5;
  padding: 40px;
  min-height: 100vh;
  border-radius: 10px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
}

.title {
  font-size: 48px;
  font-weight: bold;
  margin-bottom: 2rem;
  font-family: 'Courier New', Courier, monospace;
  color: #333;
}

.form {
  width: 100%;
  max-width: 400px;
}

.step {
  margin-bottom: 20px;
  background: #fff;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

label {
  display: block;
  margin-bottom: 8px;
  font-weight: 500;
  color: #555;
}

.inputField,
.selectField {
  width: 100%;
  padding: 12px;
  border-radius: 5px;
  border: 1px solid #d9d9d9;
  margin-top: 10px;
  margin-bottom: 10px;
  font-size: 16px;
}

.inputField:focus,
.selectField:focus {
  border-color: #40a9ff;
  outline: none;
  box-shadow: 0 0 5px rgba(64, 169, 255, 0.5);
}

.btn {
  padding: 10px 20px;
  border-radius: 5px;
  background-color: #1890ff;
  color: white;
  border: none;
  cursor: pointer;
  margin-right: 10px;
  transition: background-color 0.3s;
}

.btn:hover {
  background-color: #40a9ff;
}

.submitBtn {
  background-color: #52c41a;
}

.submitBtn:hover {
  background-color: #73d13d;
}

.result {
  margin-top: 20px;
  width: 100%;
  max-width: 400px;
  background: #fff;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}
</style>
