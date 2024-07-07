// 将 File 对象转为 ArrayBuffer 
export function fileToBuffer(file) {
    return new Promise((resolve, reject) => {
      const fr = new FileReader()
      fr.onload = e => {
        resolve(e.target.result)
      }
      fr.readAsArrayBuffer(file)
      fr.onerror = () => {
        reject(new Error('转换文件格式发生错误'))
      }
    })
  }