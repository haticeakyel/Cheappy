<template>
  <v-app id="inspire">
    <v-navigation-drawer
      v-model="drawer"
      app
    >
      <v-list-item>
        <v-list-item-content>
          <v-list-item-title class="title">
            Cheappy App
          </v-list-item-title>
          <v-list-item-subtitle>
            See the cheapest prices
          </v-list-item-subtitle>
        </v-list-item-content>
      </v-list-item>

      <v-divider></v-divider>

      <v-list
        dense
        nav
      >
        <v-list-item
          v-for="item in items"
          :key="item.title"
          :to="item.to"
          link
        >
          <v-list-item-icon>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar
      app
      color="primary"
      dark
      src="background.png"
      prominent
    >
      <template v-slot:img="{ props }">
        <v-img
          v-bind="props"
        ></v-img>
      </template>

      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>


      <v-spacer></v-spacer>

      <template>
  <v-form>
    <v-container>
      <v-row>
        <v-col cols="12">
          <v-text-field
            filled
            clear-icon="mdi-close-circle"
            clearable
            label="Search"
            type="text"
            @click:clear="clearMessage"
          ></v-text-field>
        </v-col>
      </v-row>
    </v-container>
  </v-form>
</template>

      <v-btn icon>
        <v-icon>mdi-heart</v-icon>
      </v-btn>

      <v-btn icon>
        <v-icon>mdi-dots-vertical</v-icon>
      </v-btn>
    </v-app-bar>

    <v-main>
      <router-view></router-view>
    </v-main>
  </v-app>
</template>

<script>
  export default {
    data: () => ({
      drawer: null,
      items: [
        { title: 'Home', icon: 'mdi-home account', to: '/' },
        { title: 'Contracted Brands', icon: 'mdi-handshake-outline', to: '/brands' },
        { title: 'Add New Product', icon: 'mdi-plus', to: '/addProduct' },
        { title: 'Add New Brand', icon: 'mdi-plus', to: '/addBrand' },
      ],
      password: 'Password',
      show: false,
      marker: true,
    }),
    computed: {
      icon () {
        return this.icons[this.iconIndex]
      },
    },
  components:{
  },
  methods: {
      clearMessage () {
        this.message = ''
      },
    },
  }
</script>