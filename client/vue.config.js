module.exports = {
  css: {
    loaderOptions: {
      // pass options to sass-loader
      sass: {
        // import `src/style/_main.sass` to all components
        data: `@import "~@/style/_main.sass"`
      }
    }
  },

  configureWebpack: {
    resolve: {
      alias: {
        vue$: 'vue/dist/vue.esm.js'
      }
    }
  },

  devServer: {
    proxy: {
      '/api/*': {
        target: 'http://localhost:3000',
        changeOrigin: true
      }
    }
  },

  pluginOptions: {
    i18n: {
      locale: 'ja',
      fallbackLocale: 'ja',
      localeDir: 'locales',
      enableInSFC: true
    }
  }
}
