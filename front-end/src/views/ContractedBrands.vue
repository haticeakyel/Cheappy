<template>
  <Container>
  <div class="container">
  <h6 class="title text-center blue--text pa-10 elevation-15">Contracted Brands </h6>
  <v-list-item-avatar class="blue elevation-19 ma-xl">
        <v-icon>mdi-handshake-outline</v-icon>
      </v-list-item-avatar>
    </div>
  <v-list class="brand-list">
    <v-list-item
      v-for="brand in brands"
      :key="brand.name"
      @mouseover="hoverBrand = brand.name"
      @mouseout="hoverBrand = null"
      class="brand-item"
    >
      <v-list-item-content :class="{ 'hovered': hoverBrand === brand.name }">
        <v-list-item-title>{{ brand.name }}</v-list-item-title>
      </v-list-item-content>
    </v-list-item>
  </v-list>
</Container>
</template>

<style>
.brand-list{
  display: flex;
  flex-direction: column;
  align-items: center;
  
}
.container{
  display: flex;
  flex-direction: column;
  align-items: center;
}
.brand-item {
  cursor: pointer;
}

.hovered {
  font-weight: bold; 
  color: darkblue;
  font-family: sans-serif;
  font-size: 500%;

}
</style>

<script>
import { validationMixin } from 'vuelidate'

export default {
  mixins: [validationMixin],

  data() {
    return {
      hoverBrand: null,
    };
  },

  computed: {
    brands() {
      return this.$store.getters.brands.map(brand => ({
        name: brand.name,
        brand: brand
      }));
    },
  },

  created() {
    this.$store.dispatch('listBrands');
  },
}
</script>
