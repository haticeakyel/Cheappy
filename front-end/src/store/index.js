import Vue from 'vue';
import Vuex from 'vuex';
import axios from 'axios';

Vue.use(Vuex);

const api = axios.create({
  baseURL: 'http://localhost:3001',
});

export default new Vuex.Store({
  state: {
    products: [],
    brands: [],
    websites: []
  },
  getters: {
    products: state => {
      return state.products;
    },
    brands: state => {
      return state.brands;
    },
    websites: state =>{
      return state.websites;
    }
  },
  mutations: {
    GET_PRODUCTS(state, products) {
      state.products = products;
    },
    GET_BRANDS(state, brands){
      state.brands = brands;
    },
    GET_WEBSITES(state,webssites){
      state.websites = webssites;
    }
  },
  actions: {
    async listProducts({ commit }) {
      try {
        const response = await api.get('/products');
        commit('GET_PRODUCTS', response.data);
      } catch (error) {
        console.log(error);
      }
    },

    async listBrands({ commit}){
      try {
        const response = await api.get('/brands');
        commit('GET_BRANDS', response.data)
      } catch (error) {
        console.log(error)
      }
    },
    
    async listWebsites({commit}){
      try {
        const response = await api.get('/websites');
        commit('GET_WEBSITES', response.data)
      } catch (error) {
        console.log(error)
      }
    }
  },
  modules: {},
});
