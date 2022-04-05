<script setup lang="ts">
import {onMounted, reactive} from "vue";
import {Env} from "@/env/build";
import router from "@/router";
import {ShortToName} from "@/utils/ShortToName";

interface Story {
  story: string[];
  images: ImageAddon[];
  aloud_images: ImageAddon[];
  read_aloud: string[];
  elements: StoryElement[];
}

const story = reactive<Story>({
  story: [],
  images: [],
  aloud_images: [],
  read_aloud: [],
  elements: []
})


interface ImageAddon {
  based_on_word: string;
  position: number;
  type: string;
  url: string;
}

interface StoryElement {
  text?: string;
  type?: string;
  image?: ImageAddon | undefined;
}

var elementsAdded = 0;
var imagesAdded = 0;
const getNextElement = (lastElement: StoryElement | undefined): StoryElement | undefined => {
  if (elementsAdded >= 10) {
    return {} as StoryElement;
  }
  const returnEl = {} as StoryElement;
  if (Math.random() < 0.7 || (lastElement && lastElement.type === "aloud")) {
    returnEl.text = story.story[elementsAdded]
    returnEl.type = "story"
    elementsAdded++
  } else {
    returnEl.text = story.read_aloud.pop()
    returnEl.type = "aloud"
  }
  if (Math.random() < 0.2 && imagesAdded <= 3) {
    if (returnEl.type == "story") {
      const nextImage: ImageAddon = story.images.filter((i: any) => i.type === "story" && i.position === elementsAdded)[0]
      if (nextImage !== undefined) {
        returnEl.image = nextImage;
      }

    } else {
      returnEl.image = story.aloud_images.pop()
    }
    imagesAdded++
  } else if (returnEl.type == "aloud") {
    story.aloud_images.pop()
  }
  return returnEl
}


onMounted(() => {
  fetch(Env.baseURL + "/run/chain/" + router.currentRoute.value.params.name).then(res => res.json()).then(data => {
    story.story = data.story;
    story.read_aloud = data.read_aloud;
    story.images = data.images;
    story.aloud_images = data.images.filter((i: any) => i.type == "aloud")
    var nextElement = undefined;
    while (nextElement = getNextElement(nextElement)) {
      if (nextElement.text == null) {
        break;
      }
      story.elements.push(nextElement)
    }
  })
  console.log(story.elements)
})
</script>
<template>
  <div>
    <a class="back" href="/">Back</a>
    <h1>{{ ShortToName($route.params.name) }}... Kinda</h1>
    <div v-for="element in story.elements">
      <p class="aloud" v-if="element.type === 'aloud'">{{ element.text }}</p>
      <p v-if="element.type === 'story'">{{ element.text }}</p>
      <img :class="{imageLeft: Math.random() > 0.5, imageRight: Math.random() > 0.5, imageFull: Math.random() > 0.5}"
           v-if="element.image" :src="element.image.url">
    </div>


  </div>
</template>


<style scoped lang="scss">
@import '@/assets/colors.scss';

h1 {
  font-size: 32px;
  font-weight: normal;
  border-bottom: 3px solid $red;
}

.imageLeft {
  display: inline;
  float: left;
  margin: 10px;
}

.imageRight {
  display: inline;
  float: left;
  margin: 10px;
}

p {
  line-height: 30px;
  font-size: 16px;
}

img {
  max-width: 100%;
}

.aloud {
  overflow: auto;
  display: block;
  margin: 30px 0;
  line-height: 1.6;
  font-size: 14px;
  color: var(#242527);
  border: 8px solid transparent;
  border-image-repeat: repeat;
  border-image-slice: 8 8 8 8 fill;
  background-color: transparent;
  padding: 20px 20px 10px !important;
  position: relative;
  border-image-source: url(https://media.dndbeyond.com/ddb-compendium-client/146117d0758df55ed5ff299b916e9bd1.png);
}
</style>