<template>
  <v-container>
    <div>
    <v-text-field
        v-model="searchQuery"
        label="Search Product"
        outlined
        class="mb-4"
      ></v-text-field>
      <v-icon-button></v-icon-button>
    </div>
    <v-row class="text-center">
      <v-list flat class="pt-0" style="display: contents;">
        <product-item
          v-for="product in filteredProducts"
          :key="product.id"
          :product="product"
        />
      </v-list>
    </v-row>
  </v-container>
</template>

<script>
import { mapState, mapActions } from 'vuex';
import ProductItem from '@/components/ProductListItem.vue';

export default {
  components: {
    ProductItem,
  },
  data() {
    return {
      searchQuery: '', 
    };
  },
  computed: {
    ...mapState(['products']),
    ...mapState(['brands']),
    filteredProducts() {
      const query = this.searchQuery.toLowerCase();
      return this.products.filter(product => product.name.toLowerCase().includes(query));
    },
  },
  methods: {
    ...mapActions(['listProducts']),
  },
  watch: {
    products: {
      handler(newProducts, oldProducts) {
        this.listProducts(newProducts);
      },
      deep: true,
    },
  },
  created() {
    this.listProducts();
  },
};
</script>
