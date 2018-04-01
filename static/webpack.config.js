const uglify = require('uglifyjs-webpack-plugin');

let config = require('./webpack.base.js')

config.mode = "production"
config.plugins = [
  new uglify(
    {
      uglifyOptions:{
        ecma: 6,
        output: {
          comments: false
        },
        compress: true
      }
    }
  )
];

module.exports = config
