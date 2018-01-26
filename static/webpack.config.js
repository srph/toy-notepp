const path = require('path')
const Dotenv = require('dotenv-webpack')

module.exports = {
  entry: path.resolve('src/index.js'),
  output: {
    path: path.resolve('public/dist'),
    filename: 'bundle.js'
  },
  module: {
    rules: [
      { test: /\.js$/, use: 'babel-loader', exclude: /node_modules/ },
      { test: /\.css$/, use: ['style-loader', 'css-loader'] },
      { test: /\.(woff|woff2|eot|ttf|otf|svg)$/, use: 'file-loader' }
    ]
  },
  plugins: [
    new Dotenv()
  ]
}