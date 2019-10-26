module.exports = {
  css: {
    loaderOptions: {
      // pass options to sass-loader
      sass: {
        // import `src/style/_main.sass` to all components
        prependData: '@import "~@/style/_main.sass"'
      }
    }
  },

  configureWebpack: {
    resolve: {
      alias: {
        vue$: 'vue/dist/vue.esm.js'
      }
    },
    context: __dirname,
    module: {
      rules: []
    },
    plugins: []
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
  },

  pwa: {
    name: 'My App',
    themeColor: '#4DBA87',
    msTileColor: '#000000',
    appleMobileWebAppCapable: 'yes',
    appleMobileWebAppStatusBarStyle: 'black',
    assetsVersion: '',
    manifestPath: 'manifest.json',
    iconPaths: {
      favicon32: 'img/icons/favicon-32x32.png',
      favicon16: 'img/icons/favicon-16x16.png',
      appleTouchIcon: 'img/icons/apple-touch-icon-152x152.png',
      maskIcon: 'img/icons/safari-pinned-tab.svg',
      msTileImage: 'img/icons/msapplication-icon-144x144.png'
    },
    workboxPluginMode: 'GenerateSW',
    workboxOptions: {
      // swSrc is required in InjectManifest mode.
      // swSrc: 'public/service-worker.js'
      // ...other Workbox options...
    }
  },

  productionSourceMap: false
}
