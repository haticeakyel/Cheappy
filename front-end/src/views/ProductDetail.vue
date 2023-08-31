<template>
    <div>
      <h2>{{ product.name }}</h2>
      <p>Category: {{ getCategoryName(product.categoryId) }}</p>
    </div>
  </template>
  
  <script>
  import { mapActions } from 'vuex';
  
  export default {
    props: ['productId'],
    computed: {
      product() {
        return this.$store.state.products.find(prod => prod.id === this.productId);
      },
    },
    methods: {
      getCategoryName(categoryId) {
        // Implement this method similar to your ProductItem component
      },
      methods: {
      ...mapActions(['listProducts']),
    },
      // ... other methods
    },
  };
  </script>
  