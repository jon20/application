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
          <input id="name" type="text" v-model="name" class="validate">
          <label for="username">Name</label>
        </div>
      </div>
      <div class="row">
        <div class="input-field col s12">
          <input id="email" type="email" v-model="email" class="validate">
          <label for="email">Email</label>
        </div>
      </div>
      <div class="row">
        <div class="input-field col s12">
          <input id="password" type="password" v-model="password" class="validate">
          <label for="password">Password</label>
        </div>
      </div>
      <div class="row">
           <div class="input-field col s12">
            <button class="btn cyan waves-effect waves-light right" type="submit" v-on:click="signIn">Submit
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
    signIn: function() {
      let params = new URLSearchParams();
      params.append('name', this.name);
      params.append('email', this.email);
      params.append('password', this.password);
      axios.post('http://localhost:8000/auth', params).then(res =>
        console.log(res.data)
      )
    }
  }
}
</script>
