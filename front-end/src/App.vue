<template>
  <div>
    <div v-for="product in products">
      <h2>{{ product.name }}</h2>
    </div>
    <div v-for="brand in brands">
      <h2>{{ brand.name }}</h2>
    </div>
    <div v-for="website in websites">
      <h2>{{ website.url }}</h2>
    </div>
  </div>
</template>

<script>
export default {
  name: 'App',
  computed: {
    products() {
      return this.$store.getters.products;
    },
    brands() {
      return this.$store.getters.brands;
    },
    websites() {
      return this.$store.getters.websites;
    }
  },
  created() {
    this.$store.dispatch('listProducts');
    this.$store.dispatch('listBrands');
    this.$store.dispatch('listWebsites');
  }
};
</script>
