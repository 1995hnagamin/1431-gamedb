const HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = {
  mode: 'production',
  entry: './src/index.tsx',
  module: {
    rules: [
      {
        test: /\.tsx$/,
        use: 'babel-loader',
      },
    ],
  },
  resolve: {
    extensions: [
      '.ts', '.js', '.tsx', '.jsx'
    ],
  },
  plugins: [
    new HtmlWebpackPlugin({
      filename: 'index.html',
      template: 'public/index.html'
    }),
  ],
};
