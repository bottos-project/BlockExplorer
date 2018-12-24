<template>
  <div class="header cf">
    <div class="inner cf">
      <div class="logo">
        <router-link to="/"><img src="../assets/logo.png" /></router-link>
      </div>
      <div class="menus">
        <ul>
          <li>
            <router-link to="/home">{{$t('menus.home')}}</router-link>
          </li>
          <li>
            <router-link to="/blockchain">{{$t('menus.blockchain')}}</router-link>
            <ul>
              <li>
                <router-link to="/blockchain/blocks">{{$t('menus.blocks')}}</router-link>
              </li>
              <li>
                <router-link to="/blockchain/transactions">{{$t('menus.transactions')}}</router-link>
              </li>
               <li>
                <router-link to="/blockchain/transfers">{{$t('menus.transfers')}}</router-link>
              </li>
              <li>
                <router-link to="/blockchain/accounts">{{$t('menus.accounts')}}</router-link>
              </li>
              <li>
                <router-link to="/blockchain/statistics">{{$t('menus.statistics')}}</router-link>
              </li>
            </ul>
          </li>
          <li>
            <router-link to="/contracts">{{$t('menus.contracts')}}</router-link>
          </li>
          <li>
            <router-link to="/nodes">{{$t('menus.nodes')}}</router-link>
          </li>
          <li>
            <router-link to="/supernode">{{$t('menus.supernode')}}</router-link>
          </li>
        </ul>
      </div>
      <a class="Mbtn" @click="Mmenus=!Mmenus">菜单</a>
      <div class="language">
        <ul>
          <li><a href="#"></a></li>
        </ul>
        <button class="languagebtn" @click="changeLanguage('en')">En</button>
        <button class="languagebtn cn" @click="changeLanguage('zh-cn')">中文</button>
      </div>
      
    </div>
    <transition name="fade">
    <div class="Mmenus" v-show="Mmenus" @click="MmenuHide">
      <ul>
        <li>
          <router-link to="/home">{{$t('menus.home')}}</router-link>
        </li>
        <li>
          <a href="javascript:;">{{$t('menus.blockchain')}}</a>
          <ul>
            <li>
              <router-link to="/blockchain/blocks">{{$t('menus.blocks')}}</router-link>
            </li>
            <li>
              <router-link to="/blockchain/transactions">{{$t('menus.transactions')}}</router-link>
            </li>
              <li>
              <router-link to="/blockchain/transfers">{{$t('menus.transfers')}}</router-link>
            </li>
            <li>
              <router-link to="/blockchain/accounts">{{$t('menus.accounts')}}</router-link>
            </li>
            <li>
              <router-link to="/blockchain/statistics">{{$t('menus.statistics')}}</router-link>
            </li>
          </ul>
        </li>
        <li>
          <router-link to="/nodes">{{$t('menus.nodes')}}</router-link>
        </li>
      </ul>
    </div>
    </transition>
  </div>
</template>
<script>
    export default {
        data(){
          return {
            Mmenus:false
          }
        },
        props:{
            back:{
                type:Boolean,
                default:false
            }
        },
        methods:{
          goBack(){
              this.$router.go(-1);
          },
          changeLanguage(val){
            this.$i18n.locale = val
            window.localStorage.setItem('language', val)
          },
          MmenuHide(){
            this.Mmenus = false
          }       
        }
    }
</script>
<style scoped lang="less">
  .header{
    .languagebtn {
      background-color: #333; border: 0; font-size: 12px; color: #fff;
      height: 20px; width: 40px;  line-height: 20px; cursor: pointer;
    }
    .languagebtn.cn {right: 60px;}
    .languagebtn:hover {
      opacity: 0.8;
    }
  }
  .fade-enter-active, .fade-leave-active {
    transition: opacity .5s;
  }
  .fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
    opacity: 0;
  }
</style>

