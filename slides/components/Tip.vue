<template>
  <div class="tip">
    <span class="icon">💡</span>
    <div class="content">
      <slot />
    </div>
  </div>
</template>

<style scoped>
.tip {
  display: flex;
  align-items: flex-start;
  gap: 0.5em;
  background: rgba(189, 147, 249, 0.1); /* subtle purple tint */
  border-left: 3px solid #bd93f9;       /* Dracula purple */
  border-radius: 6px;
  padding: 0.5em 0.75em;
  margin: 0.75em 0;
}

.icon {
  font-size: 1.2em;
  line-height: 1;
  color: #f1fa8c; /* yellow accent */
}

.content {
  flex: 1;
  color: #f8f8f2;
  font-size: 0.7em;
  line-height: 1.4;
}
</style>

