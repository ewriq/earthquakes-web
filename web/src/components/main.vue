<template>
<div class="flex justify-center">
  <div class="ml-80 mt-10">
        <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
      <thead  class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
        <tr class="tabletr">
          <th scope="col" class="px-6 py-3">Büyüklük</th>
          <th scope="col" class="px-6 py-3">Derinlik</th>
          <th scope="col" class="px-6 py-3">Zaman</th>
          <th scope="col" class="px-6 py-3">Konum</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="data in datam" :key="data.no" class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
          <td class="px-6 py-4">{{ data.violence }}</td>
          <td class="px-6 py-4">{{ data.depth }}km</td>
          <td class="px-6 py-4">{{ data.date }}</td>
          <td class="px-6 py-4">{{ data.location }}</td>
        </tr>
      </tbody>
    </table>
    </div>
    </div>
   
 
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        datam: [],
      };
    },
    created() {
      this.geldatam();
    },
    methods: {
      geldatam() {
        axios
          .get("http://localhost:3000/api/earthquakes")
          .then((response) => {
            this.datam = response.data;
            console.log(response.data);
          })
          .catch((error) => {
            console.error(error);
          });
      },
    },
  };
  </script>
  