<template>
  <div class="container" @keydown.enter="login">
    <div class="logo">
      <h1>OpenStack</h1>
    </div>
    <div class="loginBox">
      <a-form-model
        ref="loginFormRef"
        :rules="rules"
        :model="formdata"
        class="loginForm"
      >
        <a-form-model-item prop="username">
          <a-input
            v-model="formdata.username"
            placeholder="请输入用户名"
            class="inputField"
          >
            <a-icon slot="prefix" type="user" class="inputIcon" />
          </a-input>
        </a-form-model-item>
        <a-form-model-item prop="password">
          <a-input
            v-model="formdata.password"
            placeholder="请输入密码"
            type="password"
            class="inputField"
          >
            <a-icon slot="prefix" type="lock" class="inputIcon" />
          </a-input>
        </a-form-model-item>
        <a-form-model-item class="loginBtn">
          <a-button type="primary" @click="login" class="btn">登录</a-button>
          <a-button type="default" @click="resetForm" class="btn"
            >取消</a-button
          >
        </a-form-model-item>
      </a-form-model>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      formdata: {
        username: '',
        password: ''
      },
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          {
            min: 4,
            max: 12,
            message: '用户名必须在4到12个字符之间',
            trigger: 'blur'
          }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          {
            min: 6,
            max: 20,
            message: '密码必须在6到20个字符之间',
            trigger: 'blur'
          }
        ]
      }
    }
  },
  methods: {
    resetForm() {
      this.$refs.loginFormRef.resetFields()
    },
    login() {
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
  height: 100%;
  background-color: #f0f2f5;
  display: flex;
  align-items: center;
  justify-content: center;
}

.logo {
  text-align: center;
  margin-right: 50px;
}

.logo h1 {
  font-size: 48px;
  font-weight: bold;
  margin-bottom: 2rem;
  font-family: 'Courier New', Courier, monospace;
}

.loginBox {
  width: 450px;
  background-color: #ffffff;
  border-radius: 10px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  padding: 30px;
  box-sizing: border-box;
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

.inputIcon {
  color: rgba(0, 0, 0, 0.25);
}

.btn {
  width: 100%;
  margin: 10px 0;
  border-radius: 5px;
}

.btn:first-child {
  background-color: #1890ff;
  color: white;
}

.btn:last-child {
  background-color: #f5f5f5;
  color: #595959;
}
</style>
