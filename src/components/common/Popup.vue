<template>
  <div v-if="visible" class="popup">
    <!-- 弹窗的内容 -->
    <div class="popup-content">
      {{ content }}
    </div>
    <!-- 关闭按钮 -->
    <button class="close-button" @click="closePopup">关闭</button>
  </div>
</template>
 
<script>
export default {
  props: {
    visible: {
      type: Boolean,
      default: false
    },
    content: {
      type: String,
      default: ''
    }
  },
  methods: {
    closePopup() {
      this.$emit('close');
    }
  }
}
</script>
 
<style scoped>
.popup {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
}
 
.popup-content {
  background: #fff;
  padding: 20px;
  border-radius: 5px;
}
 
.close-button {
  margin-top: 10px;
}
</style>