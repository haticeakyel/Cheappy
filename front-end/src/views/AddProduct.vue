<template>
  <form class="pa-8">
    <h2>Add New Product</h2>
    <v-select
      v-model="selectedCategoryId"
      :items="categories"
      item-text="name"
      item-value="category.id"
      label="Category"
      @change="$v.selectedCategoryId.$touch()"
      @blur="$v.selectedCategoryId.$touch()"
      required
></v-select>
      <v-text-field
      v-model="name"
      label="Name"
      required
      @input="$v.name.$touch()"
      @blur="$v.name.$touch()"
    ></v-text-field>
    <v-text-field
      v-model="description"
      label="Description"
      required
      @input="$v.description.$touch()"
      @blur="$v.description.$touch()"
    ></v-text-field>

    <div class="flex items-center">
      <v-select
        v-model="select"
        :items="brands"
        item-text="name"
        label="Brand"
        required
        @change="$v.select.$touch()"
        @blur="$v.select.$touch()"
      ></v-select>
      <router-link to="/addBrand" class="ml-2">
        <v-btn class="mb-8 mt-8">Add New Brand</v-btn>
      </router-link>
    </div>
    <h4 class="pb-6">Add stock data </h4>
    <v-select
      v-model="selectW"
      :items="websites"
      item-text="name"
      label="Website"
      required
    ></v-select>
      <v-text-field
          label="Price"
          type="number"
          v-model="price"
          :error="price < 0"
          prefix="₺"
          min="0"
        ></v-text-field>
        <v-text-field
          label="Stock"
          type="number"
          v-model="stock"
          :error="stock < 0"
          min="0"
        ></v-text-field>
    
    <v-btn
      class="mr-4"
      @click="submit"
    >
      Add Product
    </v-btn>
    <v-btn @click="clear">
      clear
    </v-btn>
    <v-alert
      class="ma-10"
      type="success"
      v-model="productAdded"
      transition="scale-transition"
      dismissible
     >
  Product added successfully
</v-alert> </form>
</template>

<script>
import { validationMixin } from 'vuelidate'
import { required, description } from 'vuelidate/lib/validators'

export default {
  mixins: [validationMixin],

  validations: {
    description: { required },
    name: { required},
    select:{required},
    selectedCategoryId: { required },
    selectW: { required },
  },

  data () {
    return {
    name: '',
    description: '',
    price: 0,
    stock: 0,
    select: null,
    selectW: null,
    productAdded: false, 
    selectedCategoryId: null,
    };
  },

  computed: {
    brands() {
      return this.$store.getters.brands.map(brand => ({
        name: brand.name,
        brand: brand
      }));
    },
    categories() {
      return this.$store.getters.categories.map(category => ({
        name: category.name,
        category: category
      }));
    },
    websites() {
      return this.$store.getters.websites.map(website => ({
        name: website.url,
        website: website
      }))
    },
  },

  created() {
    this.$store.dispatch('listBrands');
    this.$store.dispatch('listCategories');
    this.$store.dispatch('listWebsites');
  },

  methods: {
    async addNewProduct() {
      try {
        const productData = { 
          name: this.name,
          categoryId: this.selectedCategoryId, 
        };
        await this.$store.dispatch('addProduct', productData);
        this.productAdded = true; 
        setTimeout(() => {
          this.productAdded = false;
        }, 3000);
      } catch (error) {
        console.error('Error adding product:', error);
      }
    },

    submit() {
      this.$v.$touch()
      if (!this.$v.$invalid) {
        this.addNewProduct();
        this.name = '';
        this.price = 0;
        this.stock = 0;
        this.description = '';
        this.select = null;
        this.selectC = null;
        this.selectW = null;
      }
      else {
        console.log('olmadı :()')
      }
    },
    clear() {
      this.$v.$reset()
      this.name = ''
      this.price = 0
      this.stock = 0
      this.description = ''
      this.select = null 
      this.selectC = null
      this.selectW = null
    },
  },
}
</script>

