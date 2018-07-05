<template>
  <div class="top">
  <nav>
    <div class="nav-wrapper">
      <a href="#" class="brand-logo">Sample Line</a>
      <ul id="nav-mobile" class="right hide-on-med-and-down">
      </ul>
    </div>
  </nav>

  Sample application
  <div class="button">
  <a class="waves-effect waves-light btn-large" @click="Signin">Signin</a>
  <a class="waves-effect waves-light btn-large">Signup</a>
  </div>
  </div>
</template>


<script>
export default {
  name: 'Top',
  data () {
    return {
      msg: 'Welcome to Your Vue.js App'
    }
  }
}
</script>
