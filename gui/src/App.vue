<template>
  <div id="app">
    <h1>Last api calls</h1>
    <table>
      <tr>
        <th>user.id</th>
        <th>user.email</th>
        <th>user.token</th>
        <th>user.requests</th>
      </tr>
      <tr v-for="user in users" :key="user.id">
        <td>{{ user.id }}</td>
        <td>{{ user.email }}</td>
        <td>{{ user.token }}</td>
        <td>{{ user.requests }}</td>
      </tr>
    </table>
    <hr>
    <table>
      <tr>
        <th>id</th>
        <th>url</th>
        <th>user</th>
      </tr>
      <tr v-for="call in calls" v-bind:key="call.id">
        <td>{{ call.id }}</td>
        <td><a :href="call.url" :title="call.url" target="_blank">link</a></td>
        <td>{{ call.user.email }}</td>
      </tr>
    </table>
    <hr>

    <h1>Proxy CRUD</h1>
    <select v-model="form.type">
      <option value="http">http</option>
      <option value="sock4">sock4</option>
      <option value="sock5">sock5</option>
    </select>
    <input size="40" v-model="form.address" type="text" placeholder="login:password@address:port">
    <input size="20" v-model="form.country" type="text" placeholder="ru,ua,de,us,uk etc" maxlength="3"
           @keypress.enter="sendProxy">
    <button @click="sendProxy">submit</button>

    <table>
      <tr>
        <th>id</th>
        <th>type</th>
        <th style="min-width: 250px">address</th>
        <th>port</th>
        <th>login</th>
        <th>password</th>
        <th>country</th>
        <th>delete</th>
      </tr>
      <tr v-for="proxy in proxies" v-bind:key="proxy.id">
        <td>{{ proxy.id }}</td>
        <td>{{ proxy.type }}</td>
        <td>{{ proxy.address }}</td>
        <td>{{ proxy.port }}</td>
        <td>{{ proxy.login }}</td>
        <td>{{ proxy.password }}</td>
        <td>{{ proxy.country }}</td>
        <td @click="rmProxy(proxy.id)" class="remove">x</td>
      </tr>
    </table>
  </div>
</template>

<script>

export default {
  name: 'App',
  data() {
    return {
      auth: false,
      proxies: null,
      calls: null,
      users: null,
      form: {
        type: 'http',
        address: null,
        country: null
      }
    }
  },
  methods: {
    async sendProxy() {
      let form = Object.assign({}, this.$data.form)

      // credentials and proxy address
      let sep = form.address.split('@')
      let cred = sep[0].split(':')
      let proxy = sep[1].split(':')

      // log pass
      let login = cred[0]
      let password = cred[1]

      // address port
      let address = proxy[0]
      let port = proxy[1]

      console.log(this.$data.form.country,
          this.$data.form.type,
          login,
          password,
          address,
          port,)

      const res = await fetch(`/api/proxy/create`, {
        method: 'POST',
        credentials: 'include',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          country: this.$data.form.country,
          type: this.$data.form.type,
          login: login,
          password: password,
          address: address,
          port: port,
        })
      })
      const json = await res.json()
      if (json.status !== 'ok') {
        alert('some error: ' + json.msg)
      }
      await this.updateProxies()

    },
    async rmProxy(id) {
      const res = await fetch(`/api/proxy/${id}/delete`, {
        method: 'GET',
        credentials: 'include'
      })
      const json = await res.json()
      if (json.status !== 'ok') {
        alert('some error')
      }
      await this.updateProxies()
    },

    async updateProxies() {
      const res = await fetch('/api/proxies', {
        method: 'GET',
        credentials: 'include'
      })
      if (res.statusCode === 200) {
        this.$data.auth = true
        return
      }
      this.$data.proxies = await res.json()
    },

    async getCalls() {
      const res = await fetch('/api/calls', {
        method: 'GET',
        credentials: 'include'
      })
      this.$data.calls = await res.json()

      const users = await fetch('/api/users', {
        method: 'GET',
        credentials: 'include'
      })
      this.$data.users = await users.json()
    }
  },
  async mounted() {
    await this.updateProxies()
    await this.getCalls()
  }
}
</script>

<style>
body {
  font-family: Roboto, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  font-size: 18px;
  color: #2c3e50;
  margin: 60px auto 0;
  width: 900px;
}

table {
  width: 700px;
  margin: 40px auto;
}

td, th {
  padding: 10px;
}

.remove {
  color: red;
  background: #eee;
}

input, select, button {
  padding: 5px;
  font-size: 18px;
  margin: 0 10px 0 0;
}

</style>
