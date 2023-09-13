<template >
  <v-row justify="center">
    <v-dialog
      :value="true"
      persistent
      width="auto"
    >
      
      <v-card>
        <v-card-title class="text-h5">
          Delete Product?
        </v-card-title>
        <v-card-text>Are you sure to delete this product?</v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="green"
            variant="text"
            @click="deleteProduct(product)"
          >
            Yes
          </v-btn>

          <v-btn
            color="red"
            variant="text"
            @click="$emit('close')"
          >
            No
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
</template>
<script>

  export default {
    props:['product'],
    methods:{
      deleteProduct(product){
        this.$store.dispatch('deleteProduct',this.product)
        this.$emit('close')
        
      },
    closeDialog() {
      this.$emit('close');
    }
    }
  }
</script> 