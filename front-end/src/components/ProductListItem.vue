<template>
<div>
  <v-alert
      class="ma-10"
      type="success"
      v-if="showSuccessAlert"
    >
      Product deleted successfully
    </v-alert>
  <div class="product-card elevation-5 ma-6 pa-10" style="background-color: white;">
    <h2 class="mb-8">{{ product.name }}</h2>
    <img :src="'data:image/png;base64,' + product.productImage" alt="Product Image" width="200" height="200" />
    <p>Category: <i>{{ getCategoryName(product.categoryId) }} </i></p>
    <p>Brand: <strong> {{ getBrand(product.brandId) }} </strong></p>
    <router-link :to="`/product/${product.id}`" style="background-color: orange; color: white; padding: 5px; border-radius: 5px; text-decoration: none;">View Product Detail</router-link>
    <v-icon @click="deleteProduct">mdi-delete</v-icon>
  </div>
  <delete-product
    v-if="isDeleting"
    @close="handleDeleteClose"
    :product="product.id"
  />
</div>
</template>


<script>
import { mapState, mapActions } from 'vuex';
import DeleteProduct from './DeleteProduct.vue';

export default {
props: ['product'],
data(){
  return{
    isDeleting: false,
    showSuccessAlert: false
  }
},
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
  deleteProduct() {
    this.isDeleting = true;
  },
  handleDeleteClose() {
    this.isDeleting = false;
    this.showSuccessAlert = true;
    setTimeout(() => {
          this.showSuccessAlert = false;
        }, 3000);
  },
  ...mapActions(['listBrands', 'listCategories']),
},
components:{
  'delete-product': DeleteProduct,
},
created() {
  this.listBrands();
  this.listCategories();
},
};
</script>