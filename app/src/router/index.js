import Vue from 'vue'
import Router from 'vue-router'
import Top from '@/components/Top'
import Signin from '@/components/Signin'
import Signup from '@/components/Signup'
import Home from '@/components/Home'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Top',
      component: Top
    },
    {
      path: '/signin',
      name: 'Signin',
      component: Signin
    },
    {
      path: '/signup',
      name: 'Signup',
      component: Signup
    },
    {
      path: '/home',
      name: 'home',
      component: Home
    }
  ]
})
