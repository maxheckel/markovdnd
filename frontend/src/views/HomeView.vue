<script setup lang="ts">
import {onMounted, reactive} from "vue";
import {Env} from "@/env/build";
import {ShortToName} from "@/utils/ShortToName";

const books = reactive({
  crawled: []
})



onMounted(() => {
  fetch(Env.baseURL + "/get/crawled").then(res => res.json()).then(data => {books.crawled = data.crawled})
})
</script>

<template>
  <main>
    <h1>Choose a book</h1>
    <li v-for="item in books.crawled">
      <a :href="'story/'+item">{{ ShortToName(item) }}</a>
    </li>
  </main>
</template>
