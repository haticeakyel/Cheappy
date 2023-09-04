<template>
  <v-container>
    <v-row class="text-center">
      <v-col>
      <h1 style="color: navy;"> ALL PRODUCTS </h1>
      <productList/>
</v-col>
    </v-row>
  </v-container>
</template>

<script>
import ProductList from './ProductList.vue';
  export default {
    name: 'HelloWorld',

    data: () => ({
  
    }),
    components:{
      'productList': ProductList
    },
  }
</script>
