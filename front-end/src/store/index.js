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
    websites: [],
    categories:[],
    snackbar:{
      show: false,
      text:''
    }
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
    },
    categories: state =>{
      return state.categories;
    },
    productById: state => productId => {
      return state.products.find(product => product.id === productId);
    },
  },
  mutations: {
    GET_PRODUCTS(state, products) {
      state.products = products;
    },
    GET_BRANDS(state, brands){
      state.brands = brands;
    },
    GET_WEBSITES(state,websites){
      state.websites = websites;
    },
    GET_CATEGORIES(state, categories){
      state.categories = categories;
    },
    ADD_BRAND(state,newBrand){
      state.brands.push(newBrand)
    },
    ADD_WEBSITE(state,websites){
      state.websites = websites;
    },
    ADD_PRODUCT(state,newProduct){
      state.products.push(newProduct);
    },
    DELETE_PRODUCT(state,productId){
      state.products = state.products.filter(product => product.id !== productId)
    },
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
    },
    async listCategories({commit}){
      try {
        const response = await api.get('/categories');
        commit('GET_CATEGORIES', response.data)
      } catch (error) {
        console.log(error)
      }
    },
    async addBrand({ commit }, brandData) {
      try {
        const response = await api.post('/addBrand', brandData);
        commit('ADD_BRAND', response.data);
      } catch (error) {
        console.error('Error adding brand:', error);
        throw error;
      }
    },
    async addProduct({commit}, productData){
      try {
        const response = await api.post('/addProduct',productData)
        commit('ADD_PRODUCT', response.data);
      } catch (error) {
        console.error('Error adding product:', error);
        throw error;
      }
    },
    async getProduct({commit},productId) {
      try{
        const response = await api.get(`/products/${productId}`);
        commit('GET_PRODUCTS',[response.data])
      }catch (error) {
        console.log(error);
      }
    },
    async deleteProduct({commit},productId){
      try{
        const response = await api.delete(`/products/${productId}`);
        commit('DELETE_PRODUCT', response.data)
      }catch (error) {
        console.log(error);
      }

    },
  },

  modules: {},
});
