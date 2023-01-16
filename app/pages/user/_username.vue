<template>
  <div>
    <header>
      <div>
        <h1>User</h1>
      </div>
    </header>
    <section id="user-section">
      <div>
        User: {{ $route.params.username }}
      </div>
    </section>
  </div>
</template>

<script>
  export default {
    data() {
      return {

      }
    }
  }
</script>

<style scoped>
  * {
    text-align: center;
  }
</style>
