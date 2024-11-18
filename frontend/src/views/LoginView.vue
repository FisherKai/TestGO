<template>
  <div class="login-container">
    <div class="login-box">
      <h2>欢迎登录</h2>
      <form @submit.prevent="handleLogin">
        <div class="input-group">
          <label for="username">用户名</label>
          <input type="text" id="username" v-model="username" required />
        </div>
        <div class="input-group">
          <label for="password">密码</label>
          <input type="password" id="password" v-model="password" required />
        </div>
        <button type="submit">登录</button>
      </form>
    </div>
  </div>
</template>

<script>
import api from "../api/api.js"

export default {
  data() {
    return {
      username: '',
      password: ''
    };
  },
  methods: {
    async handleLogin() {
      try {
        const token = await api.login(this.username, this.password);
        console.log('登录成功，token:', token);
        // 登录成功后，可以重定向到主页或其他页面
        this.$router.push('/');
      } catch (error) {
        console.error('登录失败:', error.message);
      }
    }
  }
};
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #87CEEB; /* 天蓝色 */
}

.login-box {
  width: 300px;
  padding: 40px;
  background-color: #fff;
  box-shadow: 0 0 15px rgba(0, 0, 0, .1);
  border-radius: 8px;
  text-align: center;
}

.input-group {
  margin-bottom: 15px;
}

.input-group label,
.input-group input {
  display: block;
  width: 100%;
  margin-top: 5px;
}

.input-group input {
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

button {
  width: 100%;
  padding: 10px;
  background-color: #007BFF;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

button:hover {
  background-color: #0056b3;
}
</style>