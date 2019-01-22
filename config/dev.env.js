'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
  NODE_ENV: '"development"',
  //BASE_API: '"http://192.168.2.123:8084"',
  // BASE_API: '"http://192.168.2.122:8080"',
  // BASE_API: '"http://127.0.0.1:8080"',
  BASE_API: '"http://localhost:8088"',
})
