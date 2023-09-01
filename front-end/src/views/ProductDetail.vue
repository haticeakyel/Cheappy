<template>
  <v-container>
    <h2>{{ product.name }}</h2>
    <p>Category: {{ getCategoryName(product.categoryId) }}</p>
    <p>{{product.description}} </p>
    <p>Brand: {{ getBrand(product.brandId) }}</p>
    <h3>Different Website Prices:</h3>
    <ul>
      <li v-for="(websitePrice, index) in product.websitePrices" :key="index">
        Website: <a :href="getWebsiteRelativeUrl(websitePrice.websiteId)" target="_blank">{{ getWebsiteName(websitePrice.websiteId) }}</a>
        Price: {{ websitePrice.price }}
        Stock: {{ websitePrice.stock }}
      </li>
    </ul>
  </v-container>
</template>

<script>
import { mapActions } from 'vuex';

export default {
  computed: {
    product() {
      return this.$store.state.products.find(
        prod => prod.id === this.$route.params.productId
      );
    },
  },
  methods: {
    getCategoryName(categoryId) {
      const category = this.$store.state.categories.find(cat => cat.id === categoryId);
      return category ? category.name : 'Unknown';
    },
    getBrand(brandId) {
      const brand = this.$store.state.brands.find(brand => brand.id === brandId);
      return brand ? brand.name : 'Unknown';
    },
    getWebsiteName(websiteId) {
        const website = this.$store.state.websites.find(web => web.id === websiteId);
        return website ? website.url : 'Unknown';
     
    },
    getWebsiteRelativeUrl(websiteId) {
    const website = this.$store.state.websites.find(web => web.id === websiteId);
    return website ? '/' + website.url : '#'; 
  },
    ...mapActions(['listProducts', 'listCategories','listBrands','listWebsites']),
  },
  created() {
    this.listBrands();
    this.listProducts();
    this.listCategories();
    this.listWebsites();
  },
};
</script>
