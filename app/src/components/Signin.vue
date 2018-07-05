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
          <input id="name" type="text" class="validate">
          <label for="username">Name</label>
        </div>
      </div>
      <div class="row">
        <div class="input-field col s12">
          <input id="email" type="email" class="validate">
          <label for="email">Email</label>
        </div>
      </div>
      <div class="row">
        <div class="input-field col s12">
          <input id="password" type="password" class="validate">
          <label for="password">Password</label>
        </div>
      </div>
      <div class="row">
           <div class="input-field col s12">
            <button class="btn cyan waves-effect waves-light right" type="submit" name="action">Submit
              <i class="mdi-content-send right"></i>
            </button>
           </div>
      </div>
    </form>
  </div>

  </div>
</template>

<script>
export default {
  name: 'signin',
  data () {
    return {
      msg: 'Welcome to Your Vue.js App'
    }
  }
}
</script>
