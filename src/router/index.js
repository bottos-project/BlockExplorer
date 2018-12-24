import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router);

export default new Router({
  routes: [
    {
      path:'/',
      redirect:'/home'
    },
    {
      path:'/home',
      component:()=>import('@/views/home'/*webpackChunkName: "home"*/)
    },
    {
      path:'/blockchain',
      redirect:'/blockchain/blocks',
      component:()=>import('@/views/blockchain'/*webpackChunkName: "blockchain"*/),
      children: [
        {
          path: 'blocks',
          component:()=>import('@/views/blockchain/blocks'/*webpackChunkName: "blocks"*/)
        },
        {
          path: 'blocksDetail/:id',
          component:()=>import('@/views/blockchain/blocksDetail'/*webpackChunkName: "blocks"*/)
        },
        {
          path: 'transactions',
          component:()=>import('@/views/blockchain/transactions'/*webpackChunkName: "transactions"*/)
        },
        {
          path: 'transactionsDetail/:id',
          component:()=>import('@/views/blockchain/transactionsDetail'/*webpackChunkName: "transactions"*/)
        },
        {
          path: 'transfers',
          component:()=>import('@/views/blockchain/transfers'/*webpackChunkName: "transfers"*/)
        },
        {
          path: 'transfersDetail/:id',
          component:()=>import('@/views/blockchain/transfersDetail'/*webpackChunkName: "transfers"*/)
        },
        {
          path: 'accounts',
          component:()=>import('@/views/blockchain/accounts'/*webpackChunkName: "accounts"*/)
        },
        {
          path: 'accountsDetail/:id',
          component:()=>import('@/views/blockchain/accountsDetail'/*webpackChunkName: "accounts"*/)
        },
        {
          path: 'statistics',
          component:()=>import('@/views/blockchain/statistics'/*webpackChunkName: "statistics"*/)
        }
      ]
    },
    {
      path:'/contracts',
      redirect:'/contracts/index',
      component:()=>import('@/views/contracts'/*webpackChunkName: "contracts"*/),
      children: [
        {
          path: 'index',
          component:()=>import('@/views/contracts/lists'/*webpackChunkName: "contracts"*/)
        },
        {
          path: 'detail/:id',
          component:()=>import('@/views/contracts/detail'/*webpackChunkName: "contracts"*/)
        }
      ]
    },
    {
      path:'/nodes',
      redirect:'/nodes/index',
      component:()=>import('@/views/nodes'/*webpackChunkName: "nodes"*/),
      children: [
        {
          path: 'index',
          component:()=>import('@/views/nodes/lists'/*webpackChunkName: "nodes"*/)
        },
        {
          path: 'detail/:id',
          component:()=>import('@/views/nodes/detail'/*webpackChunkName: "nodes"*/)
        }
      ]
    },
    {
      path:'/supernode',
      redirect:'/supernode/index',
      component:()=>import('@/views/supernode'/*webpackChunkName: "supernode"*/),
      children: [
        {
          path: 'index',
          component:()=>import('@/views/supernode/lists'/*webpackChunkName: "supernode"*/)
        },
        {
          path: 'detail/:id',
          component:()=>import('@/views/supernode/detail'/*webpackChunkName: "supernode"*/)
        }
      ]
    },
    {
      path:'*',
      redirect:'/home'
    }
  ]
})
