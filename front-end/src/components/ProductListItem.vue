<template>
      <div class="product-card elevation-5 ma-6 pa-10">
        <h2 class="mb-8">{{ product.name }}</h2>
        <p>Category: {{ getCategoryName(product.categoryId) }}</p>
        <p>Brand: {{ getBrand(product.brandId) }}</p>
      </div>
  </template>
  
  <script>
  import { mapState, mapActions } from 'vuex';
  
  export default {
    props: ['product'],
    computed: {
      ...mapState(['websites']),
      ...mapState(['brands']),
      ...mapState(['categories']),
    },
    methods: {
      getCategoryName(categoryId) {
        const category = this.categories.find(cat => cat.id === categoryId);
        return category ? category.name : 'Unknown';
      },
      getBrand(brandId) {
        const brand = this.brands.find(brand => brand.id === brandId);
        return brand ? brand.name : 'Unknown';
      },
      ...mapActions(['listWebsites']),
      ...mapActions(['listBrands']),
      ...mapActions(['listCategories']),
    },
    created() {
      this.listWebsites();
      this.listBrands();
      this.listCategories();
    },
  };
  </script>
  