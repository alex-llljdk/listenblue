module.exports = {
    baseUrl: './',
    assetsDir: 'static',
    productionSourceMap: false,
    configureWebpack: {
        module: {
          rules: [
            {
              test: /\.mjs$/,
              include: /node_modules/,
              type: "javascript/auto"
            }
          ] 
        }
    },
}

    // devServer: {
    //     proxy: {
    //         '/api':{
    //             target:'http://jsonplaceholder.typicode.com',
    //             changeOrigin:true,
    //             pathRewrite:{
    //                 '/api':''
    //             }
    //         }
    //     }
    // }