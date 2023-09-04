<template>
  <v-container>
    <div style="display: flex;">
    <div class="pt-5">
    
    <img :src="'data:image/png;base64,' + product.productImage" alt="Product Image" style="max-height: 500px; border-radius: 2px; border: 1px solid grey;" />
    </div>
    <div class="pa-16">
      <div>
      <h2>{{ product.name }}</h2>
      <p style="font-size: larger; color: grey;"> <i>{{product.description}}  </i></p>
      </div>
    <p style="color: darkcyan;"><strong> {{ getBrand(product.brandId) }}</strong></p>
    <p class="pt-2" style="color: grey;"> <i>Category: </i> <strong>{{ getCategoryName(product.categoryId) }} </strong></p>

    <h3 style="color: crimson;">Different Website Prices:</h3>
    <ul>
      <li v-for="(websitePrice, index) in product.websitePrices" :key="index">
        website: <strong><a style="text-decoration: none;" :href="getWebsiteRelativeUrl(websitePrice.websiteId)" target="_blank">{{ getWebsiteName(websitePrice.websiteId) }}</a> </strong> 
        Price:<strong>{{ websitePrice.price }}₺</strong> 
        Stock: {{ websitePrice.stock }}
      </li>
    </ul>
  </div>
  </div>

  <div style="border: 1px solid black; box-sizing: border-box; border-radius: 3px;" class="pa-6" >
    <h2 class="pa-8" style="color: crimson;">Best Price Lately</h2>
    

    <p>
      <i>  Website With Best Price:</i> <strong><a>{{ bestPriceWebsite.name }}</a></strong>
      </p>
  <p>
        Price:<strong> {{ bestPriceWebsite.price }}₺</strong> 
  </p>
  <p>
    Stock:  {{ bestPriceWebsite.stock }}
  </p>
  </div>

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
    bestPriceWebsite() {
      if (this.product && this.product.websitePrices.length > 0) {
        const minPrice = Math.min(
          ...this.product.websitePrices.map(price => price.price)
        );
        
        const bestPriceWebsite = this.product.websitePrices.find(
          price => price.price === minPrice
        );

        if (bestPriceWebsite) {
          bestPriceWebsite.name = this.getWebsiteName(bestPriceWebsite.websiteId);
        }

        return bestPriceWebsite;
      }
      return null; 
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
