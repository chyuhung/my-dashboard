<template>
  <div>
    <h1>创建服务器</h1>
    <form @submit.prevent="handleSubmit">
      <div v-if="step === 1">
        <label for="hostname">主机名:</label>
        <input type="text" v-model="hostname" required />
        <button @click="nextStep">下一步</button>
      </div>

      <div v-if="step === 2">
        <label for="image">选择镜像:</label>
        <select v-model="selectedImage" required>
          <option v-for="image in images" :key="image.id" :value="image.id">
            {{ image.name }}
          </option>
        </select>
        <button @click="nextStep">下一步</button>
      </div>

      <div v-if="step === 3">
        <label for="flavor">选择规格:</label>
        <select v-model="selectedFlavor" required>
          <option v-for="flavor in flavors" :key="flavor.id" :value="flavor.id">
            {{ flavor.name }}
          </option>
        </select>
        <button @click="nextStep">下一步</button>
      </div>

      <div v-if="step === 4">
        <label for="volume">选择卷类型:</label>
        <select v-model="selectedVolumeType" required>
          <option v-for="type in volumeTypes" :key="type" :value="type">
            {{ type }}
          </option>
        </select>
        <button @click="nextStep">下一步</button>
      </div>

      <div v-if="step === 5">
        <label for="network">选择网络:</label>
        <select v-model="selectedNetwork" required>
          <option
            v-for="network in networks"
            :key="network.id"
            :value="network.id"
          >
            {{ network.name }}
          </option>
        </select>
        <button @click="submit">创建服务器</button>
      </div>
    </form>
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
      networks: []
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
    async fetchImages() {
      const response = await axios.get('/api/images')
      this.images = response.data
    },
    async fetchFlavors() {
      const response = await axios.get('/api/flavors')
      this.flavors = response.data
    },
    async fetchVolumeTypes() {
      const response = await axios.get('/api/volumetypes')
      this.volumeTypes = response.data
    },
    async fetchNetworks() {
      const response = await axios.get('/api/networks')
      this.networks = response.data
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
        const response = await axios.post('/api/create-server', serverData)
        alert('服务器创建成功:' + response.data.id)
      } catch (error) {
        console.error('创建服务器失败:', error)
      }
    }
  }
}
</script>
