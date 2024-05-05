<script setup>
import { ref } from "vue";

const search = async () => {
  try {
    const response = await fetch(`http://localhost:8080/v1/api/movie/search?query=${field.value}`)
    const data = await response.json()
    console.log('response', response)
    console.log('data', data)
    if(data.hits?.hits) {
      res.value = [...data.hits?.hits]
    }
  } catch(err) {
    console.error(err)
  }
}

let field = ref('Harry')
let res = ref([])
</script>

<template>
  <input placeholder="Поиск"
    v-model="field"
    @keydown.enter="search()"
  />
  <button
    @click="search()"
  >
    Поиск
  </button>
  <div class="items">
    <div class="item" v-for="item of res" :key="item.id">
      <div class="title">
        <span>Title:</span> {{item._source.title}}
      </div>
      <div>
        <span>Overview:</span> {{item._source.overview}}
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
}

.item span {
  font-weight: bold;
}

.title {
  margin-bottom: 10px;
}
</style>