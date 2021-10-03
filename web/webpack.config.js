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
};
