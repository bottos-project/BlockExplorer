<template>
  <transition name="mei-message-fade">
    <div v-if="show" :class="['mei-message',type? `mei-message-${ type }` : '']">
      <i v-if="type=='warning'" class="mei-message-icon">
      <svg t="1534130512063" class="icon" style="" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="7173" xmlns:xlink="http://www.w3.org/1999/xlink" width="40" height="40"><path d="M512 989.86666668C775.85066667 989.86666667 989.86666667 775.85066667 989.86666668 512c0-263.91893333-214.016-477.86666667-477.86666668-477.86666668-263.98720001 0-477.86666667 213.94773333-477.86666668 477.86666668C34.13333333 775.85066667 248.0128 989.86666667 512 989.86666668zM512 238.93333333c28.19413333 0 51.2 22.9376 51.2 51.2S540.19413333 341.33333333 512 341.33333334c-28.39893334 0-51.2-22.9376-51.2-51.20000001S483.60106667 238.93333333 512 238.93333333zM460.8 494.93333333C460.8 466.60266667 483.60106667 443.73333332 512 443.73333333 540.19413333 443.73333333 563.2 466.60266667 563.2 494.93333333l0 238.93333334C563.2 762.19733333 540.19413333 785.06666667 512 785.06666667c-28.39893334 0-51.2-22.86933333-51.2-51.2L460.8 494.93333333z" p-id="7174" fill="#e6a23c"></path></svg>
      </i>
     <i v-if="type=='error'" class="mei-message-icon">
      <svg t="1534130023768" class="icon" style="" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4951" xmlns:xlink="http://www.w3.org/1999/xlink" width="40" height="40"><path d="M511.970515 1024a512.560207 512.560207 0 0 1-511.970515-511.970515 512.560207 512.560207 0 0 1 511.970515-512.029485 512.560207 512.560207 0 0 1 511.970516 511.970515 512.560207 512.560207 0 0 1-511.970516 512.029485z m-5.07135-433.718399l145.889778 145.889778a58.615376 58.615376 0 0 0 41.750187 17.277973 58.556407 58.556407 0 0 0 41.632249-17.277973 58.556407 58.556407 0 0 0 17.277973-41.691218 58.556407 58.556407 0 0 0-17.277973-41.691218l-145.948747-145.948747 145.948747-145.948748a58.556407 58.556407 0 0 0 17.277973-41.691218 58.556407 58.556407 0 0 0-17.277973-41.691218 58.438468 58.438468 0 0 0-41.691218-17.277973 58.615376 58.615376 0 0 0-41.691218 17.277973l-145.948747 145.948748-145.948748-145.948748a58.615376 58.615376 0 0 0-41.691218-17.277973 58.615376 58.615376 0 0 0-41.691218 17.277973 58.969191 58.969191 0 0 0 0 83.382436l145.948748 145.948748-145.948748 145.948747a58.969191 58.969191 0 0 0 0 83.382436 58.615376 58.615376 0 0 0 41.750188 17.277973 58.556407 58.556407 0 0 0 41.691218-17.277973l145.830809-145.889778z" fill="#E76B66" p-id="4952"></path></svg>
      </i>
      <i v-if="type=='success'" class="mei-message-icon">
       <svg t="1534130184516" class="icon" style="" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5686" xmlns:xlink="http://www.w3.org/1999/xlink" width="40" height="40"><path d="M512 32C246.4 32 32 246.4 32 512s214.4 480 480 480 480-214.4 480-480S777.6 32 512 32z m214.4 400l-230.4 230.4c-16 16-38.4 16-54.4 0L320 540.8c-16-16-16-38.4 0-54.4s38.4-16 54.4 0l96 96 204.8-204.8c16-16 38.4-16 54.4 0 12.8 16 12.8 38.4-3.2 54.4z" fill="#28C828" p-id="5687"></path></svg>
      </i>
      <i v-if="type=='info'" class="mei-message-icon">
       <svg t="1534130289600" class="icon" style="" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="6492" xmlns:xlink="http://www.w3.org/1999/xlink" width="40" height="40"><path d="M511.263731 64.132646c-247.085468 0-447.387423 200.430892-447.387423 447.674972 0 247.245104 200.301955 447.674972 447.387423 447.674972s447.387423-200.429869 447.387423-447.674972C958.651155 264.563538 758.349199 64.132646 511.263731 64.132646zM566.040289 709.809181c-19.158339 24.910343-55.812143 67.205635-99.207489 67.393923-30.632671 0.187265-62.426795-19.016099-57.409524-70.02996 4.983501-51.015907 33.799804-122.066103 49.570999-178.949648 12.761652-45.947472 29.160133-84.242659-33.92567-83.001389 0-4.315282 0-8.643867 0-12.95915 42.360783-16.70752 107.673302-10.604522 151.413502-25.946952 1.725294 0 3.480264 0 5.205559 0-21.070898 106.189508-65.595974 202.761983-75.691913 311.208904 2.384303 1.820462 2.069125 1.442862 5.204535 2.574638 32.859385 10.980076 51.737338-64.884776 73.089645-59.612703C605.672938 665.820318 570.900993 703.407378 566.040289 709.809181zM559.675324 351.942631c-29.503964 0-53.429887-23.749914-53.429887-53.071729 0-29.318746 23.925922-53.069683 53.429887-53.069683 29.474288 0 53.399187 23.750937 53.399187 53.069683C613.073488 328.192717 589.148589 351.942631 559.675324 351.942631z" p-id="6493"></path></svg>
      </i>
      <span class="mei-message-con">{{message}}</span>
    </div>
  </transition>
</template>

<script type="text/babel">
  const typeMap = {
    success: 'success',
    info: 'info',
    warning: 'warning',
    error: 'error',
  
  };
  export default {
    data() {
      return {
        show:false,
        message:'',
        type:''
      };
    },
    computed: {
      iconClass() {
        return this.type
          ? `mei-message-icon mei-icon-${typeMap[this.type] }`
          : '';
      }
    },
  };
</script>
<style lang="less" scoped>
   .mei-message{
          position: fixed;
          z-index: 9999;
          min-width: 380px;
          border-radius: 8px;
          position: fixed;
          left: 50%;
          top: 30%;
          height:40px;
          transform: translate(-50%,-50%);
          background-color: #edf2fc;
          transition: opacity .3s,transform .4s;
          overflow: hidden;
          padding: 15px 15px 15px 20px;
          background-color:#ccc;
          color:#000;
          box-shadow: 0 0 25px rgba(0, 0, 0, 0.3);
          font-size: 16px;
    }
    .mei-message-success{
        background-color:#fff; 
        color:#333;
    }
    .mei-message-error{
        background-color:#fff; 
        color:#333;
    }
    .mei-message-warning{
        background-color:#fff; 
        color:#e6a23c;
    }
     .mei-message-info{
        background-color:#fff; 
        color:#333;
    }
    .mei-message-icon{
      display:inline-block;
      width:40px;
      height:40px;
      float:left;
    }
    .mei-message-con{
      line-height:40px;
      height:40px;
      display:inline-block;
      margin-left:10px;
    }
.mei-message-fade-enter-active {
  transition: all 0.3s linear;
}
.mei-message-fade-leave-active {
  transition: all 0.3s linear;
}
.mei-message-fade-enter, .mei-message-fade-leave-to
/* .slide-fade-leave-active for below version 2.1.8 */ {
  opacity: 0;
}
</style>