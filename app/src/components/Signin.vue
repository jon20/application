<template>
  <div class="signin">
  <nav>
    <div class="nav-wrapper">
      <a href="#" class="brand-logo">Sample Line</a>
      <ul id="nav-mobile" class="right hide-on-med-and-down">
      </ul>
    </div>
  </nav>

  <div class="row">
    <form class="col s12">
      <div class="row">
        <div class="input-field col s12">
          <input id="name" type="text" name="name" v-model="name" data-vv-as="name" >
          <label for="username">Name</label>
          <span class="helper-text" data-error="wrong" data-success="right"></span>
        </div>
      </div>
      <div class="row">
        <div class="input-field col s12">
          <input id="email" type="email" v-model="email"    data-vv-as="email" >
          <label for="email">Email</label>
          <span class="helper-text" data-error="wrong" data-success="right"></span>
        </div>
      </div>
      <div class="row">
        <div class="input-field col s12">
          <input id="password" type="password" v-model="password"  data-vv-as="password" >
          <label for="password">Password</label>
          <span class="helper-text" data-error="wrong" data-success="right"></span>
        </div>
      </div>
      <div class="row">
           <div class="input-field col s12">
            <button class="btn cyan waves-effect waves-light right" type="submit" v-on:click="signIn" >Create Account
              <i class="mdi-content-send right"></i>
            </button>
           </div>
      </div>
    </form>
  </div>

  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'signin',
  data () {
    return {
      name: '',
      email: '',
      password: ''
    }
  },
  methods: {
    postApp: function () {
      this.$validator.validateAll().then((result) => {
        if (result) {
          signIn()
          let params = new URLSearchParams()
          params.append('name', this.name)
          params.append('email', this.email)
          params.append('password', this.password)
          axios.post('http://localhost:8000/auth', params).then(res =>
            resp = res.data,
            axios.post('http://localhost:8000/login', resp)
          )
        }
      })
    },
    signIn: function () {
      let params = new URLSearchParams()
      params.append('name', this.name)
      params.append('email', this.email)
      params.append('password', this.password)
      axios.post('http://localhost:8000/auth', {'name': this.name, 'email': this.email, 'password': this.password}).then(res =>
        console.log("ok")  
      )

      axios.post('http://localhost:8000/login', {'username': this.name, 'email': this.email, 'password': this.password}).then(res =>
        localStorage.setItem('jwt', res.data.token),
        this.$router.push('/')

      )
    }
  }
}
</script>
