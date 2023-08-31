<template>
    <v-container>
      <v-row class="text-center">
        <v-list flat class="pt-0">
          <product-item
            v-for="product in products"
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
    computed: {
      ...mapState(['products']),
      ...mapState(['brands']),
    },
    methods: {
      ...mapActions(['listProducts']),
    },
    created() {
      this.listProducts();
    },
  };
  </script>
  