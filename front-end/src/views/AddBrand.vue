<template>
    <form class="pa-8">
      <h2>Add New Brand</h2>
      <h4 class="pa-4 font-italic">Please check contracted brands before adding new one</h4>
      <v-text-field
        v-model="name"
        :error-messages="nameErrors"
        :counter="10"
        label="Name"
        required
        @input="$v.name.$touch()"
        @blur="$v.name.$touch()"
      ></v-text-field>
  
      
      <v-btn
        class="mr-4"
        @click="submit"
      >
        Add Brand
      </v-btn>
      <v-btn @click="clear">
        clear
      </v-btn>
    </form>
  </template>
  
  <script>
    import { validationMixin } from 'vuelidate'
    import { required } from 'vuelidate/lib/validators'
  
    export default {
      mixins: [validationMixin],
  
      validations: {
        name: { required},
        description: { required },
        select: { required },
        
      },
  
      data: () => ({
        name: '',
        description: '',
        select: null,
      }),
  
      computed: {
        brands() {
        return this.$store.getters.brands.map(brand => ({
          name: brand.name,
          brand:brand
        }));
      },
      },
        created() {
      this.$store.dispatch('listBrands');
    },
      methods: {
        submit () {
          this.$v.$touch()
        },
        clear () {
          this.$v.$reset()
          this.name = ''
        },
      },
    }
  </script>
  