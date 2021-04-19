<template>
 <div>
   <button @click="s()">start</button>
   <button @click="f()">finish</button>
   <button @click="g()">List</button>
 </div>
</template>

<script>
import {getAttendant, updateAttendant, getAttendantList} from "@/api/ActivitiesManagement";

export default {
  name: "start.vue",
  async created() {
    await getAttendant()
  },
  methods: {
    async s() {
      await getAttendant()
    },
    async f() {
      await updateAttendant()
    },
    async g() {
      let res = await getAttendantList()
      console.log(res.data.list)
    }
  }
}
</script>

<style scoped>

</style>
