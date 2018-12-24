<template>
    <div class="circle_container">
 
        <div :style="circleSize" :class="wrapClasses">
            <svg viewBox="0 0 100 100">
                <path :d="pathString" :stroke="trailColor" :stroke-width="5" :fill-opacity="0"/>
                <path :d="pathString" :stroke-linecap="strokeLinecap" :stroke="strokeColor" :stroke-width="6" fill-opacity="0" :style="pathStyle"/>
            </svg>
            <div :class="innerClasses">
                <slot><h3 class="">{{ jindu }}%</h3></slot>
            </div>
        </div>
 
    </div>
</template>
 
<script>
const prefixCls = 'ivu-chart-circle';
function  oneOf (value, validList) {
    for (let i = 0; i < validList.length; i++) {
        if (value === validList[i]) {
            return true;
        }
    }
    return false;
}
export default {
    props: {
        list: {
            type: Number,
            default: 0
        },
        jindu: {
            type: Number,
            default: 0
        },
        // percent: {
        //     type: Number,
        //     default: 0
        // },
        size: {
            type: Number,
            default: 100
        },
        strokeColor: {
            type: String,
            default: '#5bc1e8'
        },
        strokeLinecap: {
            validator (value) {
                return oneOf(value, ['square', 'round']);
            },
            default: 'round'
        },
        trailColor: {
            type: String,
            default: '#efefef'
        }
    },
    data () {
		return {
          percent:0
    	}
    },
    computed: {
        circleSize () {
            return {
                width: `${this.size}px`,
                height: `${this.size}px`
            };
        },
        radius () {
            return 50 - 6 / 2;
        },
        pathString () {
            return `M 50,50 m 0,-${this.radius}
            a ${this.radius},${this.radius} 0 1 1 0,${2 * this.radius}
            a ${this.radius},${this.radius} 0 1 1 0,-${2 * this.radius}`;
        },
        len () {
            return Math.PI * 2 * this.radius;
        },
        pathStyle () {
            return {
                'stroke-dasharray': `${this.len}px ${this.len}px`,
                'stroke-dashoffset': `${((100 - this.percent) / 100 * this.len)}px`,
                'transition': 'stroke-dashoffset 0.6s ease 0s, stroke 0.6s ease'
            };
        },
        wrapClasses () {
            return `${prefixCls}`;
        },
        innerClasses () {
            return `${prefixCls}-inner`;
        }
    },
    created() {
       setTimeout(()=>{
            this.percent = this.list;
        },1000);
    },
  	methods: {
 
  	}
}
</script>
 
<style scoped lang="less">
    .clearfix{zoom: 1;}
    .clearfix:after{
        content: '';
        display: block;
        height: 0;
        visibility: hidden;
        clear: both;
    }
    .circle_container{
        width: 100%;
        position: relative;
        line-height: 0;
    }
    .ivu-chart-circle-inner {
        width: 100%;
        text-align: center;
        position: absolute;
        left: 0;
        top: 50%;
        -webkit-transform: translateY(-50%);
        transform: translateY(-50%);
        line-height: 1;
    }
    .ivu-chart-circle {
        display: inline-block;
        position: relative;
    }
</style>