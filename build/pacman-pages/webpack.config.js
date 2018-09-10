'use strict';

const WebpackShellPlugin = require('webpack-shell-plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin')
const path = require('path');
const autoprefixer = require('autoprefixer')

var config = {
  context: path.join(__dirname, 'src'),
  entry: {
    index: './index.js'
  },
  output: {
    path: path.resolve(__dirname, 'dist'), // regular webpack
    filename: 'bundle.js'
  },
  devServer: {
    contentBase: path.resolve(__dirname, 'src') // dev server
  },
  plugins: [
    new WebpackShellPlugin({
      onBuildStart: ['gopherjs build --tags=pacman --output=dist/pacman.js --minify'],
    }),
    new HtmlWebpackPlugin({
      title: 'Pacman',
      template: 'index.html'
    }),
    new HtmlWebpackPlugin({
      template: 'pacman.html',
      filename: 'pacman.html',
      excludeChunks: ['index']
    })
  ],
  module: {
    rules: [
      {
        test: /\.js$/,
        use: {
          loader: 'babel-loader',
          options: {
            plugins: ['@babel/transform-runtime'],
            presets: ['@babel/env'],
            cacheDirectory: true,
            babelrc: false
          }
        },
        exclude: /node_modules/,
      },
      {
        test: /\.css$/,
        use: [
          'style-loader',
          {
            loader: 'css-loader',
            options: {
              importLoaders: 1,
            },
          },
          {
            loader: 'postcss-loader',
            options: {
              ident: 'postcss',
              plugins: [
                autoprefixer({
                  browsers: [
                    '>1%',
                    'last 4 versions',
                    'Firefox ESR',
                    'not ie < 9',
                  ],
                  flexbox: 'no-2009',
                })
              ]
            }
          }
        ]
      },
      {
        test: /\.scss$/,
        use: [
          'style-loader',
          {
            loader: 'sass-loader',
            options: {
              importLoaders: 1,
            },
          },
          {
            loader: 'postcss-loader',
            options: {
              ident: 'postcss',
              plugins: [
                autoprefixer({
                  browsers: [
                    '>1%',
                    'last 4 versions',
                    'Firefox ESR',
                    'not ie < 9',
                  ],
                  flexbox: 'no-2009',
                })
              ]
            }
          }
        ]
      },
      { test: /\.(png|jpg|svg)$/, loader: 'file-loader?name=images/[name].[ext]&publicPath=/'}
    ]
  },
  resolve: {
    extensions: ['.js', '.json']
  }
}

module.exports = config;
