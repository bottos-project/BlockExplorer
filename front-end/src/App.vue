<template>
  <div id="app">
    <MHeader></MHeader>
    <router-view></router-view>
    <div class="footer-sink"></div>
    <MFooter></MFooter>
  </div>
</template>

<script>
import MHeader from './components/Header.vue';
import MFooter from './components/Footer.vue';
export default {
  name: 'app',
  components:{
    MHeader,MFooter
  }
}
</script>