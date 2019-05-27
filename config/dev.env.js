'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
  NODE_ENV: '"development"',
  // BASE_API: '"http://192.168.2.123:8084"',
  // BASE_API: '"http://114.67.80.209:8080"',
  BASE_API: '"http://125.94.34.23:8080"',
  // BASE_API: '"http://localhost:8080"',
  // BASE_API: '"http://explorerapi.chainbottos.com"',
  // BASE_API: '"http://explorerapi.chainbottos.com:8080"',
})
