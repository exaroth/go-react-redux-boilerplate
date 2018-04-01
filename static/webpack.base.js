const path = require('path');
const webpack = require('webpack');

const config = {
  entry: {
    'babel-polyfill': ['babel-polyfill'],
    'app': './js/index.js',
  },
  output: {
    path: path.resolve('./build/js'),
    filename: '[name].min.js'
  },
  resolve: {
    modules: [
      path.resolve('./js'),
      path.resolve('./node_modules')
    ]
  },
  module: {
    rules: [
      {
        test: /\.js$/,
        loader: 'babel-loader',
        exclude: /node_modules/,
        query: {
            presets:['es2015', 'react']
        }
      },
      {
        test: /\.jsx$/,
        loader: 'babel-loader',
        exclude: /node_modules/,
        query: {
            presets:['es2015', 'react']
        },
      },
      {
        test: /\.less$/,
        loader: "style-loader!css-loader!autoprefixer-loader!less-loader"
      },
      {
        test: /\.css$/,
        loader: "style-loader!css-loader!autoprefixer-loader!less-loader"
      }
    ],
  },
  watch: false
}

module.exports = config
