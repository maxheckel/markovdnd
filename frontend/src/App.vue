<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import HelloWorld from '@/components/HelloWorld.vue'
</script>

<template>
  <RouterView />
</template>

<style lang="scss">
@import '@/assets/colors.scss';

#app {
  max-width: 1280px;
  margin: 0 auto;
  padding: 2rem;
  font-family: Roboto,Helvetica,sans-serif;
  font-weight: normal;
}

html, body, #app{
  background-color: $background;
}
</style>
