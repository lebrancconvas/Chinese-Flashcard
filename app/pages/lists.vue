<template>
  <div>
    <header>
      <div>
        <h1>Word List</h1>
      </div>
    </header>
    <section id="word-list-section">
      <div v-for="word in results" id="word-list" :key="word.id">
        <v-card>
          <v-card-title>
            <h2>{{ word.vocab.simplified }} / {{ word.vocab.traditional }}</h2>
          </v-card-title>
          <v-card-text>
            <p>Category: {{ word.category }}</p>
            <p>Part of Speech: {{ word.pos }}</p>
          </v-card-text>
          <v-card-text>
            <h3>Pinyin</h3>
            <p>Mandarin: {{ word.pinyin.mandarin }}</p>
            <p>Cantonese: {{ word.pinyin.cantonese }}</p>
            <p>Hokkien: {{ word.pinyin.hokkien }}</p>
            <p>Teochew: {{ word.pinyin.teochew }}</p>
          </v-card-text>
          <v-card-text>
            <h3>Translation</h3>
            <p>English: {{ word.translation.english }}</p>
            <p>Thai: {{ word.translation.thai }}</p>
          </v-card-text>
        </v-card>
      </div>
    </section>
  </div>
</template>

<script>
  import axios from 'axios';
  export default {
    data() {
      return {
        results: []
      }
    },
    created() {
      const url = `http://localhost:9002/words`;
      axios.get(url)
        .then(res => {
          this.results = res.data;
        })
        .catch(err => {
          console.error(err);
          console.log(`[ERROR] Fetching Data Error.`);
        })
    }
  }
</script>

<style scoped>
  header {
    text-align: center;
  }

  #word-list-section {
    margin: 0 auto;
    display: grid;
    grid-template-columns: repeat(3, 1fr);
  }

  #word-list {
    width: 60%;
    margin: 5px auto;
  }
</style>
