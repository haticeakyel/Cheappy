<template>
  <div class="product-card elevation-5 ma-6 pa-10">
    <h2 class="mb-8">{{ product.name }}</h2>
    <img :src="'data:image/png;base64,' + product.productImage" alt="Product Image" width="200" height="200" />
    <p>Category: {{ getCategoryName(product.categoryId) }}</p>
    <p>Brand: {{ getBrand(product.brandId) }}</p>
    <router-link :to="`/product/${product.id}`" >View Product Detail</router-link>
  
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex';

export default {
  props: ['product'],
  computed: {
    ...mapState(['brands', 'categories']),
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
    goToProductDetail() {
      this.$router.push({ name: 'productDetail', params: { productId: this.product.id } });
    },
    ...mapActions(['listBrands', 'listCategories']),
  },
  created() {
    this.listBrands();
    this.listCategories();
  },
};
</script>
