const debug   = process.env.NODE_ENV !== "production";
const webpack = require('webpack');
const path    = require('path');

module.exports = {
  context: path.join(__dirname, ""),
  entry: "./js/src/go-react.js",
  module: {
    rules: [{
      test: /\.jsx?$/,
        exclude: /(node_modules|bower_components)/,
        use: [{
          loader: 'babel-loader',
          options: {
            presets: ['@babel/preset-react', '@babel/preset-env']
          }
        }]
      }]
    },
    output: {
      path: __dirname + "/js/assets",
      filename: "go-react.min.js"
    },
    plugins: debug ? [] : [
      new webpack.optimize.OccurrenceOrderPlugin(),
      new webpack.optimize.UglifyJsPlugin({ mangle: false, sourcemap: false }),
    ]
};
