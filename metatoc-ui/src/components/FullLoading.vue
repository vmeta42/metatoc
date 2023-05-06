<template>
  <div id="app">
    <!-- 其他组件和内容 -->
    <transition name="fade">
      <div v-if="fullLoading" class="fullscreen-spin">
        <a-spin />
        <div class="loading-text">
          Data initialization in progress, please wait...
        </div>
      </div>
    </transition>
  </div>
</template>

<script>
export default {
  props: {
    fullLoading: {
      type: Boolean,
    },
    data() {
      return {
        tip: "Data initialization in progress, please wait...",
      };
    },
  },
  setup() {
    return {};
  },
};
</script>
<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s;
}
.fade-enter,
.fade-leave-to {
  opacity: 0;
}

.fullscreen-spin {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.8); /* 透明度从0.3增加到0.8 */
  z-index: 1000;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
}

.loading-text {
  margin-top: 1rem;
  font-size: 1.2rem; /* 字体大小从1rem增加到1.2rem */
  color: #fff; /* 字体颜色从黑色更改为白色 */
}
</style>
