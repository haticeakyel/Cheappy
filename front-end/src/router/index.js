import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from '../views/HomeView.vue'
import AddProduct from '../views/AddProduct.vue'
import AddBrand from '../views/AddBrand.vue'
import ContractedBrands from '../views/ContractedBrands.vue'
import ProductDetail from '../views/ProductDetail.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/addProduct',
    name: 'addProduct',
    component:  AddProduct
  },
  {
    path: '/addBrand',
    name: 'addBrand',
    component:  AddBrand
  },
  {
    path: '/brands',
    name: 'contractedBrands',
    component:  ContractedBrands
  },
  {
    path: '/product/:productId',
    name: 'productDetail',
    component: ProductDetail
  }
]

const router = new VueRouter({
  routes,
  mode: 'history',
  base: '/',
})

export default router
