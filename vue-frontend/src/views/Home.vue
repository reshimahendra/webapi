<template>
  <div class="home">
    <h2>HELLO</h2>
    <router-view />
    <table class="table">
        <thead>
            <th>ID</th>
            <th>Nama</th>
            <th>Harga</th>
            <th>Rating</th>
            <th>Deskripsi</th>
        </thead>
        <tbody>
            <tr v-for="d in books.data" v-bind:key="d.id">
                <td class="has-text-left">{{ d.id }}</td>
                <td class="has-text-left">{{ d.title}}</td>
                <td class="has-text-right">{{ d.price }}</td>
                <td class="has-text-right">{{ d.rating }}</td>
                <td class="has-text-left">{{ d.description }}</td>
            </tr>
        </tbody>
    </table>
  </div>
</template>

<script>
import axios from 'axios'

export default {
    name: 'Home',
    data() {
        return {
            books: {}
        }
    },
    mounted() {
        this.loadBooks()
    },
    methods: {
        async loadBooks(){
            await axios
                .get('http://127.0.0.1:3888/v2/books')
                .then(res => {
                    console.log(res, 'success')
                    this.books = res.data
                })
            .catch(err => {
                console.log(err, 'error')
            })
        }
    },
}
</script>
