<script setup>
import {computed, onMounted, ref} from "vue";

const search = async (cache = false) => {
  try {
    clearRes()
    clearMaybe()
    const response = await fetch(`http://localhost:8080/v1/api/movie/search?query=${field.value}&cache=${cache}`)
    const data = await response.json()
    console.log('response', response)
    console.log('data', data)
    if(data.hits?.hits) {
      res.value = [...data.hits?.hits]
    }
    if(data.suggest?.simple_phrase[0].options) {
      maybeItems.value = [...data.suggest?.simple_phrase[0].options]
    }
  } catch(err) {
    console.error(err)
  }
}

const clearRes = () => {
  res.value = []
}

const clearMaybe = () => {
  maybeItems.value = []
}

const handleMaybeClick = async (text) => {
  field.value = text
  await search(true)
}

const openModal = () => {
  modalIsActive.value = true
}

const closeModal = () => {
  modalIsActive.value = false
}

const setSrcVideo = (src) => {
  srcVideo.value = `https://www.youtube.com/embed/${src}`
}

const handleClickMore = (src) => {
  setSrcVideo(src)
  openModal()
}

const showSearchBtn = computed(() => {
  return field.value.length
})

const showMaybe = computed(() => {
  return maybeItems.value.length
})

const sortedRes = computed(() => {
  return res.value.sort((a, b) => {
    return new Date(b._source.release_date) - new Date(a._source.release_date)
  })
})

let field = ref('')
let res = ref([])
let maybeItems = ref([])
let modalIsActive = ref(false)
let srcVideo = ref('')

onMounted(() => {
  document.addEventListener(('keydown'), event => {
    if(event.code === 'Escape') {
      closeModal()
    }
  })
})
</script>

<template>
  <div class="search-w">
    <input class="search"
           placeholder="Поиск"
           v-model="field"
           @keydown.enter="field.length ? search() : ''"
    />
    <button
        class="search-btn"
        :style="{ display: showSearchBtn ? 'block' : 'none' }"
        @click="search()"
    >
      Поиск
    </button>
  </div>
  <div class="maybe"
    :style="{opacity: showMaybe ? '1' : '0'}"
  >
    Возможно вы имели ввиду:
    <div class="maybe-item"
      v-for="(el, index) of maybeItems" :key="index"
      @click="handleMaybeClick(el.text)"
    >
      {{el.text}}
      <span v-if="index === 0 && maybeItems.length > 1">
        или
      </span>
    </div>
  </div>
  <div class="items">
    <div class="item" v-for="item of sortedRes" :key="item.id">
      <div class="item-header">
        <div class="title">
          <span>Title:</span> {{item._source.title}}
        </div>
        <div class="subtitle">
          {{item._source.vote_average}}
        </div>
      </div>
      <div class="desc">
        <span>Overview:</span> {{item._source.overview}}
      </div>
      <div class="item-footer">
        <div class="release_date">
          <span>Release date:</span> {{item._source.release_date}}
        </div>
        <div class="more"
             v-if="item._source.trailer_yt"
             @click="handleClickMore(item._source.trailer_yt)"
        >
          Подробнее
        </div>
      </div>
    </div>
  </div>
  <div class="popup"
       v-if="modalIsActive"
       @click.self="closeModal()"
  >
    <div class="popup__body">
      <div class="popup__content">
        <iframe class="iframe" :src="srcVideo" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>
        <div class="popup-close" @click="closeModal()">
          <svg viewBox="0 -0.5 25 25" fill="none" xmlns="http://www.w3.org/2000/svg"><g id="SVGRepo_bgCarrier" stroke-width="0"></g><g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g><g id="SVGRepo_iconCarrier"> <path d="M6.96967 16.4697C6.67678 16.7626 6.67678 17.2374 6.96967 17.5303C7.26256 17.8232 7.73744 17.8232 8.03033 17.5303L6.96967 16.4697ZM13.0303 12.5303C13.3232 12.2374 13.3232 11.7626 13.0303 11.4697C12.7374 11.1768 12.2626 11.1768 11.9697 11.4697L13.0303 12.5303ZM11.9697 11.4697C11.6768 11.7626 11.6768 12.2374 11.9697 12.5303C12.2626 12.8232 12.7374 12.8232 13.0303 12.5303L11.9697 11.4697ZM18.0303 7.53033C18.3232 7.23744 18.3232 6.76256 18.0303 6.46967C17.7374 6.17678 17.2626 6.17678 16.9697 6.46967L18.0303 7.53033ZM13.0303 11.4697C12.7374 11.1768 12.2626 11.1768 11.9697 11.4697C11.6768 11.7626 11.6768 12.2374 11.9697 12.5303L13.0303 11.4697ZM16.9697 17.5303C17.2626 17.8232 17.7374 17.8232 18.0303 17.5303C18.3232 17.2374 18.3232 16.7626 18.0303 16.4697L16.9697 17.5303ZM11.9697 12.5303C12.2626 12.8232 12.7374 12.8232 13.0303 12.5303C13.3232 12.2374 13.3232 11.7626 13.0303 11.4697L11.9697 12.5303ZM8.03033 6.46967C7.73744 6.17678 7.26256 6.17678 6.96967 6.46967C6.67678 6.76256 6.67678 7.23744 6.96967 7.53033L8.03033 6.46967ZM8.03033 17.5303L13.0303 12.5303L11.9697 11.4697L6.96967 16.4697L8.03033 17.5303ZM13.0303 12.5303L18.0303 7.53033L16.9697 6.46967L11.9697 11.4697L13.0303 12.5303ZM11.9697 12.5303L16.9697 17.5303L18.0303 16.4697L13.0303 11.4697L11.9697 12.5303ZM13.0303 11.4697L8.03033 6.46967L6.96967 7.53033L11.9697 12.5303L13.0303 11.4697Z" fill="#000000"></path> </g></svg>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.items {
  display: flex;
  flex-wrap: wrap;
  gap: 20px;
  margin-top: 20px;
}
.item {
  box-sizing: border-box;
  flex: 0 0 calc(33.333% - 14px);
  padding: 10px;
  border: 1px solid #969696;
  color: #fff;
  display: flex;
  flex-direction: column;
}

.desc {
  flex: 1;
}

.item span {
  font-weight: bold;
}

.title {
  margin-bottom: 10px;
}

.release_date{
}

.search-w {
  position: relative;
  max-width: 500px;
  margin: 0 auto;
  width: 100%;
  box-sizing: border-box;
}

.search {
  border: 2px solid #2a2a2c;
  border-radius: 8px;
  padding: 10px;
  font-size: 22px;
  background: #232325;
  color: #5f5f60;
  box-sizing: border-box;
  display: block;
  width: 100%;
}

.search-btn {
  position: absolute;
  right: 4px;
  top: 4px;
  background: #fc0;
  border: none;
  border-radius: 8px;
  color: #000;
  cursor: pointer;
  font-size: .9rem;
  font-weight: 500;
  padding: 0 1rem;
  height: calc(100% - 8px);
  font-size: 22px;
}

.maybe {
  color: #fff;
  font-size: 22px;
  max-width: 500px;
  margin: 10px auto 0;
  display: flex;
  gap: 10px;
  min-height: 25px;
}

.maybe-item {
  cursor: pointer;
}

.item-header {
  display: flex;
  gap: 5px;
  justify-content: space-between;
}

.item-footer {
  display: flex;
  gap: 5px;
  justify-content: space-between;
  align-items: center;
  margin-top: 10px;

}

.more {
  cursor: pointer;
}

.popup {
  width: 100%;
  height: 100%;
  position: fixed;
  top: 0;
  left: 0;
  overflow-y: auto;
  overflow-x: hidden;
}

.popup::before {
  content: "";
  background: #000000;
  opacity: 0.4;
  position: fixed;
  left: 0px;
  top: 0px;
  width: 100%;
  height: 100%;
  z-index: 104;
}

.popup__body {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 30px 10px;
  min-height: 100%;
}

.popup__content {
  position: relative;
  background: #ffffff;
  border-radius: 15px;
  width: 95%;
  max-width: 946px;
  padding: 40px 20px;
  z-index: 105;
}

.popup-close {
  width: fit-content;
  position: absolute;
  right: 5px;
  top: 5px;
  cursor: pointer;
  & svg {
    width: 40px;
  }
}

.iframe {
  width: 100%;
  height: 500px;
  border-radius: 8px;
}
</style>