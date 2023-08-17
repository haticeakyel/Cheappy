<template>
  <form class="pa-8">
    <h2>Add New Product</h2>
    <v-select
        v-model="selectC"
        :items="categories"
        item-text="name"
        item-value="category"
        :error-messages="selectCErrors"
        label="Category"
        required
        @change="$v.selectC.$touch()"
        @blur="$v.selectC.$touch()"
      ></v-select>
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
        item-value="brand"
        :error-messages="selectErrors"
        label="Brand"
        required
        @change="$v.select.$touch()"
        @blur="$v.select.$touch()"
      ></v-select>
      <router-link to="/addBrand" class="ml-2">
        <v-btn class="mb-8 mt-8">Add New Brand</v-btn>
      </router-link>
    </div>
    
    <v-btn
      class="mr-4"
      @click="submit"
    >
      Add Product
    </v-btn>
    <v-btn @click="clear">
      clear
    </v-btn>
  </form>
</template>

<script>
  import { validationMixin } from 'vuelidate'
  import { required, maxLength, description } from 'vuelidate/lib/validators'

  export default {
    mixins: [validationMixin],

    validations: {
      description: { required },
      select: { required },
      selectC: { required },
      
    },

    data: () => ({
      name: '',
      description: '',
      select: null,
      selectC:null,
    }),

    computed: {
      selectErrors () {
        const errors = []
        if (!this.$v.select.$dirty) return errors
        !this.$v.select.required && errors.push("Please select brand, if doesn't exist please add")
        return errors
      },
      selectCErrors () {
        const errors = []
        if (!this.$v.selectC.$dirty) return errors
        !this.$v.selectC.required && errors.push("Please select category")
        return errors
      },
      brands() {
      return this.$store.getters.brands.map(brand => ({
        name: brand.name,
        brand:brand
      }));
    },
    categories() {
      return this.$store.getters.categories.map(category => ({
        name: category.name,
        category:category
      }));
    },
    },
      created() {
    this.$store.dispatch('listBrands');
    this.$store.dispatch('listCategories');
  },
    methods: {
      submit () {
        this.$v.$touch()
      },
      clear () {
        this.$v.$reset()
        this.name = ''
        this.description = ''
        this.select = null
      },
    },
  }
</script>
