<template>
  <form class="pa-8">
    <h2>Add New Brand</h2>
    <h4 class="pa-4 font-italic">Please check contracted brands before adding a new one</h4>
    <v-text-field
      v-model="name"
      :counter="10"
      label="Name"
      required
      @input="$v.name.$touch()"
      @blur="$v.name.$touch()"
    ></v-text-field>

    <v-btn class="mr-4" @click="submit" :disabled="isAddButtonDisabled">
      Add Brand
    </v-btn>
    <v-btn @click="clear">Clear</v-btn>
    <v-alert class="ma-10" type="success"  v-model="brandAdded">Brand added successfully</v-alert>
  </form>
</template>

<script>
import { validationMixin } from 'vuelidate';
import { required } from 'vuelidate/lib/validators';

export default {
  mixins: [validationMixin],

  validations: {
    name: { required },
  },

  data() {
    return {
      name: '',
      brandAdded: false, // New property to control alert visibility
    };
  },

  computed: {
    isAddButtonDisabled() {
      return this.$v.name.$invalid;
    },
  },

  methods: {
    async addNewBrand() {
      try {
        const brandData = { name: this.name };
        await this.$store.dispatch('addBrand', brandData);
        this.brandAdded = true; 
        setTimeout(() => {
          this.brandAdded = false;
        }, 3000);
      } catch (error) {
        console.error('Error adding brand:', error);
      }
    },

    async submit() {
      this.$v.$touch();

      if (!this.$v.$invalid) {
        this.addNewBrand();
        this.name = '';
      }
    },

    clear() {
      this.$v.$reset();
      this.name = '';
    },
  },
};
</script>
