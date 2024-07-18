import CryptoJS from 'crypto-js'
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

export function encodeWithSalt(input, salt) {
  return btoa(input + salt);
}

export function decodeWithSalt(input, salt) {
  try {
    const base64Data = atob(input);
    return base64Data.slice(0, -salt.length);
  } catch (e) {
    return 'Invalid encoded data';
  }
}

export const Salt = "j3j214hh421hk"
